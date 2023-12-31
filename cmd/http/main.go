package main

import (
	"flag"
	"log"
	"net/http"

	"acsq.me/dblog"
)

var (
	langs  = []string{"en-US", "es-US", "de-DE"}
	scheme string
)

func main() {
	if err := dblog.MakeDB(); err != nil {
		log.Fatal(err)
	}

	addr := flag.String("addr", ":4000", "HTTP Network Address")
	https := flag.Bool("https", false, "TLS Encryption")
	flag.Parse() // required before flag is used

	if *https {
		scheme = "https"
	} else {
		scheme = "http"
	}

	mux := http.NewServeMux()

	// TODO: Make cooler router
	mux.HandleFunc("/", pageHandler)
	for _, lang := range langs {
		mux.HandleFunc(translatePath(lang, "/posts"), pageHandler)
		mux.HandleFunc(translatePath(lang, "/posts/"), postHandler)
		mux.HandleFunc(translatePath(lang, "/tags/"), tagHandler)
		mux.HandleFunc(translatePath(lang, "/recommend"), recommendHandler)
		mux.HandleFunc(translatePath(lang, "/recommend/"), recommendHandler)
	}
	mux.HandleFunc("/favicon", redirectHandler("/favicon.ico"))
	mux.HandleFunc("/favicon.ico", faviconHandler)
	mux.HandleFunc("/cv", redirectHandler("/cv.pdf"))
	mux.HandleFunc("/cv.pdf", cvHandler)
	mux.HandleFunc("/pgp", redirectHandler("/angelcastaneda.asc"))
	mux.HandleFunc("/angelcastaneda.asc", pgpHandler)
	mux.HandleFunc("/atom.xml", feedHandler)
	mux.HandleFunc("/submit", apiHandler)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting the server on %s", *addr)

	err := http.ListenAndServe(*addr, gzipHandler(redirectWWW(mux)))
	log.Fatal(err)
}
