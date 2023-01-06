package main

import (
	"fmt"
	"net/http"
)

func rm(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Quickstart:\n\n")
	fmt.Fprintf(w, "GET\n")
	fmt.Fprintf(w, "/readme	    - This Readme!\n")
	fmt.Fprintf(w, "/           - Request IP\n")
	fmt.Fprintf(w, "/agent      - Request User-Agent\n")
}

func ip(w http.ResponseWriter, req *http.Request) {
	ri := req.Header.Get("x-real-ip")
	fmt.Fprintf(w, ri+"\n")
}

func ua(w http.ResponseWriter, req *http.Request) {
	ua := req.Header.Get("User-Agent")
	fmt.Fprintf(w, ua+"\n")
}

func main() {
	fmt.Println("Server listening on 8080/tcp")
	http.HandleFunc("/readme", rm)
	http.HandleFunc("/", ip)
	http.HandleFunc("/agent", ua)
	http.ListenAndServe(":8080", nil)
}
