package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	//whoisServerList()
	loadConfig()

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", viewHandler("index.html"))
	e.GET("/contact.html", viewHandler("contact.html"))

	e.GET("/rdap/domain/:domain", rdapHandler)

	theStaticHandler := echo.WrapHandler(staticHandler())
	e.GET("/status.json", echo.WrapHandler(&Status{}))
	e.GET("/robots.txt", theStaticHandler)
	e.GET("/favicon.ico", theStaticHandler)
	e.GET("/favicon.svg", theStaticHandler)
	e.GET("/css/:filename", theStaticHandler)
	e.GET("/js/:filename", theStaticHandler)

	e.GET("/config.json", configHandler)

	// Start server
	e.Logger.Fatal(e.Start(bindHost + ":" + strconv.Itoa(port)))
}
