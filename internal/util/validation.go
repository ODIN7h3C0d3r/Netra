package util

import (
    "fmt"
    "net"
    "net/url"
    "os"
    "regexp"
    "strings"
)

// IsValidIP validates an IPv4 or IPv6 address
func IsValidIP(ip string) bool {
    return net.ParseIP(ip) != nil
}

// IsValidHostname validates a domain name or hostname
func IsValidHostname(host string) bool {
    if host == "" || len(host) > 253 {
        return false
    }
    parts := strings.Split(host, ".")
    for _, part := range parts {
        if !regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`).MatchString(part) {
            return false
        }
    }
    return true
}

// IsValidURL validates a URL with optional scheme check
func IsValidURL(rawurl string) bool {
    u, err := url.Parse(rawurl)
    if err != nil {
        return false
    }
    return u.Scheme != "" && u.Host != ""
}

// FileExists checks if a file exists
func FileExists(path string) bool {
    info, err := os.Stat(path)
    return err == nil && !info.IsDir()
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
    info, err := os.Stat(path)
    return err == nil && info.IsDir()
}

// IsValidFormat checks if the output format is supported
func IsValidFormat(format string) bool {
    switch strings.ToLower(format) {
    case "text", "json", "csv", "yaml":
        return true
    default:
        return false
    }
}

// IsPrivateIP checks if the IP is private (RFC 1918 and RFC 4193)
func IsPrivateIP(ipStr string) bool {
    ip := net.ParseIP(ipStr)
    if ip == nil {
        return false
    }
    privateRanges := []string{
        "0.0.0.0/8",
        "10.0.0.0/8",
        "172.16.0.0/12",
        "192.168.0.0/16",
        "100.64.0.0/10",
        "::1/128",
        "fc00::/7",
    }
    for _, cidr := range privateRanges {
        _, block, _ := net.ParseCIDR(cidr)
        if block.Contains(ip) {
            return true
        }
    }
    return false
}

// ValidateIPList validates a list of IPs and returns valid ones
func ValidateIPList(ips []string) []string {
    var valid []string
    for _, ip := range ips {
        if IsValidIP(ip) {
            valid = append(valid, ip)
        }
    }
    return valid
}

// ValidateHostOrIP validates either a hostname or an IP address
func ValidateHostOrIP(input string) error {
    if IsValidIP(input) {
        return nil
    }
    if IsValidHostname(input) {
        return nil
    }
    return fmt.Errorf("invalid input: must be a valid IP address or hostname")
}