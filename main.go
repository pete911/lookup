package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	defaultWhoisAddress = "whois.iana.org"
	defaultWhoisPort    = "43"
)

var Version = "dev"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no host specified, run ")
		os.Exit(0)
	}
	host := args[0]

	printWhois(host)
	printCerts(host)
	printLookup(host)
}

func printWhois(host string) {
	addr, response, err := Whois(defaultWhoisAddress, defaultWhoisPort, host)
	printHeader(fmt.Sprintf("whois %s", addr))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(response))
}

func printCerts(host string) {
	certs, err := GetCerts(host)
	if err != nil {
		printHeader("certs")
		fmt.Println(err.Error())
		return
	}
	printHeader(fmt.Sprintf("certs %s", certs.TLSVersion))
	fmt.Printf("SANs: %s\n", strings.Join(certs.SANs, ", "))
	fmt.Printf("Expiry: %s\n", certs.Expiry)
	fmt.Println()
}

func printLookup(host string) {
	printHeader("lookup")
	response, err := LookupHost(host)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	response.PrettyPrint()
	fmt.Println()
}

func printHeader(msg string) {
	fmt.Printf(" --- [ %s ] ---\n", msg)
}
