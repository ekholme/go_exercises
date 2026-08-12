package main

import (
	"fmt"
	"net/http"
)

const addr = ":8080"

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", handleIndex)

	s := http.Server{
		Addr:    addr,
		Handler: r,
	}

	fmt.Printf("Running web server on %v", addr)

	s.ListenAndServe()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)

	msg := []byte("Hello World")

	w.Write(msg)
}
