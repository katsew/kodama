package main

import (
	"log"

	"os"

	env "github.com/katsew/go-getenv"
	"github.com/katsew/kodama/pkg/client"
	"github.com/katsew/kodama/pkg/common"
	"github.com/katsew/kodama/pkg/server"
)

const (
	LaunchTypeServer = "server"
	LaunchTypeClient = "client"

	Usage = `Usage: kodama [protocol] [launchType]

protocols: http, grpc
launchTypes: server, client

e.g. kodama http server
show this help: kodama help
`
)

func main() {

	if len(os.Args) < 2 {
		os.Stdout.WriteString(Usage)
		return
	}

	t := os.Args[1]
	if t == "help" {
		os.Stdout.WriteString(Usage)
		return
	}
	l := os.Args[2]
	log.Printf("Service: %s %s", t, l)

	pcol := common.Protocol(t)
	if !pcol.Validate() {
		panic(pcol.ValidateError())
	}

	h := env.GetEnv("LAUNCH_HOST", "localhost").String()
	p := env.GetEnv("LAUNCH_PORT", "8080").String()

	var s common.Servable
	switch l {
	case LaunchTypeClient:
		s = client.GetClient()
		s.Use(pcol)
		th := env.GetEnv("BACKEND_HOST", "localhost").String()
		tp := env.GetEnv("BACKEND_PORT", "").String()
		if tp == "" {
			log.Printf("Register backend %s", th)
		} else {
			log.Printf("Register backend %s:%s", th, tp)
		}
		s.RegisterBackend(th, tp)
	case LaunchTypeServer:
		s = server.GetServer()
		s.Use(pcol)
	default:
		panic("You MUST specify launch server type: server/client")
	}
	log.Printf("%s %s running on %s:%s", t, l, h, p)
	s.Serve(h, p)
}
