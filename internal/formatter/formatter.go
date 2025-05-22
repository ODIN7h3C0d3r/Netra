package formatter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// IPInfo represents structured geolocation and network metadata for an IP
// Moved from core/ipinfo.go to avoid import cycles
// Helper functions also moved here
type IPInfo struct {
	IP        string  `json:"ip"`
	Country   string  `json:"country_name"`
	Region    string  `json:"region"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	ISP       string  `json:"organization"`
	Postal    string  `json:"postal"`
	ASN       string  `json:"asn"`
	IsMobile  bool    `json:"device_is_mobile"`
	IsProxy   bool    `json:"proxy"`
	IsHosting bool    `json:"is_hosting"`
	Org       string  `json:"organization"`
	Continent string  `json:"continent"`
}

func (i *IPInfo) FromJSON(data []byte) error {
	type alias IPInfo
	aux := &struct {
		*alias
	}{
		alias: (*alias)(i),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to parse IP info JSON: %v", err)
	}

	// Post-processing
	i.IsHosting = detectHosting(i.ISP, i.ASN)
	return nil
}

func (i *IPInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ip":         i.IP,
		"country":    i.Country,
		"region":     i.Region,
		"city":       i.City,
		"latitude":   i.Latitude,
		"longitude":  i.Longitude,
		"timezone":   i.Timezone,
		"isp":        i.ISP,
		"postal":     i.Postal,
		"asn":        i.ASN,
		"is_mobile":  boolToString(i.IsMobile),
		"is_proxy":   boolToString(i.IsProxy),
		"is_hosting": boolToString(i.IsHosting),
		"org":        i.Org,
		"continent":  i.Continent,
	}
}

// Helper functions
func parseFields(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(strings.ToLower(s), ",")
}

func boolToString(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

func titleCase(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func detectHosting(isp, asn string) bool {
	hostingKeywords := []string{
		"Amazon", "AWS", "Azure", "DigitalOcean", "Linode",
		"Google", "Cloudflare", "Cloud", "Vultr", "Scaleway",
		"OVH", "Hetzner", "Namecheap", "Fastly",
	}

	for _, keyword := range hostingKeywords {
		if strings.Contains(strings.ToLower(isp), strings.ToLower(keyword)) {
			return true
		}
	}
	return false
}

var validFieldMap = map[string]bool{
	"ip":         true,
	"country":    true,
	"region":     true,
	"city":       true,
	"latitude":   true,
	"longitude":  true,
	"timezone":   true,
	"isp":        true,
	"postal":     true,
	"asn":        true,
	"is_mobile":  true,
	"is_proxy":   true,
	"is_hosting": true,
}

func validFields(fields []string) bool {
	for _, f := range fields {
		if !validFieldMap[f] && f != "" {
			return false
		}
	}
	return true
}

func getAllFields() []string {
	fields := make([]string, 0, len(validFieldMap))
	for k := range validFieldMap {
		fields = append(fields, k)
	}
	return fields
}
