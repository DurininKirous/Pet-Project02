package scanner_test

import (
	"project02/internal/scanner"
	"os"
	"testing"
	"path/filepath"
)

func TestScanDir(t *testing.T) {
	tmpDir := t.TempDir()
	smallFile := filepath.Join(tmpDir, "small.txt")
	bigFile := filepath.Join(tmpDir, "big.txt")

	err := os.WriteFile(smallFile, []byte("12345"), 0644) 
	if err != nil {
		t.Fatalf("failed to create small file: %v", err)
	}

	err = os.WriteFile(bigFile, []byte("1234567890abcdef"), 0644) 
	if err != nil {
		t.Fatalf("failed to create big file: %v", err)
	}

	result := scanner.ScanDir(tmpDir, 0)
	
	if result.Path != bigFile {
		t.Errorf("expected biggest file to be %s, got %s", bigFile, result.Path)
	}
}
