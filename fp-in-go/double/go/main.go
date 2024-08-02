package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// NumberData represents the input data
type NumberData struct {
	numbers []int
}

// ProcessedData represents the processed output data
type ProcessedData struct {
	numbers []int
}

// NewNumberData creates and returns a new NumberData instance
func NewNumberData() *NumberData {
	return &NumberData{numbers: []int{}}
}

// AddNumber adds a number to NumberData
func (nd *NumberData) AddNumber(num int) {
	nd.numbers = append(nd.numbers, num)
}

// Process doubles all numbers in NumberData and returns ProcessedData
func (nd *NumberData) Process() ProcessedData {
	processed := ProcessedData{numbers: make([]int, len(nd.numbers))}
	for i, num := range nd.numbers {
		processed.numbers[i] = num * 2
	}
	return processed
}

// FileProcessor handles file operations and data processing
type FileProcessor struct {
	inputFile  string
	outputFile string
}

// NewFileProcessor creates and returns a new FileProcessor instance
func NewFileProcessor(input, output string) *FileProcessor {
	return &FileProcessor{
		inputFile:  input,
		outputFile: output,
	}
}

// ReadAndDeserialize reads data from input file and deserializes it into NumberData
func (fp *FileProcessor) ReadAndDeserialize() (*NumberData, error) {
	file, err := os.Open(fp.inputFile)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %w", err)
	}
	defer file.Close()

	data := NewNumberData()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error converting to number: %w", err)
		}
		data.AddNumber(num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file: %w", err)
	}

	return data, nil
}

// SerializeAndWrite serializes ProcessedData and writes it to output file
func (fp *FileProcessor) SerializeAndWrite(data ProcessedData) error {
	file, err := os.Create(fp.outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, num := range data.numbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}

	return nil
}

// Process orchestrates the entire data processing workflow
func (fp *FileProcessor) Process() error {
	// Read and deserialize input data
	inputData, err := fp.ReadAndDeserialize()
	if err != nil {
		return err
	}

	// Process data
	processedData := inputData.Process()

	// Serialize and write output data
	err = fp.SerializeAndWrite(processedData)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	processor := NewFileProcessor("input.txt", "output.txt")
	if err := processor.Process(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Processing completed successfully.")
}
