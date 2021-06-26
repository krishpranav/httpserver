package binder

import (
	"net"

	"github.com/projectdiscovery/gologger"
)

func CanListenOn(address string) bool {
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		return false
	}
	if err := listener.Close(); err != nil {
		gologger.Info().Msgf("%s\n", err)
	}
	return true
}
