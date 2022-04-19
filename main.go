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

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no host specified, run ")
		os.Exit(0)
	}
	host := args[0]

	printWhois(host)
	printSANs(host)
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

func printSANs(host string) {
	tlsVersion, response, err := SANs(host)
	if err != nil {
		printHeader(fmt.Sprintf("SANs %s", host))
		fmt.Println(err.Error())
		return
	}
	printHeader(fmt.Sprintf("SANs %s %s", tlsVersion, host))
	fmt.Println(strings.Join(response, ", "))
	fmt.Println()
}

func printLookup(host string) {
	printHeader("lookup")
	response, err := LookupHost(host)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response.PrettyString())
	fmt.Println()
}

func printHeader(msg string) {
	fmt.Printf(" --- [ %s ] ---\n", msg)
}
