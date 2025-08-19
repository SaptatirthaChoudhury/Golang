package main

import (
	"fmt"
	"io"
)

// Define small, focused interfaces
type Reader interface {
	Read(data []byte) (int, error)
}

type Writer interface {
	Write(data []byte) (int, error)
}

type Closer interface {
	Close() error
}

// Compose interfaces into a ReadWriteCloser
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// FileProcessor struct  (simulates a file)
type FileProcessor struct {
	name     string
	contents []byte
}

// Implement Reader interface
func (f *FileProcessor) Read(data []byte) (int, error) {
	if len(f.contents) == 0 {
		return 0, io.EOF
	}

	n := copy(data, f.contents)
	f.contents = f.contents[n:] // Simulate reading by removing read data
	return n, nil

}

// Implement Writer interface
func (f *FileProcessor) Write(data []byte) (int, error) {
	f.contents = append(f.contents, data...)
	return len(data), nil
}

// Implement Closer interface
func (f *FileProcessor) Close() error {
	fmt.Printf("Closing %s\n", f.name)
	return nil
}

// Function that uses the composite interface
func ProcessFile(rwc ReadWriteCloser) error {
	// Write data
	_, err := rwc.Write([]byte("Hello, World"))
	if err != nil {
		return err
	}

	// Read data
	buffer := make([]byte, 100)
	n, err := rwc.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}

	fmt.Printf("Read: %s\n", string(buffer[:n]))

	// Close the resource
	return rwc.Close()
}

func main() {
	// Create a FileProcessor instance
	file := &FileProcessor{name: "example.txt", contents: []byte{}}

	// Pass it to ProcessFile (satisfies ReadWriteCloser)
	err := ProcessFile(file)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
