package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// LATER: move to be embedded
var defaultConfig = []byte(`
whois:
    kr: whois.kr
    io: whois.nic.io
    me: whois.nic.me
    rs: whois.rnids.rs
    sh: whois.nic.sh
redirect:
    ch: https://rdap.nic.ch/
    de: https://rdap.denic.de/
    pr: https://rdap.afilias-srs.net/rdap/pr/
    us: https://rdap.nic.us/
    ve: https://rdap.nic.ve/rdap/
    vu: https://rdap.dnrs.neustar/
    ws: https://rdap.website.ws/
`)

var (
	whoisMap       map[string]string
	redirectMap    map[string]string
	redirectStatus int
	timeout        time.Duration
	port           int
	devMode        bool
	bindHost       string
	allowed        []string
	allowedSet     map[string]bool
)

func loadConfig() {

	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	pflag.BoolVar(&devMode, "dev", false, "Run in development mode")
	pflag.IntVar(&port, "port", 4000, "Port to run on")
	pflag.StringVar(&bindHost, "bind", "", "Network to bind to, usually either localhost (for development) or 0.0.0.0 (default, for production)")
	pflag.StringSliceVar(&allowed, "allowed", []string{}, "List of allowed TLDs")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetDefault("redirectStatus", http.StatusTemporaryRedirect)
	viper.SetDefault("timeout", "10s")
	viper.SetDefault("port", "4000")
	viper.SetDefault("dev", "false")
	viper.SetDefault("bind", "0.0.0.0")

	viper.SetConfigFile("rdap-proxy.yaml")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("config: using default (%s)", err)
			viper.ReadConfig(bytes.NewBuffer(defaultConfig))
		} else {
			log.Fatalf("Unable to load config: %s", err)
		}
	} else {
		log.Printf("config: loaded from %s", viper.ConfigFileUsed())
	}
	whoisMap = viper.GetStringMapString("whois")
	redirectMap = viper.GetStringMapString("redirect")
	redirectStatus = viper.GetInt("redirectStatus")
	var timeoutErr error
	timeout, timeoutErr = time.ParseDuration(viper.GetString("timeout"))
	if timeoutErr != nil {
		timeout = time.Duration(10) * time.Second
	}
	devMode = viper.GetBool("dev")

	allowedSet = make(map[string]bool)
	for allowedTld := range allowed {
		allowedSet[allowed[allowedTld]] = true
	}

	log.Printf("devmode: %s", viper.GetString("dev"))

	//LATER
	//res, _ := json.Marshal(viper.AllSettings())
	//log.Printf("whois: %s", res)
}

// just for debugging
func configHandler(c echo.Context) error {
	return c.JSONPretty(200, viper.AllSettings(), "  ")
}
