package util

import (
    "fmt"
    "os"
    "strings"
)

const (
    colorReset  = "\033[0m"
    colorRed    = "\033[31m"
    colorGreen  = "\033[32m"
    colorYellow = "\033[33m"
    colorBlue   = "\033[34m"
)

var (
    quietMode bool
)

// SetQuiet enables/disables all output
func SetQuiet(q bool) {
    quietMode = q
}

// LogInfo prints a formatted info message
func LogInfo(format string, args ...interface{}) {
    if quietMode {
        return
    }
    msg := fmt.Sprintf(format, args...)
    fmt.Fprintf(os.Stdout, "%s[INFO]%s %s\n", colorBlue, colorReset, msg)
}

// LogWarning prints a formatted warning message
func LogWarning(format string, args ...interface{}) {
    if quietMode {
        return
    }
    msg := fmt.Sprintf(format, args...)
    fmt.Fprintf(os.Stderr, "%s[WARN]%s %s\n", colorYellow, colorReset, msg)
}

// LogError prints a formatted error message
func LogError(format string, args ...interface{}) {
    if quietMode {
        return
    }
    msg := fmt.Sprintf(format, args...)
    fmt.Fprintf(os.Stderr, "%s[ERROR]%s %s\n", colorRed, colorReset, msg)
}

// PrintBanner prints a stylized ASCII banner
func PrintBanner(text string) {
    if quietMode {
        return
    }
    lines := strings.Split(text, "\n")
    for _, line := range lines {
        fmt.Fprintf(os.Stdout, "%s%s%s\n", colorGreen, line, colorReset)
    }
}