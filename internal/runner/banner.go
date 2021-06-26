package runner

import "github.com/projectdiscovery/gologger"

const banner = `SIMPLE HTTP SERVER`

const Version = `1.0.0`

func showBanner() {
	gologger.Print().Msgf("%s\n", banner)

	gologger.Print().Msgf("Use with caution. You are responsible for your actions\n")
	gologger.Print().Msgf("Developers assume no liability and are not responsible for any misuse or damage.\n")
}
