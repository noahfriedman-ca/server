package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strings"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
}

func Router() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/projects/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		re := regexp.MustCompile("^/projects/([^/]*)/(.*)$")
		m := re.FindStringSubmatch(r.URL.Path)

		if m[2] == "" {
			var trailingSlash string
			if !strings.HasSuffix(r.URL.Path, "/") {
				trailingSlash = "/"
			}

			http.ServeFile(w, r, "."+r.URL.Path+trailingSlash+"build/index.html")
		} else {
			http.ServeFile(w, r, fmt.Sprintf("./projects/%s/build/%s", m[1], m[2]))
		}
	})
	r.Path("/sitemap.xml").HandlerFunc(serveFile)
	r.Path("/LICENSE").HandlerFunc(serveFile)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/build/"))))

	return r
}
