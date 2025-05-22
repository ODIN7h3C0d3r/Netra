package formatter

import (
    "fmt"
    "strings"
)

// FormatText converts IPInfo slice into human-readable text with optional field filtering
func FormatText(data []*IPInfo, fieldsStr string) (string, error) {
    fields := parseFields(fieldsStr)
    if !validFields(fields) {
        return "", fmt.Errorf("invalid field(s) specified for text format")
    }

    var b strings.Builder

    for i, info := range data {
        ipMap := info.ToMap()

        if len(fields) == 0 {
            fields = getAllFields()
        }

        for _, f := range fields {
            b.WriteString(fmt.Sprintf("%s: %v\n", titleCase(f), ipMap[f]))
        }

        if i < len(data)-1 {
            b.WriteString("\n---\n\n")
        }
    }

    return b.String(), nil
}