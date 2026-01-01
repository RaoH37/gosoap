package zadmin

import (
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/RaoH37/gosoap/zimbraCommon"
	"github.com/RaoH37/gosoap/zimbraConnector"
)

const Equipment = "Equipment"
const Emplacement = "Emplacement"

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
		urlAdmin:             buildUrl(urlAdmin),
		tls:                  tls,
		login:                login,
		password:             password,
		debug:                debug,
		RetryWaitingDuration: retryWaitingDuration,
		userAgent:            userAgent,
		timeout:              timeout,
	}
}

func buildUrl(inputUrl string) string {
	if !strings.HasPrefix(inputUrl, "http://") && !strings.HasPrefix(inputUrl, "https://") {
		inputUrl = "https://" + inputUrl
	}

	u, err := url.Parse(inputUrl)
	if err != nil {
		return ""
	}

	if u.Path == "" || u.Path == "/" {
		u.Path = "/service/admin/soap"
	}

	if u.Port() == "" {
		u.Host = net.JoinHostPort(u.Hostname(), "7071")
	}

	return u.String()
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

func (s *ZAdmin) byNode(id string, name string) zimbraCommon.ByNode {
	if id != "" {
		return zimbraCommon.NewByNode(zimbraCommon.ID, id)
	}

	return zimbraCommon.NewByNode(zimbraCommon.NAME, name)
}
