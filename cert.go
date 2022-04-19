package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

const certTLSDialTimeout = 5 * time.Second

type Certs struct {
	TLSVersion string
	SANs       []string
	Expiry     time.Time
}

func GetCerts(host string) (Certs, error) {
	address := net.JoinHostPort(host, "443")
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: certTLSDialTimeout}, "tcp", address, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return Certs{}, fmt.Errorf("dial %s: %w", address, err)
	}
	defer conn.Close()

	// search only unique dns names
	dnsSet := make(map[string]struct{})
	var dnsNames []string
	var expiry time.Time
	for i, cert := range conn.ConnectionState().PeerCertificates {
		if i == 0 || cert.NotAfter.Before(expiry) {
			expiry = cert.NotAfter
		}
		for _, dnsName := range cert.DNSNames {
			if _, ok := dnsSet[dnsName]; ok {
				continue
			}
			dnsNames = append(dnsNames, dnsName)
			dnsSet[dnsName] = struct{}{}
		}
	}

	return Certs{
		TLSVersion: tlsFormat(conn.ConnectionState().Version),
		SANs:       dnsNames,
		Expiry:     expiry,
	}, nil
}

func tlsFormat(tlsVersion uint16) string {
	switch tlsVersion {
	case 0:
		return ""
	case tls.VersionSSL30:
		return "SSLv3 - Deprecated!"
	case tls.VersionTLS10:
		return "TLS 1.0 - Deprecated!"
	case tls.VersionTLS11:
		return "TLS 1.1 - Deprecated!"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "TLS Version %d (unknown)"
	}
}
