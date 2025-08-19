package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger interface - defines the behavior
type Logger interface {
	Log(message string)
}

// ConsoleLogger struct - logs to terminal
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[Console] %s: %s\n", time.Now().Format(time.RFC3339), message)
}

// FileLogger struct - logs to a file
type FileLogger struct {
	Filename string
}

func (f FileLogger) Log(message string) {
	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("could not open log file: %v", err)
	}

	defer file.Close()

	logger := log.New(file, "[File]", log.LstdFlags)
	logger.Println(message)
}

// Function that works with ANY Logger
func processSomething(logger Logger) {
	logger.Log("Starting process ...")
	time.Sleep(1 * time.Second) // simulate work
	logger.Log("Proces finished successfully!")
}

func main() {
	// Use console logging
	consoleLogger := ConsoleLogger{}
	processSomething(consoleLogger)

	// Use file logging
	fileLogger := FileLogger{Filename: "app.log"}
	processSomething(fileLogger)

	fmt.Println("Check app.log for file logs")
}
