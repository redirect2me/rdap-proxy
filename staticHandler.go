package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed static
var embeddedFiles embed.FS

func staticHandler() http.Handler {
	if devMode {
		log.Print("staticHandler: using file system")
		return http.FileServer(http.FS(os.DirFS("static")))
	}

	log.Print("staticHandler: using embedded")
	fsys, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(fsys))
}
