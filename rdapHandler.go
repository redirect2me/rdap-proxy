package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func rdapHandler(c echo.Context) error {
	domain := c.Param("domain")
	if domain == "" {
		return c.String(http.StatusNotAcceptable, "No domain specified")
	}
	lastDot := strings.LastIndex(domain, ".")
	if lastDot == -1 {
		return c.String(http.StatusNotAcceptable, "No . in domain")
	}
	if lastDot == len(domain)-1 {
		return c.String(http.StatusNotAcceptable, "Domain cannot end with .")
	}
	tld := domain[lastDot+1:]

	redirect := redirectMap[tld]
	if redirect != "" {
		return c.Redirect(redirectStatus, fmt.Sprintf("%sdomain/%s", redirect, domain))
	}

	whois := whoisMap[tld]
	if whois == "" {
		return c.String(http.StatusOK, fmt.Sprintf("Unable to handle %s (%s)\n", domain, tld))
	}

	return c.String(whoisLookup(whois, domain, c.QueryParam("format")))
}
