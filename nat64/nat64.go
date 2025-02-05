package nat64

import (
    "fmt"
    "net"
)

func IPv6toIPv4(ipv6Addr string) (string, error) {
    prefix := "64:ff9b::"
    ip := net.ParseIP(ipv6Addr)

    if ip == nil {
        return "", fmt.Errorf("invalid IPv6 address")
    }

    ipv4Addr := fmt.Sprintf("%d.%d.%d.%d", ip[12], ip[13], ip[14], ip[15])
    return prefix + ipv4Addr, nil
}

func IPv4toIPv6(ipv4Addr string) (string, error) {
    ip := net.ParseIP(ipv4Addr)

    if ip == nil {
        return "", fmt.Errorf("invalid IPv4 address")
    }

    ipv6Addr := fmt.Sprintf("64:ff9b::%02x%02x%02x%02x", ip[12], ip[13], ip[14], ip[15])
    return ipv6Addr, nil
}
