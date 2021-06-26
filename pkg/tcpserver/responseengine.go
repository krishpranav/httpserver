package tcpserver

import (
	"errors"
)

func (t *TCPServer) BuildResponse(data []byte) ([]byte, error) {

	for _, rule := range t.options.rules {
		if rule.matchRegex.Match(data) {
			return []byte(rule.Response), nil
		}
	}
	return nil, errors.New("no matched rule")
}
