package main

import (
	"fmt"
	"os"
	"strings"
)

// IO represents an IO operation that, when run, produces a value of type any or an error
type IO struct {
	run func() (any, error)
}

// NewIO creates a new IO monad
func NewIO(f func() (any, error)) IO {
	return IO{run: f}
}

// Bind chains IO operations, allowing for type changes
func (io IO) Bind(f func(any) IO) IO {
	return NewIO(func() (any, error) {
		v, err := io.run()
		if err != nil {
			return nil, err
		}
		return f(v).run()
	})
}

// Map transforms the value inside IO
func (io IO) Map(f func(any) any) IO {
	return io.Bind(func(v any) IO {
		return NewIO(func() (any, error) {
			return f(v), nil
		})
	})
}

// Pure lifts a value into the IO context
func Pure(x any) IO {
	return NewIO(func() (any, error) { return x, nil })
}

// ReadFile is an IO operation that reads a file
func ReadFile(filename string) IO {
	return NewIO(func() (any, error) {
		content, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		return string(content), nil
	})
}

// WriteFile is an IO operation that writes to a file
func WriteFile(filename string, content string) IO {
	return NewIO(func() (any, error) {
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to write file: %w", err)
		}
		return true, nil
	})
}

// Print is an IO operation that prints to stdout
func Print(x any) IO {
	return NewIO(func() (any, error) {
		fmt.Println(x)
		return x, nil
	})
}

func main() {
	// Example: Read a file, transform its content, and write it back
	program := ReadFile("input.txt").
		Map(func(v any) any {
			return strings.ToUpper(v.(string))
		}).
		Bind(func(v any) IO {
			return WriteFile("output.txt", v.(string))
		}).
		Bind(func(v any) IO {
			success := v.(bool)
			if success {
				return Pure("File processed successfully")
			}
			return Pure("Failed to process file")
		}).
		Bind(func(v any) IO {
			return Print(v)
		})

	// Run the IO operation
	result, err := program.run()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Program completed: %s\n", result)
	}
}
