package main

import (
	"bufio"
	"bytes" // Added for request body
	"context"
	"encoding/json"
	"fmt"
	"io"       // Added for reading response body
	"net/http" // Added for HTTP requests
	"os"
	"path"
	"path/filepath"
	"strings"
	"time" // Added for http client timeout

	// We keep jsonschema for convenience, although technically schema could be built manually
	"github.com/invopop/jsonschema"
)

// --- Configuration ---
// Read from environment variables
var (
	openaiAPIKey      = os.Getenv("OPENAI_API_KEY")                           // Use OPENAI_API_KEY now
	openaiAPIEndpoint = os.Getenv("OPENAI_API_BASE") + "/v1/chat/completions" // Allow overriding base URL
	openaiModel       = os.Getenv("OPENAI_MODEL")                             // Allow specifying model
)

// --- OpenAI API Structs ---

type OpenAIChatCompletionRequest struct {
	Model       string                        `json:"model"`
	Messages    []OpenAIChatCompletionMessage `json:"messages"`
	Tools       []OpenAIChatCompletionTool    `json:"tools,omitempty"`
	ToolChoice  any                           `json:"tool_choice,omitempty"` // "auto" or specific tool
	MaxTokens   int                           `json:"max_tokens,omitempty"`
	Temperature float32                       `json:"temperature,omitempty"`
	// Add other OpenAI parameters as needed (top_p, stream, etc.)
}

type OpenAIChatCompletionMessage struct {
	Role       string                         `json:"role"`                   // "system", "user", "assistant", "tool"
	Content    string                         `json:"content,omitempty"`      // For text content or tool result
	ToolCalls  []OpenAIChatCompletionToolCall `json:"tool_calls,omitempty"`   // For assistant requesting tools
	ToolCallID string                         `json:"tool_call_id,omitempty"` // For tool role messages
	Name       string                         `json:"name,omitempty"`         // For tool role messages (function name) - Optional by OpenAI spec but sometimes useful
}

type OpenAIChatCompletionTool struct {
	Type     string                                 `json:"type"` // Always "function" for now
	Function OpenAIChatCompletionFunctionDefinition `json:"function"`
}

type OpenAIChatCompletionFunctionDefinition struct {
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Parameters  map[string]any `json:"parameters"` // Use map[string]any to represent JSON schema object
}

type OpenAIChatCompletionResponse struct {
	ID      string                       `json:"id"`
	Object  string                       `json:"object"`
	Created int64                        `json:"created"`
	Model   string                       `json:"model"`
	Choices []OpenAIChatCompletionChoice `json:"choices"`
	Usage   OpenAIUsage                  `json:"usage"`
}

type OpenAIChatCompletionChoice struct {
	Index        int                         `json:"index"`
	Message      OpenAIChatCompletionMessage `json:"message"`       // The assistant's response message
	FinishReason string                      `json:"finish_reason"` // e.g., "stop", "tool_calls"
}

type OpenAIChatCompletionToolCall struct {
	ID       string                           `json:"id"`   // ID to match with tool response
	Type     string                           `json:"type"` // Always "function"
	Function OpenAIChatCompletionFunctionCall `json:"function"`
}

type OpenAIChatCompletionFunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"` // Arguments are a *string* containing JSON
}

type OpenAIUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// --- Tool Definition (Mostly Unchanged, but Schema Generation Adapted) ---

type ToolDefinition struct {
	Name        string
	Description string
	// InputSchema now map[string]any to match OpenAI's parameter schema format
	InputSchema map[string]any
	Function    func(input json.RawMessage) (string, error) // Input is JSON string from OpenAI args
}

// GenerateSchema adapted to return map[string]any
func GenerateSchema[T any]() map[string]any {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties:  false,
		DoNotReference:             true, // Keep definitions inline for OpenAI
		RequiredFromJSONSchemaTags: true, // Respect `jsonschema:"required"`
	}
	var v T
	schema := reflector.Reflect(v)

	// Convert the jsonschema.Schema to map[string]any expected by OpenAI
	// This is a simplification; a full conversion might be more complex
	schemaBytes, _ := json.Marshal(schema)
	var schemaMap map[string]any
	_ = json.Unmarshal(schemaBytes, &schemaMap)

	// OpenAI expects parameters schema directly, remove unnecessary outer layers if present
	if props, ok := schemaMap["properties"]; ok {
		schemaMap["properties"] = props
	}
	if req, ok := schemaMap["required"]; ok {
		schemaMap["required"] = req
	}
	schemaMap["type"] = "object" // Ensure root type is object

	return schemaMap
}

