# kodama

Go echo client/server application for testing connectivity of servers(and containers) with ease.

# Docker image

https://cloud.docker.com/swarm/katsew/repository/docker/katsew/kodama/general

# Installation

```
go get github.com/katsew/kodama
```

# Usage

```
kodama [protocol] [launchType]

e.g. kodama http server
```

## Supported protocol

- http
- grpc

## Launch types

- client
- server

# License

MIT
