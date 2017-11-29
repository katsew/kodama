package server

import (
	"github.com/katsew/kodama/pkg/common"
)

type Server struct {
	common.ServeBase
}

var s *Server

func init() {
	s = new(Server)
	s.SetStrategies(common.StrategyMap{
		common.ProtocolHTTP: &HTTPStrategy{},
		common.ProtocolGrpc: &GrpcStrategy{},
	})
}

func (s *Server) Use(p common.Protocol) {
	s.SetCurrent(p)
}

func (s *Server) Serve(h string, p string) {
	str := s.GetCurrent()
	str.Serve(h, p)
}

func (s *Server) RegisterBackend(h string, p string) {
	// noop
}

func GetServer() *Server {
	return s
}
