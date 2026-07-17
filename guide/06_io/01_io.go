// I/O (Input / Output)
// In Go, I/O operations are centered around two core interfaces: io.Reader and io.Writer.
// The "os" package provides platform-independent interfaces to operating system functionality,
// while "bufio" implements buffered I/O to optimize reading and writing processes.

package io

import (
	"bufio"
	"fmt"
	"os"
)

// Creating Files
// os.Create creates or truncates (empties) a file.
// It returns a file pointer (*os.File) which acts as an io.ReadWriteCloser.
// Always defer Close() right after a successful open/create to release resources.
func CreatingFiles() {
	file, err := os.Create("resources/example.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
	}
	defer file.Close()
}

// Creating Directories
// os.Mkdir creates a single directory with the specified permissions.
// The parent directory must already exist, or it will return an error.
func CreatingDirectories() {
	err := os.Mkdir("resources/data", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}
}

// Creating Nested Directories
// os.MkdirAll creates a directory along with any necessary parent directories (like mkdir -p).
// If the path already exists, it does nothing and returns nil.
func CreatingNestedDirectories() {
	err := os.MkdirAll("resources/nested/deep/dir", 0755)
	if err != nil {
		fmt.Println("Error creating nested directories:", err)
	}
}

// Removing Files
// os.Remove deletes a single file or an empty directory.
func RemovingFiles() {
	err := os.Remove("resources/example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}
}

// Removing Directories
// os.Remove can delete a directory only if it is completely empty.
// os.RemoveAll removes a path and any children it contains (like rm -rf).
func RemovingDirectories() {
	err := os.Remove("resources/data")
	if err != nil {
		fmt.Println("Error deleting empty directory:", err)
	}

	errAll := os.RemoveAll("resources/nested")
	if errAll != nil {
		fmt.Println("Error deleting directory tree:", errAll)
	}
}

// Reading Files
// os.ReadFile is the simplest way to read an entire file into memory at once.
// It automatically handles opening and closing the file for you.
// Idiomatic for small files where memory footprint is not a concern.
// data is returned as []byte, so we cast it to string to display/use it.
func ReadingFiles() {
	data, err := os.ReadFile("resources/example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(data))
}

// Writing Files
// os.WriteFile writes data to a path. If the file doesn't exist, it creates it.
// The 0644 parameter sets standard unix permissions (Read/Write for owner, Read-only for others).
func WritingFiles() {
	content := "Hello, Go I/O World!"
	err := os.WriteFile("resources/example.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

// Opening Large Files
// os.Open opens the file for reading without loading its entire content into memory.
// bufio.NewScanner reads the file line-by-line, keeping RAM usage low and constant.
func OpeningLargeFiles() {
	file, err := os.Open("large_log.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

// Creating Temp Files
// os.CreateTemp creates a random, unique file in the specified directory.
// Passing an empty string "" targets the default temporary directory of the OS.
func CreatingTempFiles() {
	tempFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer tempFile.Close()

	fmt.Println("Temp file created at:", tempFile.Name())
}

// Removing Temp Files
// Temporary files are not auto-deleted by Go when closed; they must be removed from disk.
// The idiomatic way is to defer os.Remove right after a successful creation.
func RemovingTempFiles() {
	tempFile, err := os.CreateTemp("", "transient-*.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	// This ensures the file is wiped from the disk when the function scope ends
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()
}