// --- Tools (ReadFile, ListFiles, EditFile - Implementations Unchanged, Definitions Adapted) ---

// ReadFile Tool Definition
type ReadFileInput struct {
	Path string `json:"path" jsonschema_description:"The relative path of a file in the working directory." jsonschema:"required"`
}

var ReadFileDefinition = ToolDefinition{
	Name:        "read_file",
	Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
	InputSchema: GenerateSchema[ReadFileInput](),
	Function:    ReadFile, // Function implementation remains the same
}

func ReadFile(input json.RawMessage) (string, error) { // Implementation unchanged
	readFileInput := ReadFileInput{}
	err := json.Unmarshal(input, &readFileInput)
	if err != nil {
		return "", fmt.Errorf("failed to parse input for read_file: %w. Input was: %s", err, string(input))
	}
	if readFileInput.Path == "" {
		return "", fmt.Errorf("missing required parameter 'path' for read_file")
	}
	content, err := os.ReadFile(readFileInput.Path)
	if err != nil {
		return "", fmt.Errorf("error reading file '%s': %w", readFileInput.Path, err)
	}
	return string(content), nil
}

// ListFiles Tool Definition
type ListFilesInput struct {
	Path string `json:"path,omitempty" jsonschema_description:"Optional relative path to list files from. Defaults to current directory if not provided."`
}

var ListFilesDefinition = ToolDefinition{
	Name:        "list_files",
	Description: "List files and directories at a given path. If no path is provided, lists files in the current directory. Returns a JSON array of strings, directories have a trailing slash.",
	InputSchema: GenerateSchema[ListFilesInput](),
	Function:    ListFiles, // Function implementation remains the same
}

// ListFiles function implementation unchanged
func ListFiles(input json.RawMessage) (string, error) {
	listFilesInput := ListFilesInput{}
	if len(input) > 0 && string(input) != "null" {
		err := json.Unmarshal(input, &listFilesInput)
		if err != nil {
			return "", fmt.Errorf("failed to parse input for list_files: %w. Input was: %s", err, string(input))
		}
	}
	dir := "."
	if listFilesInput.Path != "" {
		dir = listFilesInput.Path
	}
	var files []string
	err := filepath.WalkDir(dir, func(currentPath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(dir, currentPath)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", currentPath, err)
		}
		if relPath == "." {
			return nil
		}
		if d.IsDir() {
			files = append(files, relPath+"/")
		} else {
			files = append(files, relPath)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error listing files in '%s': %w", dir, err)
	}
	result, err := json.Marshal(files)
	if err != nil {
		return "", fmt.Errorf("failed to marshal file list to JSON: %w", err)
	}
	return string(result), nil
}

// EditFile Tool Definition
type EditFileInput struct {
	Path   string `json:"path" jsonschema_description:"The path to the file" jsonschema:"required"`
	OldStr string `json:"old_str" jsonschema_description:"Text to search for. If empty and file doesn't exist, creates the file with new_str as content. If not empty, MUST match exactly (limitation)."`
	NewStr string `json:"new_str" jsonschema_description:"Text to replace old_str with, or the initial content if creating a new file." jsonschema:"required"`
}

var EditFileDefinition = ToolDefinition{
	Name:        "edit_file",
	Description: `Make edits to a text file. Replaces ALL occurrences of 'old_str' with 'new_str'. If 'old_str' is empty and the file doesn't exist, it creates it with 'new_str'.`,
	InputSchema: GenerateSchema[EditFileInput](),
	Function:    EditFile, // Function implementation remains the same
}

// EditFile and createNewFile function implementations unchanged
func createNewFile(filePath, content string) (string, error) {
	dir := path.Dir(filePath)
	if dir != "." && dir != "" {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create directory '%s': %w", dir, err)
		}
	}
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to create file '%s': %w", filePath, err)
	}
	return fmt.Sprintf("Successfully created file %s", filePath), nil
}
func EditFile(input json.RawMessage) (string, error) {
	editFileInput := EditFileInput{}
	err := json.Unmarshal(input, &editFileInput)
	if err != nil {
		return "", fmt.Errorf("failed to parse input for edit_file: %w. Input was: %s", err, string(input))
	}
	if editFileInput.Path == "" {
		return "", fmt.Errorf("invalid input: 'path' cannot be empty")
	}
	content, err := os.ReadFile(editFileInput.Path)
	if err != nil {
		if os.IsNotExist(err) && editFileInput.OldStr == "" {
			return createNewFile(editFileInput.Path, editFileInput.NewStr)
		}
		return "", fmt.Errorf("error reading file '%s': %w", editFileInput.Path, err)
	}
	oldContent := string(content)
	newContent := strings.Replace(oldContent, editFileInput.OldStr, editFileInput.NewStr, -1)
	if oldContent == newContent && editFileInput.OldStr != "" {
		if oldContent == editFileInput.NewStr {
			return "OK (no change needed, content already matched)", nil
		}
		return "", fmt.Errorf("old_str '%s' not found in file '%s'", editFileInput.OldStr, editFileInput.Path)
	}
	err = os.WriteFile(editFileInput.Path, []byte(newContent), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing file '%s': %w", editFileInput.Path, err)
	}
	return "OK", nil
}

