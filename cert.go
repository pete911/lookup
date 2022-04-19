package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

const certTLSDialTimeout = 5 * time.Second

func SANs(host string) ([]string, error) {
	address := net.JoinHostPort(host, "443")
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: certTLSDialTimeout}, "tcp", address, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, fmt.Errorf("dial %s: %w", address, err)
	}
	defer conn.Close()

	// search only unique dns names
	dnsSet := make(map[string]struct{})
	var dnsNames []string
	for _, cert := range conn.ConnectionState().PeerCertificates {
		for _, dnsName := range cert.DNSNames {
			if _, ok := dnsSet[dnsName]; ok {
				continue
			}
			dnsNames = append(dnsNames, dnsName)
			dnsSet[dnsName] = struct{}{}
		}
	}
	return dnsNames, nil
}
