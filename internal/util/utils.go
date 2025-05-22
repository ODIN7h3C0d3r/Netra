package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// ReadLines reads a file and returns its lines as a slice
func ReadLines(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var result []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}
	return result, nil
}

// SaveToFile writes content to a file
func SaveToFile(path, content string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(content), 0644)
}

// Truncate returns a truncated string with ellipsis if needed
func Truncate(s string, maxLen int) string {
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}
	return string([]rune(s)[:maxLen-1]) + "…"
}

// TitleCase converts snake_case or camelCase to Title Case
func TitleCase(s string) string {
	words := strings.Split(strings.ReplaceAll(s, "_", " "), " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// MapKeys returns sorted keys of a map
func MapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// ExpandHome replaces ~ with user home directory
func ExpandHome(path string) string {
	if strings.HasPrefix(path, "~") {
		home, _ := os.UserHomeDir()
		return home + path[1:]
	}
	return path
}

// IsEmpty checks if a string is empty after trimming
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
