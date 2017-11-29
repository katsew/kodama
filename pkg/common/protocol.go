package common

import "fmt"

type Protocol string

const (
	ProtocolHTTP   = "http"
	ProtocolGrpc   = "grpc"
	ProtocolThrift = "thrift"
)

func (p *Protocol) Validate() bool {
	switch *p {
	case ProtocolHTTP:
		fallthrough
	case ProtocolGrpc:
		fallthrough
	case ProtocolThrift:
		return true
	default:
		return false
	}
}

func (p *Protocol) ValidateError() error {
	return ValidateError{
		Protocol: *p,
	}
}

type ValidateError struct {
	Protocol
}

func (v ValidateError) Error() string {
	return fmt.Sprintf("Failed to validate protocol: %s", v.Protocol)
}
