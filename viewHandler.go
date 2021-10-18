package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"os"

	"github.com/aymerick/raymond"
	"github.com/labstack/echo/v4"
)

//go:embed views
var embeddedViewFS embed.FS

func viewFS() fs.FS {
	if devMode {
		return os.DirFS("views")
	}

	fsys, err := fs.Sub(embeddedViewFS, "views")
	if err != nil {
		panic(err)
	}

	return fsys
}

func getTemplate(view string) string {
	fs := viewFS()

	f, openErr := fs.Open(view)
	if openErr != nil {
		panic(openErr)
	}
	defer f.Close()

	retVal, readErr := io.ReadAll(f)
	if readErr != nil {
		panic(readErr)
	}
	return string(retVal)
}

func viewHandler(view string) func(c echo.Context) error {

	rawTemplate := getTemplate(view)
	parsedTemplate, parseErr := raymond.Parse(rawTemplate)
	if parseErr != nil {
		panic(parseErr)
	}

	return func(c echo.Context) error {
		//content := fmt.Sprintf("from %s", view)
		content, execErr := parsedTemplate.Exec(nil)
		if execErr != nil {
			return c.String(http.StatusInternalServerError, execErr.Error())
		}
		return c.HTML(http.StatusOK, content)
	}
}
