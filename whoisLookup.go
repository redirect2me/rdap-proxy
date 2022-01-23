package main

import (
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func whoisLookup(server, domain, format string) (int, interface{}) {

	whoisClient := whois.NewClient()

	whoisClient.SetTimeout(timeout)

	raw, err := whoisClient.Whois(domain, server)

	if err != nil {
		return 500, err.Error()
	}

	if format == "raw" {
		return 200, raw
	}

	parsed, parseErr := whoisparser.Parse(raw)
	if parseErr != nil {
		return 500, parseErr.Error()
	}

	if format == "parsed" {
		return 200, parsed
	}

	return convertToRDAP(domain, parsed)
}
