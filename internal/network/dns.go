package network

import (
    "context"
    "fmt"
    "net"
    "time"
)

// DNSResolver provides custom DNS resolution capabilities
type DNSResolver struct {
    Resolver *net.Resolver
    Timeout  time.Duration
}

// NewDNSResolver creates a new DNS resolver with optional custom DNS servers
func NewDNSResolver(servers []string) *DNSResolver {
    dialer := func(ctx context.Context, network, address string) (net.Conn, error) {
        d := net.Dialer{
            Timeout: 5 * time.Second,
        }
        if len(servers) > 0 {
            address = servers[0]
        }
        return d.DialContext(ctx, network, address)
    }

    return &DNSResolver{
        Resolver: &net.Resolver{
            PreferGo: true,
            Dial:     dialer,
        },
        Timeout: 5 * time.Second,
    }
}

// ResolveHostnameToIP resolves a hostname to an IP address
func (r *DNSResolver) ResolveHostnameToIP(hostname string) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
    defer cancel()

    ips, err := r.Resolver.LookupIP(ctx, "ip", hostname)
    if err != nil {
        return "", err
    }

    if len(ips) == 0 {
        return "", fmt.Errorf("no IP addresses found for %s", hostname)
    }

    return ips[0].String(), nil
}

// ResolveIPToHostname resolves an IP address to a hostname
func (r *DNSResolver) ResolveIPToHostname(ip string) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
    defer cancel()

    names, err := r.Resolver.LookupAddr(ctx, ip)
    if err != nil {
        return "", err
    }

    if len(names) == 0 {
        return "", fmt.Errorf("no hostname found for %s", ip)
    }

    return names[0], nil
}

// GetAllIPs returns all IP addresses (IPv4 and IPv6) for a hostname
func (r *DNSResolver) GetAllIPs(hostname string) ([]string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
    defer cancel()

    ipv4, err4 := r.Resolver.LookupIP(ctx, "ip4", hostname)
    ipv6, err6 := r.Resolver.LookupIP(ctx, "ip6", hostname)

    var ips []string

    if err4 == nil {
        for _, ip := range ipv4 {
            ips = append(ips, ip.String())
        }
    }

    if err6 == nil {
        for _, ip := range ipv6 {
            ips = append(ips, ip.String())
        }
    }

    if len(ips) == 0 {
        return nil, fmt.Errorf("no IPs found for %s", hostname)
    }

    return ips, nil
}