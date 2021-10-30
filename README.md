# RDAP &rarr; WHOIS Proxy  [<img alt="Logo" src="static/favicon.svg" height="96" align="right"/>](https://rdap.redirect2.me/)

[![build](https://github.com/redirect2me/rdap-proxy/actions/workflows/gcr-deploy.yaml/badge.svg)](https://github.com/redirect2me/rdap-proxy/actions/workflows/gcr-deploy.yaml)
[![30 day uptime](https://img.shields.io/nodeping/uptime/akjuezyz-cdli-4wxo-8ay6-6frbmos55ik3.svg?label=30-day%20uptime&style=flat)](https://nodeping.com/reports/uptime/akjuezyz-cdli-4wxo-8ay6-6frbmos55ik3)
[![# of WHOIS servers](https://img.shields.io/badge/dynamic/json.svg?label=WHOIS+Servers&url=https%3A%2F%2Frdap.redirect2.me%2Fstatus.json&query=%24.whoisCount)](https://rdap.redirect2.me/config.json)
[![# of RDAP servers](https://img.shields.io/badge/dynamic/json.svg?label=RDAP+Servers&url=https%3A%2F%2Frdap.redirect2.me%2Fstatus.json&query=%24.rdapCount)](https://rdap.redirect2.me/config.json)


## Resource Links

- [IANA root domain database](https://www.iana.org/domains/root/db) - info for each domain, including (if available), the whois server(s).  This isn't really machine-readable.  Try [rfc1036/whois](https://github.com/rfc1036/whois/blob/next/tld_serv_list) or [whois/ianawhois](https://github.com/whois/ianawhois/blob/master/.code/update.rb).
- [IANA RDAP for domains bootstrap file (JSON)](https://data.iana.org/rdap/dns.json) - the official list of RDAP servers for domain lookups
- [IANA list of RDAP servers for registrars](https://www.iana.org/assignments/registrar-ids/registrar-ids.xhtml) - useful for finding a fallback RDAP server, especially if you buy all your domains from a single registrar.
- [Resolve.rs report of unofficial/missing RDAP servers](https://resolve.rs/domains/rdap-missing.html)

## How to Run

To run locally for development:
```
./run.sh
```

## Contributing

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

golang
google (favicon)

* [echo](https://echo.labstack.com/)
* [rfc1036/whois](https://github.com/rfc1036/whois/blob/next/tld_serv_list) - list of whois servers
* [resolve.rs](https://resolve.rs/domains/rdap.html) - list of rdap servers
<!-- to update:
curl https://resolve.rs/domains/rdap.json\?apikey\=sysadmin+rdap-proxy@redirect2.me | jq --sort-keys . >data/rdap.json
-->
* [viper](https://github.com/spf13/viper)
* [echo](https://echo.labstack.com/guide/)
* [raymond](https://github.com/aymerick/raymond)
* [zerolog](https://github.com/rs/zerolog)

## To Do

- [x] load config
- [x] redirect tlds that already have an RDAP server
- [x] parse requested domain to determine correct WHOIS server
- [x] query WHOIS
- [x] parse response
- [ ] format response as RDAP
- [ ] pass-through error responses

- [x] LICENSE
- [x] README
- [ ] configurable limit message ("Free for light non-commerical use")
- [x] views: html -> hbs
- [ ] /index.html: footer
- [ ] /index.html: note about allowed domains
- [ ] /index.html: link to README for detailed info (after 1st para)
- [x] POST for /index.html
- [ ] list of allowed TLDs to proxy (redirect always works)

- [ ] /config.html
- [x] /config.json
- [ ] remove filesystem access to files in dev mode: air works

- [ ] fallback whois/rdap server for TLDs without
- [x] config via env vars

- [x] status.json
- [ ] metrics
- [ ] status.html

- [ ] embed assets in binary
- [x] favicon: U+1F4C7
- [ ] banner
- [ ] pico
- [ ] home page with info

- [ ] Dockerfile
- [ ] deploy to CloudRun
- [ ] GHA to deploy
- [ ] release: GHCR docker image
- [ ] cron GHA to update whois list
- [ ] cron GHA to update rdap list

- [ ] unify logging
- [ ] disable page logging when running in CloudRun

- [ ] 404/500 pages

- [ ] rate-limiter
- [ ] timeouts
- [ ] time delay (per IP?)
- [ ] flag to override specific whois/redirect
- [ ] flag for port
- [ ] basic auth + config/flag for user/password
- [ ] better logging
- [ ] page to display recent TLDs queried w/success ratio
- [ ] page showing configuration

- [ ] release: binary for Linux
- [ ] release: binary for MacOS
- [ ] release: binary for Windows

- [ ] [graceful shutdown](https://echo.labstack.com/cookbook/graceful-shutdown/)
- [ ] compression
- [ ] security headers

## Potential golang whois libraries

- [undiabler/golang-whois](github.com/undiabler/golang-whois): 6 years old, doesn't parse expiration, can't specify server (but has big list built-in)
- [likexian/whois](https://github.com/likexian/whois/): able to specify server and timeout
- [domainr/whois](https://github.com/domainr/whois): includes parsers for HTTP-based whois servers

- [registrobr/rdap](https://github.com/registrobr/rdap): JSON schema is in `protocol` subdirectory
- [openrdap](https://github.com/openrdap)

- [official whois](https://serverfault.com/questions/343941/how-can-i-find-the-whois-server-for-any-tld) - Server Fault answer with various ways to figure out the correct WHOIS server
