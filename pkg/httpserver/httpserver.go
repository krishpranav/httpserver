package httpserver

import (
	"net/http"
)

type Options struct {
	Folder            string
	EnableUpload      bool
	ListenAddress     string
	TLS               bool
	Certificate       string
	CertificateKey    string
	CertificateDomain string
	BasicAuthUsername string
	BasicAuthPassword string
	BasicAuthReal     string
	Verbose           string
}

type HTTPServer struct {
	options *Options
	layers  http.Handler
}

func New(options *Options) (*HTTPServer, error) {
	var h HTTPServer
	EnableUpload = options.EnableUpload
	EnableVerbose = options.Verbose
	h.layers = h.loglayer(http.FileServer(http.Dir(options.Folder)))
	if options.BasicAuthUsername != "" || options.BasicAuthPassword != "" {
		h.layers = h.loglayer(h.basicauthlayer(http.FileServer(http.Dir(options.Folder))))
	}
	h.options = options

	return &h, nil
}

func (t *HTTPServer) ListenAndServe() error {
	return http.ListenAndServe(t.options.ListenAddress, t.layers)
}
