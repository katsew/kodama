package client

import (
	"fmt"
	"log"
	"net/http"

	"context"

	"github.com/katsew/kodama/protobuf/ping/pb"
	"google.golang.org/grpc"
)

type GrpcService struct {
	reqScheme string
}

func (s *GrpcService) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("Request URL: %s", r.URL.String())
	if r.URL.Path == "/healthz" {
		log.Print("Server status: Healthy")
		w.WriteHeader(http.StatusOK)
		return
	}

	conn, err := grpc.Dial(s.reqScheme, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	cli := pb.NewPingServiceClient(conn)
	pong, err := cli.Send(context.Background(), &pb.Ping{})
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Failed to handle request: %s", err.Error())))
		return
	}
	log.Printf("Success handle request: %s", pong.String())
	w.Write([]byte{'o', 'k'})

}

type GrpcStrategy struct {
	host string
	port string
}

func (s *GrpcStrategy) RegisterBackend(h string, p string) {
	s.host = h
	s.port = p
}

func (s *GrpcStrategy) Serve(h string, p string) {
	http.ListenAndServe(fmt.Sprintf("%s:%s", h, p), &GrpcService{
		reqScheme: fmt.Sprintf("%s:%s", s.host, s.port),
	})
}
