package zadmin

import (
	"time"

	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraConnector"
)

func NewZAdmin(
	urlAdmin string,
	tls bool,
	login string,
	password string,
	debug bool,
	retryWaitingDuration time.Duration,
	userAgent string,
	timeout time.Duration) ZAdmin {
	return ZAdmin{
		urlAdmin:             urlAdmin,
		tls:                  tls,
		login:                login,
		password:             password,
		debug:                debug,
		RetryWaitingDuration: retryWaitingDuration,
		userAgent:            userAgent,
		timeout:              timeout,
	}
}

type ZAdmin struct {
	AuthToken            string
	urlAdmin             string
	tls                  bool
	login                string
	password             string
	debug                bool
	RetryWaitingDuration time.Duration
	userAgent            string
	timeout              time.Duration
}

func (s *ZAdmin) buildZimbraConnector() *zimbraConnector.Connector {
	return zimbraConnector.BuildConnector(s.urlAdmin, s.tls, s.userAgent, nil, s.debug, s.timeout)
}

func (s *ZAdmin) byNode(id string, name string) zimbraAdmin.ByNode {
	if id != "" {
		return zimbraAdmin.NewByNode("id", id)
	}

	return zimbraAdmin.NewByNode("name", name)
}
