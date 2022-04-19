package main

import (
	"fmt"
	"net"
	"strings"
)

type Host struct {
	Host    string
	Address []Address
	CNAME   string
	NS      []string
	MX      []string
	TXT     []string
}

type Address struct {
	Address string
	Names   []string
}

func (a Address) String() string {
	if len(a.Names) == 0 {
		return a.Address
	}
	return fmt.Sprintf("%s domain name pointer %s", a.Address, strings.Join(a.Names, ", "))
}

func (h Host) PrettyPrint() {
	fmt.Printf("%s has address\n", h.Host)
	for _, address := range h.Address {
		fmt.Printf("  %s\n", address)
	}
	fmt.Printf("CNAME %s\n", h.CNAME)

	fmt.Println("NS records")
	for _, record := range h.NS {
		fmt.Printf("  %s\n", record)
	}

	fmt.Println("MX records")
	for _, record := range h.MX {
		fmt.Printf("  %s\n", record)
	}

	fmt.Println("TXT records")
	for _, record := range h.TXT {
		fmt.Printf("  %s\n", record)
	}
}

func LookupHost(host string) (Host, error) {
	address, err := net.LookupHost(host)
	if err != nil {
		return Host{}, fmt.Errorf("lookup host %s: %w", host, err)
	}

	return Host{
		Host:    host,
		MX:      lookupMX(host),
		CNAME:   lookupCNAME(host),
		NS:      lookupNS(host),
		TXT:     lookupTXT(host),
		Address: lookupAddr(address),
	}, nil
}

func lookupAddr(address []string) []Address {
	var addresses []Address
	for _, a := range address {
		names, err := net.LookupAddr(a)
		if err != nil {
			addresses = append(addresses, Address{Address: a})
			continue
		}
		addresses = append(addresses, Address{Address: a, Names: names})
	}
	return addresses
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
