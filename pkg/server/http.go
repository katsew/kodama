package server

import (
	"fmt"
	"net/http"

	"log"
)

type SimpleHTTPServer struct{}

func (s *SimpleHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("Requested URL: %s", r.URL.String())
	if r.URL.Path == "/healthz" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.URL.Path == "/panic" {
		log.Panic("panic occurred, auto recover server")
	}
	if r.URL.Path == "/fatal" {
		log.Fatal("fatal error occurred, shutdown server")
	}

	log.Printf("Got request, send back to client.")
	w.Write([]byte("ok"))
}

type HTTPStrategy struct{}

func (s *HTTPStrategy) Serve(h string, p string) {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", h, p), &SimpleHTTPServer{})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *HTTPStrategy) RegisterBackend(h string, p string) {
	// noop
}
