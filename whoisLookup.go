package main

import (
	"encoding/json"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func whoisLookup(server, domain, format string) (int, string) {

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
		jsonStr, jsonErr := json.MarshalIndent(parsed, "", "  ")
		if jsonErr != nil {
			return 500, jsonErr.Error()
		}

		return 200, string(jsonStr)
	}

	converted, convertErr := convertToRDAP(parsed)
	if convertErr != nil {
		return 500, convertErr.Error()
	}

	return 200, converted
}
