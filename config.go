package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	jww "github.com/spf13/jwalterweatherman"
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
)

func loadConfig() {

	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	viper.SetDefault("redirectStatus", http.StatusTemporaryRedirect)
	viper.SetDefault("timeout", "10s")
	viper.SetDefault("port", "4000")

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
	port = viper.GetInt("port")

	//res, _ := json.Marshal(viper.AllSettings())
	//log.Printf("whois: %s", res)
}
