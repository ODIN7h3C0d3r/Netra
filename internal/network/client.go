package network

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"
)

// HTTPClientConfig holds configurable options for the HTTP client
type HTTPClientConfig struct {
	Timeout      time.Duration
	RetryLimit   int
	ProxyURL     string
	UserAgent    string
	MaxIdleConns int
}

// CustomHTTPClient wraps http.Client with enhanced capabilities
type CustomHTTPClient struct {
	client *http.Client
	cfg    HTTPClientConfig
}

// NewCustomHTTPClient creates a new HTTP client with customizable options
func NewCustomHTTPClient(cfg HTTPClientConfig) (*CustomHTTPClient, error) {
	// Set default values if not provided
	if cfg.Timeout == 0 {
		cfg.Timeout = 10 * time.Second
	}
	if cfg.RetryLimit == 0 {
		cfg.RetryLimit = 3
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 100
	}

	transport := &http.Transport{
		Proxy: nil,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost: cfg.MaxIdleConns,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	if cfg.ProxyURL != "" {
		proxy, err := ParseProxyURL(cfg.ProxyURL)
		if err != nil {
			return nil, err
		}
		transport.Proxy = proxy
	}

	return &CustomHTTPClient{
		client: &http.Client{
			Transport: transport,
			Timeout:   cfg.Timeout,
		},
		cfg: cfg,
	}, nil
}

// Do executes a request with retry logic and context
func (c *CustomHTTPClient) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt < c.cfg.RetryLimit; attempt++ {
		req = req.Clone(context.Background())
		resp, err = c.client.Do(req)
		if err == nil {
			if resp.StatusCode < 500 || resp.StatusCode == 429 {
				break
			}
		}

		if attempt < c.cfg.RetryLimit-1 {
			time.Sleep(time.Second * time.Duration(attempt+1))
		}
	}

	return resp, err
}

// ParseProxyURL parses and validates a proxy URL
func ParseProxyURL(proxy string) (func(*http.Request) (*url.URL, error), error) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}
	return http.ProxyURL(proxyURL), nil
}
