package formatter

import (
	"encoding/json"
	"fmt"
)

// FormatJSON converts IPInfo slice into JSON format with optional field filtering
func FormatJSON(data []*IPInfo, fieldsStr string) (string, error) {
	fields := parseFields(fieldsStr)
	if !validFields(fields) {
		return "", fmt.Errorf("invalid field(s) specified for JSON")
	}

	var result []map[string]interface{}

	for _, info := range data {
		raw := info.ToMap()

		if len(fields) > 0 {
			filtered := make(map[string]interface{})
			for _, f := range fields {
				filtered[f] = raw[f]
			}
			result = append(result, filtered)
		} else {
			result = append(result, raw)
		}
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
