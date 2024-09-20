package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// expectedChecksum defines the expected SHA-256 checksums for specific files.
var expectedChecksum = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

// ValidateFileChecksum verifies if the SHA-256 checksum of the given file matches the expected checksum.
// It returns an error if the file doesn't exist or if the checksum verification fails.
func ValidateFileChecksum(file string) error {
	checksum, err := calculateChecksum(file)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%v", err)
		}
		return fmt.Errorf("error calculating checksum: %w", err)
	}

	// Retrieve the expected checksum for the given file from expectedChecksum map.
	expected := expectedChecksum[file]

	if checksum != expected {
		return fmt.Errorf("500: checksum verification failed for: %s, banner file has been modified", file)
	}

	return nil
}

// calculateChecksum computes the SHA-256 checksum of a file given its filename.
// It returns the hexadecimal representation of the checksum and an error if any.
func calculateChecksum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Compute the final hash sum and convert it to hexadecimal string.
	return hex.EncodeToString(hash.Sum(nil)), nil
}
