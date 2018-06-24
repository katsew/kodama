# kodama

Simple echo client/server application for testing connectivity of containers(or k8s pods) with ease.

# Docker image

docker pull katsew/kodama:latest

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

# Development

```
go get github.com/katsew/kodama
```

# License

MIT
