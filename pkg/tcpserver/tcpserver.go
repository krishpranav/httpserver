package tcpserver

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"time"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/sslcert"
	"gopkg.in/yaml.v2"
)

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

func (t *TCPServer) handleConnection(conn net.Conn) error {
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		if err := conn.SetReadDeadline(time.Now().Add(readTimeout * time.Second)); err != nil {
			gologger.Info().Msgf("%s\n", err)
		}
		_, err := conn.Read(buf)
		if err != nil {
			return err
		}

		gologger.Print().Msgf("%s\n", buf)

		resp, err := t.BuildResponse(buf)
		if err != nil {
			return err
		}

		if _, err := conn.Write(resp); err != nil {
			gologger.Info().Msgf("%s\n", err)
		}

		gologger.Print().Msgf("%s\n", resp)
	}
}

func (t *TCPServer) ListenAndServeTLS() error {
	var tlsConfig *tls.Config
	if t.options.Certificate != "" && t.options.Key != "" {
		cert, err := tls.LoadX509KeyPair(t.options.Certificate, t.options.Key)
		if err != nil {
			return err
		}
		tlsConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	} else {
		tlsOptions := sslcert.DefaultOptions
		tlsOptions.Host = t.options.Domain
		cfg, err := sslcert.NewTLSConfig(tlsOptions)
		if err != nil {
			return err
		}
		tlsConfig = cfg
	}

	listener, err := tls.Listen("tcp", t.options.Listen, tlsConfig)
	if err != nil {
		return err
	}
	t.listener = listener
	return t.run()
}

func (t *TCPServer) run() error {
	for {
		c, err := t.listener.Accept()
		if err != nil {
			return err
		}
		go t.handleConnection(c)
	}
}

func (t *TCPServer) Close() error {
	return t.listener.Close()
}

func (t *TCPServer) LoadTemplate(templatePath string) error {
	var config RulesConfiguration
	yamlFile, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	for _, ruleTemplate := range config.Rules {
		rule, err := NewRule(ruleTemplate.Match, ruleTemplate.Response)
		if err != nil {
			return err
		}
		t.options.rules = append(t.options.rules, *rule)
	}

	return nil
}
