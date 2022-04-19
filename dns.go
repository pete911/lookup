package main

import (
	"fmt"
	"net"
	"strings"
)

type Host struct {
	Host           string
	Address        []string
	Mx             []string
	NamesByAddress map[string][]string
}

func (h Host) PrettyString() string {
	lines := []string{
		fmt.Sprintf("%s has address %v", h.Host, strings.Join(h.Address, ", ")),
		fmt.Sprintf("%s mail is handled by %s", h.Host, strings.Join(h.Mx, ", ")),
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
		Mx:             lookupMX(host),
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
		if record == nil {
			continue
		}
		records = append(records, record.Host)
	}
	return records
}
