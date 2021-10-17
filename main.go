package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	//whoisServerList()
	loadConfig()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/rdap/domain/:domain", rdapHandler)

	theStaticHandler := echo.WrapHandler(staticHandler(true))
	e.GET("/robots.txt", theStaticHandler)
	e.GET("/favicon.ico", theStaticHandler)
	e.GET("/favicon.svg", theStaticHandler)

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}