// --- Agent Logic (Adapted for OpenAI) ---

type Agent struct {
	httpClient     *http.Client // Use standard HTTP client
	getUserMessage func() (string, bool)
	tools          map[string]ToolDefinition // Use map for easy lookup by name
	model          string                    // Store the target model name
	systemPrompt   string                    // Store the system prompt
}

func NewAgent(
	getUserMessage func() (string, bool),
	tools []ToolDefinition,
	model string,
) *Agent {
	toolMap := make(map[string]ToolDefinition)
	for _, tool := range tools {
		toolMap[tool.Name] = tool
	}
	return &Agent{
		httpClient:     &http.Client{Timeout: 60 * time.Second}, // Add a timeout
		getUserMessage: getUserMessage,
		tools:          toolMap,
		model:          model,
		// Define the system prompt here or pass it in
		systemPrompt: "You are a helpful Go programmer assistant. You have access to tools to interact with the local filesystem (read, list, edit files). Use them when appropriate to fulfill the user's request. When editing, be precise about the changes. Respond ONLY with tool calls if you need to use tools, otherwise respond with text.",
	}
}

// Run method now manages OpenAI message history
func (a *Agent) Run(ctx context.Context) error {
	// Use OpenAI message structure for history
	conversation := []OpenAIChatCompletionMessage{
		{Role: "system", Content: a.systemPrompt}, // Start with system prompt
	}

	fmt.Println("Chat with AI (use 'ctrl-c' to quit)")

	for {
		fmt.Print("\u001b[94mYou\u001b[0m: ") // Blue prompt for user
		userInput, ok := a.getUserMessage()
		if !ok { // Handle EOF or scanner error (e.g., ctrl-d)
			fmt.Println("\nExiting.")
			break
		}
		if userInput == "" {
			continue
		}

		// Add user message to conversation
		conversation = append(conversation, OpenAIChatCompletionMessage{Role: "user", Content: userInput})

		// --- Main Loop: Call API -> Handle Response -> Execute Tools -> Call API ---
		for { // Inner loop to handle potential multi-turn tool calls
			// Call OpenAI API
			resp, err := a.callOpenAICompletion(ctx, conversation)
			if err != nil {
				fmt.Printf("\u001b[91mAPI Error\u001b[0m: %s\n", err.Error())
				// Remove the last user message before retrying? Or let user re-prompt.
				// For simplicity, break inner loop and let user re-prompt.
				break
			}

			// Process the response
			if len(resp.Choices) == 0 {
				fmt.Println("\u001b[91mError\u001b[0m: OpenAI response contained no choices.")
				break // Break inner loop, let user re-prompt
			}
			assistantMessage := resp.Choices[0].Message

			// Add assistant's message (text and/or tool calls) to conversation
			conversation = append(conversation, assistantMessage)

			// --- Handle Text Response ---
			if assistantMessage.Content != "" {
				fmt.Printf("\u001b[93mAI\u001b[0m: %s\n", assistantMessage.Content) // Yellow for AI
			}

			// --- Handle Tool Calls ---
			if len(assistantMessage.ToolCalls) == 0 {
				// No tools called, break inner loop and wait for next user input
				break
			}

			// Execute tools and collect results
			toolResults := []OpenAIChatCompletionMessage{}
			for _, toolCall := range assistantMessage.ToolCalls {
				if toolCall.Type != "function" {
					continue // Skip non-function tool calls if any
				}

				toolName := toolCall.Function.Name
				toolArgs := toolCall.Function.Arguments // This is a JSON *string*

				fmt.Printf("\u001b[92mTool Call\u001b[0m: %s(%s)\n", toolName, toolArgs) // Green

				toolDef, found := a.tools[toolName]
				var resultMsg OpenAIChatCompletionMessage
				if !found {
					errorMsg := fmt.Sprintf("tool '%s' not found by agent", toolName)
					fmt.Printf("\u001b[91mTool Error\u001b[0m: %s\n", errorMsg)
					resultMsg = OpenAIChatCompletionMessage{
						Role:       "tool",
						ToolCallID: toolCall.ID,
						Content:    errorMsg, // Report error back to OpenAI
						Name:       toolName,
					}
				} else {
					// Execute the actual tool function
					// Note: toolArgs is a JSON string, pass it as json.RawMessage
					toolOutput, err := toolDef.Function(json.RawMessage(toolArgs))
					if err != nil {
						errorMsg := fmt.Sprintf("error executing tool '%s': %s", toolName, err.Error())
						fmt.Printf("\u001b[91mTool Error\u001b[0m: %s\n", errorMsg)
						resultMsg = OpenAIChatCompletionMessage{
							Role:       "tool",
							ToolCallID: toolCall.ID,
							Content:    errorMsg, // Report error back to OpenAI
							Name:       toolName,
						}
					} else {
						// Log successful tool execution result (optional)
						// fmt.Printf("\u001b[92mTool Result\u001b[0m: %s\n", toolOutput)
						resultMsg = OpenAIChatCompletionMessage{
							Role:       "tool",
							ToolCallID: toolCall.ID,
							Content:    toolOutput, // Send success result back to OpenAI
							Name:       toolName,
						}
					}
				}
				toolResults = append(toolResults, resultMsg)
			} // End of processing tool calls for one response

			// Add all tool results to conversation history
			conversation = append(conversation, toolResults...)

			// Continue inner loop to call API again with tool results
		} // End of inner loop (API call -> handle response/tools)

	} // End of outer loop (read user input)

	return nil
}

