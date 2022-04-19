package main

import (
	"fmt"
	"net"
	"strings"
)

type Host struct {
	Host           string
	Address        []string
	MX             []string
	CNAME          string
	NS             []string
	TXT            []string
	NamesByAddress map[string][]string
}

func (h Host) PrettyString() string {
	lines := []string{
		fmt.Sprintf("%s has address %v", h.Host, strings.Join(h.Address, ", ")),
		fmt.Sprintf("  mail is handled by %s", strings.Join(h.MX, ", ")),
		fmt.Sprintf("  CNAME %s", h.CNAME),
		fmt.Sprintf("  DNS NS records %s", strings.Join(h.NS, ", ")),
		fmt.Sprintf("  TXT NS records %s", strings.Join(h.TXT, ", ")),
	}
	for _, a := range h.Address {
		names := h.NamesByAddress[a]
		lines = append(lines, fmt.Sprintf("%s domain name pointer %s", a, strings.Join(names, ", ")))
	}
	return strings.Join(lines, "\n")
}

func LookupHost(host string) (Host, error) {
	address, err := net.LookupHost(host)
	if err != nil {
		return Host{}, fmt.Errorf("lookup host %s: %w", host, err)
	}

	return Host{
		Host:           host,
		MX:             lookupMX(host),
		CNAME:          lookupCNAME(host),
		NS:             lookupNS(host),
		TXT:            lookupTXT(host),
		Address:        address,
		NamesByAddress: lookupAddr(address),
	}, nil
}

func lookupAddr(address []string) map[string][]string {
	namesByAddress := make(map[string][]string)
	for _, a := range address {
		names, err := net.LookupAddr(a)
		if err != nil {
			namesByAddress[a] = []string{err.Error()}
			continue
		}
		namesByAddress[a] = names
	}
	return namesByAddress
}

func lookupMX(host string) []string {
	mx, err := net.LookupMX(host)
	if err != nil {
		return []string{err.Error()}
	}

	var records []string
	for _, record := range mx {
		if record != nil {
			records = append(records, record.Host)
		}
	}
	return records
}

func lookupNS(host string) []string {
	ns, err := net.LookupNS(host)
	if err != nil {
		return []string{err.Error()}
	}

	var records []string
	for _, record := range ns {
		if record != nil {
			records = append(records, record.Host)
		}
	}
	return records
}

func lookupTXT(host string) []string {
	txt, err := net.LookupTXT(host)
	if err != nil {
		return []string{err.Error()}
	}
	return txt
}

func lookupCNAME(host string) string {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		return err.Error()
	}
	return cname
}
