package formatter

import (
    "encoding/csv"
    "fmt"
    "strings"
)

// FormatCSV converts IPInfo slice into CSV format with optional field filtering
func FormatCSV(data []*IPInfo, fieldsStr string) (string, error) {
    fields := parseFields(fieldsStr)
    if !validFields(fields) {
        return "", fmt.Errorf("invalid field(s) specified for CSV")
    }

    var b strings.Builder
    writer := csv.NewWriter(&b)

    // Write header
    if err := writer.Write(fields); err != nil {
        return "", err
    }

    // Write rows
    for _, info := range data {
        row := make([]string, len(fields))
        ipMap := info.ToMap()

        for i, field := range fields {
            row[i] = fmt.Sprintf("%v", ipMap[field])
        }

        if err := writer.Write(row); err != nil {
            return "", err
        }
    }

    writer.Flush()
    return b.String(), nil
}