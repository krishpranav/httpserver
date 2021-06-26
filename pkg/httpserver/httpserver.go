package httpserver

import "net/http"

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
