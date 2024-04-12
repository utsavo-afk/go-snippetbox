package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr string
	fs   string
}

var cfg config

func main() {
	// configuration
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// fs := flag.String("static", "./ui/static/", "Path to serve static files from")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.fs, "fs", "./ui/static", "Path to serve static files from")
	flag.Parse()

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir(cfg.fs))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("server started on %s", cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
