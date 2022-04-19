package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no domain specified")
		os.Exit(0)
	}

	b, err := Whois("whois.iana.org", "43", args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(b))

	dnsNames, err := SANs(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("SANs: %s\n", strings.Join(dnsNames, ", "))

	host, err := LookupHost(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(host.PrettyString())
}
