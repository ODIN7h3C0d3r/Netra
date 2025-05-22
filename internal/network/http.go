package network

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ODIN7h3C0d3r/Netra/internal/formatter"
)

// FetchIPInfo fetches IP geolocation data from the specified API endpoint
func FetchIPInfo(client *CustomHTTPClient, baseURL, ip string) (*formatter.IPInfo, error) {
	url := strings.TrimSuffix(baseURL, "/") + "/"
	if ip != "" {
		url += ip + "/json/"
	} else {
		url += "json/"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Netra/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := resp.Header.Get("Retry-After")
		return nil, fmt.Errorf("rate limit exceeded. retry after: %s", retryAfter)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var info formatter.IPInfo
	if err := info.FromJSON(body); err != nil {
		return nil, err
	}

	return &info, nil
}
