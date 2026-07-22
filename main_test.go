package main

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCountingFrom1To100(t *testing.T) {
	// Capture stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	
	// Redirect stdout
	oldStdout := os.Stdout
	os.Stdout = w
	
	// Run main
	main()
	
	// Restore stdout
	os.Stdout = oldStdout
	w.Close()
	
	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()
	
	// Verify output
	lines := strings.Split(strings.TrimSpace(output), "\n")
	
	// Should have 100 lines
	if len(lines) != 100 {
		t.Errorf("Expected 100 lines, got %d", len(lines))
	}
	
	// Verify first and last numbers
	if lines[0] != "1" {
		t.Errorf("Expected first line to be '1', got '%s'", lines[0])
	}
	
	if lines[99] != "100" {
		t.Errorf("Expected last line to be '100', got '%s'", lines[99])
	}
	
	// Verify sequence integrity
	for i := 0; i < 100; i++ {
		expected := strconv.Itoa(i + 1)
		if lines[i] != expected {
			t.Errorf("Line %d: expected '%s', got '%s'", i, expected, lines[i])
		}
	}
}
