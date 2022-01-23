package main

import (
	//"encoding/json"
	"time"

	//"github.com/likexian/whois"

	whoisparser "github.com/likexian/whois-parser"
	"github.com/registrobr/rdap/protocol"
)

func convertToRDAP(domain string, whoisInfo whoisparser.WhoisInfo) (int, interface{}) {

	tld := "" //LATER

	retVal := protocol.Domain{}
	retVal.ObjectClassName = "domain"
	retVal.Handle = domain
	retVal.LDHName = domain

	if whoisInfo.Domain.CreatedDate != "" {
		retVal.Events = append(retVal.Events, protocol.Event{
			Action: protocol.EventActionRegistration,
			Date:   parseDate(tld, whoisInfo.Domain.CreatedDate),
		})
	}
	if whoisInfo.Domain.UpdatedDate != "" {
		retVal.Events = append(retVal.Events, protocol.Event{
			Action: protocol.EventActionLastUpdate,
			Date:   parseDate(tld, whoisInfo.Domain.UpdatedDate),
		})
	}
	if whoisInfo.Domain.ExpirationDate != "" {
		retVal.Events = append(retVal.Events, protocol.Event{
			Action: protocol.EventActionExpiration,
			Date:   parseDate(tld, whoisInfo.Domain.ExpirationDate),
		})
	}

	retVal.Links = []protocol.Link{
		{
			Value: "https://rdap.redirect2.me/rdap/domain/" + domain,
			Rel:   "self",
			Href:  "https://rdap.redirect2.me/rdap/domain/" + domain,
			Type:  "application/rdap+json",
		},
	}

	if len(whoisInfo.Domain.NameServers) > 0 {
		for _, ns := range whoisInfo.Domain.NameServers {
			retVal.Nameservers = append(retVal.Nameservers, protocol.Nameserver{
				ObjectClassName: "nameserver",
				LDHName:         ns,
				Status: []protocol.Status{
					"active",
					"associated",
				},
			})
		}
	}

	if len(whoisInfo.Domain.Status) > 0 {
		for _, status := range whoisInfo.Domain.Status {
			retVal.Status = append(retVal.Status, protocol.Status(status))
		}
	}

	/* LATER
	if len(whoisInfo.Registrar.Name) > 0 {
		retVal.Entities = append(retVal.Entities, protocol.Entity{
			ObjectClassName: "entity",
			Handle:          "",
			VCardArray:      []interface{},
		})
	}
	*/

	return 200, retVal
}

func getDateLayout(tld string, input string) string {
	//LATER: lookup by TLD?

	//layouts: https://golang.org/src/time/format.go
	if input[2:3] == "." {
		return "02.01.2006 15:04:05"
	}

	return "2006-01-02T15:04:05Z"
}

func parseDate(tld string, input string) protocol.EventDate {

	t, err := time.Parse(getDateLayout(tld, input), input)
	if err != nil {
		//LATER: log error
		t = time.Time{}
	}
	return protocol.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
