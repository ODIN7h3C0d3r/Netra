package test

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func binaryPath() string {
	return filepath.Join("..", "test", "netra")
}

func outputFilePath() string {
	return filepath.Join("..", "test", "test_output.txt")
}

func TestNetraHelp(t *testing.T) {
	cmd := exec.Command(binaryPath(), "--help")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run netra --help: %v", err)
	}
	if !strings.Contains(string(output), "Usage: netra") {
		t.Errorf("Help output missing expected usage text. Output: %s", output)
	}
}

func TestNetraVersionFlag(t *testing.T) {
	cmd := exec.Command(binaryPath(), "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run netra --version: %v", err)
	}
	if !strings.Contains(string(output), "Netra v") {
		t.Errorf("Version output missing expected text. Output: %s", output)
	}
}

func TestNetraInvalidIP(t *testing.T) {
	cmd := exec.Command(binaryPath(), "999.999.999.999")
	output, _ := cmd.CombinedOutput()
	if !strings.Contains(string(output), "Skipping invalid IP") {
		t.Errorf("Expected warning for invalid IP, got: %s", output)
	}
}

func TestNetraInteractiveMode(t *testing.T) {
	cmd := exec.Command(binaryPath(), "--interactive")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("Failed to get stdin pipe: %v", err)
	}
	go func() {
		stdin.Write([]byte("exit\n"))
		stdin.Close()
	}()
	output, _ := cmd.CombinedOutput()
	if !strings.Contains(string(output), "Interactive Mode") {
		t.Errorf("Expected interactive mode banner, got: %s", output)
	}
}

func TestNetraOutputFile(t *testing.T) {
	outputFile := outputFilePath()
	cmd := exec.Command(binaryPath(), "8.8.8.8", "-output", outputFile)
	_ = cmd.Run()
	f, err := os.Open(outputFile)
	if err != nil {
		t.Fatalf("Expected output file to be created, error: %v", err)
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	if !strings.Contains(string(data), "8.8.8.8") {
		t.Errorf("Output file missing expected IP info, got: %s", data)
	}
	os.Remove(outputFile)
}
