package main

import (
	"github.com/krishpranav/httpserver/internal/runner"
	"github.com/projectdiscovery/gologger"
)

func main() {
	options := runner.ParseOptions()
	r, err := runner.New(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	if err := r.Run(); err != nil {
		gologger.Info().Msgf("%s\n", err)
	}
	defer r.Close()
}
