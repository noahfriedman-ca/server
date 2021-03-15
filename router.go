package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
}

func Router() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/{dev:dev|projects}/{project}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)

		prefix := fmt.Sprintf("/%s/%s/", v["dev"], v["project"])

		if r.URL.Path+"/" == prefix {
			http.ServeFile(w, r, "."+prefix+"build/index.html")
		} else {
			http.StripPrefix(prefix, http.FileServer(http.Dir("."+prefix+"build/"))).ServeHTTP(w, r)
		}
	})
	r.Path("/sitemap.xml").HandlerFunc(serveFile)
	r.Path("/LICENSE").HandlerFunc(serveFile)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/build/"))))

	return r
}
