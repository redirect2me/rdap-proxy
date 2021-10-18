# RDAP &rarr; WHOIS Proxy  [<img alt="Logo" src="static/favicon.svg" height="96" align="right"/>](https://rdap-proxy.redirect2.me/)

## How to Run

To run locally for development:
```
./run.sh
```

## Contributing

## License

## Credits

golang
google (favicon)

* [rfc1036/whois](https://github.com/rfc1036/whois/blob/next/tld_serv_list) - list of whois servers
* [resolve.rs](https://resolve.rs/domains/rdap.html) - list of rdap servers
<!-- to update:
curl https://resolve.rs/domains/rdap.json\?apikey\=sysadmin+rdap-proxy@redirect2.me | jq --sort-keys . >data/rdap.json
-->
* [viper](https://github.com/spf13/viper)
* [echo](https://echo.labstack.com/guide/)

## To Do

- [x] load config
- [x] parse requested domain to determine correct WHOIS server
- [x] query WHOIS
- [x] parse response
- [ ] format response as RDAP
- [ ] pass-through error responses

- [ ] load list of existing RDAP servers
- [x] redirect tlds that already have an RDAP server

- [x] status.json
- [ ] metrics

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
