package network

import (
	"net/http"

	"github.com/ODIN7h3C0d3r/Netra/internal/formatter"
)

// HTTPClient defines the interface for making HTTP requests
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// FetchIPInfoFunc is a function signature for fetching IP info
type FetchIPInfoFunc func(client HTTPClient, baseURL, ip string) (*formatter.IPInfo, error)
