package formatter

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// FormatYAML converts IPInfo slice into YAML format with optional field filtering
func FormatYAML(data []*IPInfo, fieldsStr string) (string, error) {
	fields := parseFields(fieldsStr)
	if !validFields(fields) {
		return "", fmt.Errorf("invalid field(s) specified for YAML")
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

	yamlData, err := yaml.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}
