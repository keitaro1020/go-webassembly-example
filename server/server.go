package main

import (
	"flag"
	"net/http"
	"log"
	"regexp"
)

var wasmFile = regexp.MustCompile("\\.wasm$")

func myfileserver(w http.ResponseWriter, r *http.Request) {
	ruri := r.RequestURI
	if wasmFile.MatchString(ruri) {
		w.Header().Set("Content-Type", "application/wasm")
	}
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}

func main() {
	addr := flag.String("a", ":5555", "address:port")
	flag.Parse()
	log.Printf("listening on %q\n", *addr)
	http.HandleFunc("/", myfileserver)
	log.Fatal(http.ListenAndServe(*addr, nil))
}