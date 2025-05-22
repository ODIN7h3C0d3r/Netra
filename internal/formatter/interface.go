package formatter

import (
	"fmt"
)

// Format converts IPInfo slice into the specified format (text/json/csv/yaml)
func Format(data []*IPInfo, format, fields string) (string, error) {
	switch format {
	case "csv":
		return FormatCSV(data, fields)
	case "json":
		return FormatJSON(data, fields)
	case "yaml":
		return FormatYAML(data, fields)
	case "text", "":
		return FormatText(data, fields)
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}
