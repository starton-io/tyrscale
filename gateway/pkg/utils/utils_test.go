package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"testing"
)

func TestGetSHA256Checksum(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write some data to the file
	data := []byte("hello world")
	if _, err := tmpfile.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Calculate the expected checksum for "hello world"
	hash := sha256.Sum256(data)
	expectedChecksum := hex.EncodeToString(hash[:])

	// Test the function
	match, err := GetSHA256Checksum(tmpfile.Name(), expectedChecksum)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !match {
		t.Fatalf("expected checksum to match")
	}
}
