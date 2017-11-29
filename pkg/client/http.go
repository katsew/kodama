package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type SimpleHTTPServer struct {
	reqUrl string
}

func (s *SimpleHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("Request URL: %s", r.URL.String())
	if r.URL.Path == "/healthz" {
		log.Print("Server status: Healthy")
		w.WriteHeader(http.StatusOK)
		return
	}

	cli := http.Client{Timeout: 10 * time.Second}
	res, err := cli.Get(s.reqUrl)
	if err != nil {
		log.Printf("Failed to get: %s", errors.WithStack(err).Error())
		w.Write([]byte(err.Error()))
		return
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read response body: %s", errors.WithStack(err).Error())
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("Success send response: %s", b)
	resText := fmt.Sprintf("Response sent from server: %s", b)
	w.Write([]byte(resText))

}

type HTTPStrategy struct {
	backendHost string
	backendPort string
}

func (s *HTTPStrategy) RegisterBackend(h string, p string) {
	s.backendHost = h
	s.backendPort = p
}

func (s *HTTPStrategy) Serve(h string, p string) {

	var scheme string
	if s.backendPort == "" {
		scheme = s.backendHost
	} else {
		scheme = fmt.Sprintf("%s:%s", s.backendHost, s.backendPort)
	}
	http.ListenAndServe(fmt.Sprintf("%s:%s", h, p), &SimpleHTTPServer{
		reqUrl: fmt.Sprintf("http://%s", scheme),
	})
}
