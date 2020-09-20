package main

import (
	"log"
	"net/http"
)

type Route struct {
	Endpoint string
	HTMLFile string
}

const (
	Port string = ":80"
)

var (
	Routes []Route = []Route{
		Route{
			Endpoint: "./",
			HTMLFile: "./html/index.html",
		},
		Route{
			Endpoint: "./contributors",
			HTMLFile: "./html/contributors.html",
		},
	}
)

func main() {
	log.Printf("Now listening on port %s", Port)
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(Port, nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path

	for _, r := range Routes {
		if r.Endpoint == p {
			p = r.HTMLFile
			break
		}
	}
	
	http.ServeFile(w, r, p)
}