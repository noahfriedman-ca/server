package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
}

func Router() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/projects/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(r.URL.Path, http.FileServer(http.Dir("."+r.URL.Path+"/build/"))).ServeHTTP(w, r)
	})
	r.Path("/sitemap.xml").HandlerFunc(serveFile)
	r.Path("/LICENSE").HandlerFunc(serveFile)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/build/"))))

	return r
}
