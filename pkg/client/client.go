package client

import (
	"github.com/katsew/kodama/pkg/common"
)

type Client struct {
	common.ServeBase
}

var s *Client

func init() {
	s = new(Client)
	s.SetStrategies(common.StrategyMap{
		common.ProtocolHTTP: &HTTPStrategy{},
		common.ProtocolGrpc: &GrpcStrategy{},
	})
}

func (s *Client) Use(protocol common.Protocol) {
	s.SetCurrent(protocol)
}

func (s *Client) Serve(h string, p string) {
	s.GetCurrent().Serve(h, p)
}

func (s *Client) RegisterBackend(h string, p string) {
	s.GetCurrent().RegisterBackend(h, p)
}

func GetClient() *Client {
	return s
}
