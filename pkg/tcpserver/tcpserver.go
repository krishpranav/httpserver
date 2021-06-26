package tcpserver

import "net"

const readTimeout = 5

type Options struct {
	Listen      string
	TLS         bool
	Certificate string
	Key         string
	Domain      string
	rules       []Rule
	Verbose     bool
}

type TCPServer struct {
	options  *Options
	listener net.Listener
}
