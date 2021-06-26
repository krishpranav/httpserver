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

func New(options *Options) (*TCPServer, error) {
	return &TCPServer{options: options}, nil
}

func (t *TCPServer) AddRule(rule Rule) error {
	t.options.rules = append(t.options.rules, rule)
	return nil
}

func (t *TCPServer) ListenAndServe() error {
	listener, err := net.Listen("tcp4", t.options.Listen)
	if err != nil {
		return err
	}
	t.listener = listener
	return t.run()
}
