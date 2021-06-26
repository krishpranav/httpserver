package runner

import (
	"github.com/krishpranav/httpserver/pkg/binder"
	"github.com/krishpranav/httpserver/pkg/httpserver"
	"github.com/krishpranav/httpserver/pkg/tcpserver"
	"github.com/projectdiscovery/gologger"
)

type Runner struct {
	options    *Options
	serverTCP  *tcpserver.TCPServer
	httpServer *httpserver.HTTPServer
}

// New instance of runner
func New(options *Options) (*Runner, error) {
	r := Runner{options: options}
	// Check if the process can listen on the specified ip:port
	if !binder.CanListenOn(r.options.ListenAddress) {
		newListenAddress, err := binder.GetRandomListenAddress(r.options.ListenAddress)
		if err != nil {
			return nil, err
		}
		gologger.Print().Msgf("Can't listen on %s: %s - Using %s\n", r.options.ListenAddress, err, newListenAddress)
		r.options.ListenAddress = newListenAddress
	}

	if r.options.EnableTCP {
		serverTCP, err := tcpserver.New(&tcpserver.Options{
			Listen:  r.options.ListenAddress,
			TLS:     r.options.TCPWithTLS,
			Domain:  "local.host",
			Verbose: r.options.Verbose,
		})
		if err != nil {
			return nil, err
		}
		err = serverTCP.LoadTemplate(r.options.RulesFile)
		if err != nil {
			return nil, err
		}
		r.serverTCP = serverTCP
		return &r, nil
	}

	httpServer, err := httpserver.New(&httpserver.Options{
		Folder:            r.options.Folder,
		EnableUpload:      r.options.EnableUpload,
		ListenAddress:     r.options.ListenAddress,
		TLS:               r.options.HTTPS,
		Certificate:       r.options.TLSCertificate,
		CertificateKey:    r.options.TLSKey,
		CertificateDomain: r.options.TLSDomain,
		BasicAuthUsername: r.options.username,
		BasicAuthPassword: r.options.password,
		BasicAuthReal:     r.options.Realm,
		Verbose:           r.options.Verbose,
	})
	if err != nil {
		return nil, err
	}
	r.httpServer = httpServer

	return &r, nil
}

// Run logic
func (r *Runner) Run() error {
	if r.options.EnableTCP {
		gologger.Print().Msgf("Serving TCP rule based server on tcp://%s", r.options.ListenAddress)
		return r.serverTCP.ListenAndServe()
	}

	if r.options.HTTPS {
		gologger.Print().Msgf("Serving %s on https://%s/", r.options.FolderAbsPath(), r.options.ListenAddress)
		return r.httpServer.ListenAndServeTLS()
	}

	gologger.Print().Msgf("Serving %s on http://%s/", r.options.FolderAbsPath(), r.options.ListenAddress)
	return r.httpServer.ListenAndServe()
}

func (r *Runner) Close() error {
	if r.serverTCP != nil {
		if err := r.serverTCP.Close(); err != nil {
			return err
		}
	}
	if r.httpServer != nil {
		if err := r.httpServer.Close(); err != nil {
			return err
		}
	}
	return nil
}