// callOpenAICompletion uses standard library http client
func (a *Agent) callOpenAICompletion(ctx context.Context, conversation []OpenAIChatCompletionMessage) (*OpenAIChatCompletionResponse, error) {

	// Prepare tools in OpenAI format
	openaiTools := []OpenAIChatCompletionTool{}
	for _, toolDef := range a.tools {
		openaiTools = append(openaiTools, OpenAIChatCompletionTool{
			Type: "function",
			Function: OpenAIChatCompletionFunctionDefinition{
				Name:        toolDef.Name,
				Description: toolDef.Description,
				Parameters:  toolDef.InputSchema,
			},
		})
	}

	// Build request payload
	requestPayload := OpenAIChatCompletionRequest{
		Model:       a.model,
		Messages:    conversation,
		Tools:       openaiTools,
		ToolChoice:  "auto", // Let the model decide when to use tools
		MaxTokens:   2048,   // Or make configurable
		Temperature: 0.7,    // Reasonable default
	}

	// Marshal payload to JSON
	jsonPayload, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", openaiAPIEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)
	// Add other headers if required by the specific OpenAI-compatible provider

	// Send request
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		// Try to include error message from response body
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Unmarshal response JSON
	var response OpenAIChatCompletionResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		// Include raw body in error for debugging
		return nil, fmt.Errorf("failed to unmarshal response JSON: %w. Body: %s", err, string(bodyBytes))
	}

	return &response, nil
}

// --- Main Function (Adapted) ---

func main() {
	// --- Configuration Checks ---
	if openaiAPIKey == "" {
		fmt.Fprintln(os.Stderr, "\u001b[91mError: OPENAI_API_KEY environment variable not set.\u001b[0m")
		os.Exit(1)
	}
	if os.Getenv("OPENAI_API_BASE") == "" {
		// Default to official OpenAI endpoint if base URL not set
		openaiAPIEndpoint = "https://api.openai.com/v1/chat/completions"
		fmt.Println("Info: OPENAI_API_BASE not set, defaulting to https://api.openai.com")
	}
	if openaiModel == "" {
		// Default model if not set
		openaiModel = "gpt-4o" // Or "gpt-3.5-turbo" or another compatible model
		fmt.Printf("Info: OPENAI_MODEL not set, defaulting to %s\n", openaiModel)
	}

	// --- Setup Input ---
	scanner := bufio.NewScanner(os.Stdin)
	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "\u001b[91mError reading input: %v\u001b[0m\n", err)
				return "", false
			}
			return "", false // EOF
		}
		return scanner.Text(), true
	}

	// --- Define Tools ---
	tools := []ToolDefinition{
		ReadFileDefinition,
		ListFilesDefinition,
		EditFileDefinition,
	}

	// --- Create and Run Agent ---
	agent := NewAgent(getUserMessage, tools, openaiModel)
	err := agent.Run(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "\u001b[91mAgent exited with error: %s\u001b[0m\n", err.Error())
		os.Exit(1)
	}
}
