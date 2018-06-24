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
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.URL.Path == "/panic" {
		log.Panic("client not sent request, panic occurred, auto recover client")
	}
	if r.URL.Path == "/fatal" {
		log.Fatal("client not sent request, fatal error occurred, shutdown client")
	}

	cli := http.Client{Timeout: 10 * time.Second}

	var res *http.Response
	var err error
	var reqUrl = s.reqUrl

	switch r.URL.Path {
	case "/server/healthz":
		log.Print("check backend server health")
		reqUrl = fmt.Sprintf("%s/%s", reqUrl, "healthz")
	case "/server/fatal":
		log.Print("make backend server fatal")
		reqUrl = fmt.Sprintf("%s/%s", reqUrl, "fatal")
	case "/server/panic":
		log.Print("make backend server panic")
		reqUrl = fmt.Sprintf("%s/%s", reqUrl, "panic")
	default:
		reqUrl = fmt.Sprintf("%s%s", reqUrl, r.URL.Path)
	}
	res, err = cli.Get(reqUrl)
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

	resText := fmt.Sprintf("Receive response from server: %s", b)
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
