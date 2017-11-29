package server

import (
	"fmt"
	"log"
	"net"

	"context"

	"github.com/katsew/kodama/protobuf/ping/pb"
	"google.golang.org/grpc"
)

type GrpcService struct{}

func (s *GrpcService) Send(ctx context.Context, req *pb.Ping) (*pb.Pong, error) {
	return &pb.Pong{}, nil
}

type GrpcStrategy struct{}

func (s *GrpcStrategy) Serve(h string, p string) {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", h, p))
	if err != nil {
		log.Fatal(err)
	}
	svc := GrpcService{}
	svr := grpc.NewServer()
	pb.RegisterPingServiceServer(svr, &svc)
	if err := svr.Serve(lis); err != nil {
		log.Fatal(err)
	}

}

func (s *GrpcStrategy) RegisterBackend(h string, p string) {
	// noop
}
