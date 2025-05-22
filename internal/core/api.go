package core

import (
	"fmt"
	"time"

	"github.com/ODIN7h3C0d3r/Netra/internal/formatter"
	"github.com/ODIN7h3C0d3r/Netra/internal/network"
	"github.com/ODIN7h3C0d3r/Netra/internal/util"
)

const (
	DefaultAPIBaseURL = "https://ipapi.co"
	RetryBackoffTime  = 3 * time.Second
	MaxRetries        = 3
	CacheTTL          = 24 * time.Hour
)

var (
	cache = NewIPInfoCache(CacheTTL)
)

// GetIPInfo fetches IP information from API or cache
func GetIPInfo(ip string) (*formatter.IPInfo, error) {
	// Check cache first
	if cached, ok := cache.Get(ip); ok {
		util.LogInfo("Using cached result for %s", ip)
		return cached, nil
	}

	// Rate limit check
	if cache.AttemptCount(ip) >= MaxRetries {
		util.LogWarning("Too many failed attempts for %s. Skipping request.", ip)
		return nil, fmt.Errorf("too many failed attempts")
	}

	cfg := network.HTTPClientConfig{
		Timeout:    10 * time.Second,
		RetryLimit: 3,
		UserAgent:  "Netra/1.0 (+https://github.com/ODIN7h3C0d3r/Netra )",
	}

	client, err := network.NewCustomHTTPClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %v", err)
	}

	var result *formatter.IPInfo
	var fetchErr error

	// Attempt retries with exponential backoff
	for attempt := 1; attempt <= MaxRetries; attempt++ {
		result, fetchErr = network.FetchIPInfo(client, DefaultAPIBaseURL, ip)
		if fetchErr == nil {
			break
		}

		util.LogWarning("Attempt %d failed for %s: %v", attempt, ip, fetchErr)
		cache.RecordAttempt(ip)

		if attempt < MaxRetries {
			time.Sleep(RetryBackoffTime * time.Duration(attempt))
		}
	}

	if fetchErr != nil {
		return nil, fetchErr
	}

	// Cache successful result
	cache.Set(ip, result)
	return result, nil
}
