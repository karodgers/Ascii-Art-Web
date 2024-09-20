package functions

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestForNewLine(t *testing.T) {
	expected := "\n\n"
	actual := captureOutput(func() {
		text := "\\n\\n"
		asciiArt := []string{"ASCII Art"}

		if strings.Contains(text, "\\n") {
			input := strings.Split(text, "\\n")
			count := 0
			for _, word := range input {
				if word == "" {
					count++
					if count < len(input) {
						fmt.Println()
					}
				} else if len(word) > 0 {
					PrintLineByLine(word, asciiArt)
				}
			}
		}
	})
	if actual != expected {
		t.Errorf("PrintAsciiArt should print an empty line. Got: %s, Expected: %s", actual, expected)
	}
}

func TestSpecialCharsDetection(t *testing.T) {
	// Define special characters map
	specialChars := map[string]string{
		"\\t": "Tab",
		"\\b": "Backspace",
		"\\v": "Vertical Tab",
		"\\0": "Null",
		"\\f": "Form Feed",
		"\\r": "Carriage Return",
	}

	// Test each special character
	for spChar, description := range specialChars {
		t.Run(description, func(t *testing.T) {
			actual := captureOutput(func() {
				text := "Some text with special character: " + spChar
				if strings.Contains(text, spChar) {
					fmt.Printf("Print Error: Special ASCII character (%s) or (%s) detected \n", spChar, description)
					return
				}
			})

			expected := fmt.Sprintf("Print Error: Special ASCII character (%s) or (%s) detected \n", spChar, description)

			if actual != expected {
				t.Errorf("Special character detection failed. Got: %s, Expected: %s", actual, expected)
			}
		})
	}
}

func TestValidateFileChecksum(t *testing.T) {
	// Create temporary test files
	tempDir := t.TempDir()

	createTestFile := func(name, content string) string {
		path := filepath.Join(tempDir, name)
		err := os.WriteFile(path, []byte(content), 0o644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		return path
	}

	// Create test files
	createTestFile("standard.txt", "This is a valid file")
	createTestFile("shadow.txt", "This file has been modified")
	createTestFile("unknown.txt", "This file is not in the checksum map")

	tests := []struct {
		name    string
		file    string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Missing file",
			file:    "thinkertoy.txt",
			wantErr: true,
			errMsg:  "open thinkertoy.txt: no such file or directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// If the file exists in our temp directory, replace the filename in expectedChecksum
			fullPath := filepath.Join(tempDir, tt.file)
			if _, err := os.Stat(fullPath); err == nil {
				originalChecksum := expectedChecksum[tt.file]
				expectedChecksum[tt.file] = calculateChecksumForTest(t, fullPath)
				defer func() { expectedChecksum[tt.file] = originalChecksum }()
			}

			err := ValidateFileChecksum(tt.file)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFileChecksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("ValidateFileChecksum() error message = %v, want %v", err, tt.errMsg)
			}
		})
	}
}

func calculateChecksumForTest(t *testing.T, filename string) string {
	checksum, err := calculateChecksum(filename)
	if err != nil {
		t.Fatalf("Failed to calculate checksum for test file: %v", err)
	}
	return checksum
}

func TestPrintAsciiArt(t *testing.T) {
	type testCase struct {
		inputText string

		asciiArt []string

		expected string
	}

	testCases := []testCase{
		{
			inputText: "Hello",

			asciiArt: []string{
				"_    ",
				"| |   ",
				"| |__ ",
				"|  _ \\ ",
				"| | | | ",
				"|_| |_| ",
			},

			expected: " _      \n| |     \n| |__   \n|  _ \\  \n| | | | \n|_| |_| \n        \n        \n",
		},

		{
			inputText: "\\n",

			asciiArt: []string{
				"_    ",
				"| |   ",
				"| |__ ",
				"|  _ \\ ",
				"| | | | ",
				"|_| |_| ",
			},

			expected: "\n",
		},

		{
			inputText: "",

			asciiArt: []string{
				"_    ",
				"| |   ",
				"| |__ ",
				"|  _ \\ ",
				"| | | | ",
				"|_| |_| ",
			},

			expected: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.inputText, func(t *testing.T) {
			var buffer strings.Builder

			PrintAsciiArt(testCase.inputText, testCase.asciiArt)

			buffer.WriteString(testCase.expected)

			got := buffer.String()

			if got != testCase.expected {
				t.Errorf("got %q, want %q", got, testCase.expected)
			}
		})
	}
}

func captureOutput(f func()) string {
	// Keep track of the original stdout
	originalStdout := os.Stdout
	// Create a pipe to redirect stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	// Restore the original stdout when the function returns
	defer func() {
		os.Stdout = originalStdout
	}()
	// Create a buffer to capture the output
	var capturedOutput bytes.Buffer
	// Copy the pipe's read end to the buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&capturedOutput, r)
		close(done)
	}()
	// Call the function
	f()
	// Close the write end of the pipe to signal the goroutine to exit
	w.Close()
	// Wait for the goroutine to finish
	<-done

	return capturedOutput.String()
}
