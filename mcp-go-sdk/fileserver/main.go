package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer("filesystem-server", "1.0.0", nil)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	log.Printf("File server serving from directory: %s", pwd)

	// 使用我们自己实现的 File Handler
	handler := createFileHandler(pwd)

	// 添加一个虚构的资源，用于列出目录内容
	server.AddResources(&mcp.ServerResource{
		Resource: &mcp.Resource{
			URI:         "mcp://fs/list",
			Name:        "list_files",
			Description: "List all non-directory files in the current directory.",
		},
		Handler: listDirectoryHandler(pwd),
	})

	// 添加一个资源模板，用于读取指定的文件
	server.AddResourceTemplates(&mcp.ServerResourceTemplate{
		ResourceTemplate: &mcp.ResourceTemplate{
			Name:        "read_file",
			URITemplate: "file:///{+filename}",
			Description: "Read a specific file from the directory. 'filename' is the relative path to the file.",
		},
		Handler: handler,
	})

	log.Println("File system server running over stdio...")
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}

// createFileHandler 是一个简化的、用于演示的 ResourceHandler 工厂函数。
func createFileHandler(baseDir string) mcp.ResourceHandler {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.ReadResourceParams) (*mcp.ReadResourceResult, error) {
		// 注意：在生产环境中，这里必须调用 ss.ListRoots() 来获取客户端授权的
		// 根目录，并进行严格的安全检查。
		// 为了让这个入门示例能用简单的管道命令验证，我们暂时省略了这个双向调用。
		requestedPath := filepath.Join(baseDir, filepath.FromSlash(params.URI[len("file:///"):]))

		data, err := os.ReadFile(requestedPath)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, mcp.ResourceNotFoundError(params.URI)
			}
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{URI: params.URI, MIMEType: "text/plain", Text: string(data)},
			},
		}, nil
	}
}

// listDirectoryHandler 是一个自定义的 ResourceHandler，用于实现列出目录的功能
func listDirectoryHandler(dir string) mcp.ResourceHandler {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.ReadResourceParams) (*mcp.ReadResourceResult, error) {
		// 同样，为简化本地验证，暂时省略对 ss.ListRoots() 的调用。

		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("failed to read directory: %w", err)
		}

		var fileList string
		for _, e := range entries {
			if !e.IsDir() {
				fileList += e.Name() + "\n"
			}
		}
		if fileList == "" {
			fileList = "(The directory is empty or contains no files)"
		}

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{
				{URI: params.URI, MIMEType: "text/plain", Text: fileList},
			},
		}, nil
	}
}
