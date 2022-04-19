package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

const (
	whoisWriteTimeout = 5 * time.Second
	whoisReadTimeout  = 5 * time.Second
)

func Whois(whoisHost, whoisPort, domain string) ([]byte, error) {
	address := net.JoinHostPort(whoisHost, whoisPort)
	fmt.Printf("lookup %s - %s\n", address, domain)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("dial %s: %w", address, err)
	}
	defer conn.Close()

	if err := conn.SetWriteDeadline(time.Now().Add(whoisWriteTimeout)); err != nil {
		return nil, fmt.Errorf("set write deadline: %w", err)
	}
	if err := conn.SetReadDeadline(time.Now().Add(whoisReadTimeout)); err != nil {
		return nil, fmt.Errorf("set read deadline: %w", err)
	}

	request := fmt.Sprintf("%s\r\n", domain)
	if _, err := conn.Write([]byte(request)); err != nil {
		return nil, fmt.Errorf("write request: %w", err)
	}

	b, err := io.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if refer := getRefer(b); refer != "" {
		return Whois(refer, whoisPort, domain)
	}
	return b, nil
}

func getRefer(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "refer:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "refer:"))
		}
	}
	return ""
}
