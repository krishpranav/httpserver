package tcpserver

import "regexp"

// RulesConfiguration from yaml
type RulesConfiguration struct {
	Rules []Rule `yaml:"rules"`
}

// Rule to apply to various requests
type Rule struct {
	Match      string `yaml:"match,omitempty"`
	matchRegex *regexp.Regexp
	Response   string `yaml:"response,omitempty"`
}
