package zmailbox

import (
	"net/url"
	"strings"
	"time"

	"github.com/RaoH37/gosoap/zimbraCommon"
	"github.com/RaoH37/gosoap/zimbraConnector"
)

func NewZMailbox(
	url string,
	tls bool,
	id string,
	name string,
	password string,
	domainKey string,
	debug bool,
	retryWaitingDuration time.Duration,
	userAgent string,
	timeout time.Duration) ZMailbox {
	return ZMailbox{
		url:                  buildUrl(url),
		tls:                  tls,
		id:                   id,
		name:                 name,
		password:             password,
		domainKey:            domainKey,
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
		u.Path = "/service/soap"
	}

	return u.String()
}

type ZMailbox struct {
	Token                *zimbraCommon.Token
	url                  string
	tls                  bool
	id                   string
	name                 string
	password             string
	domainKey            string
	debug                bool
	RetryWaitingDuration time.Duration
	userAgent            string
	timeout              time.Duration
}

func (s *ZMailbox) GetToken() string {
	if s.Token == nil {
		return ""
	}

	return s.Token.String()
}

func (s *ZMailbox) SetToken(rawToken string) {
	s.Token = zimbraCommon.NewToken(rawToken)
}

func (s *ZMailbox) IsTokenValid() bool {
	return s.Token != nil && !s.Token.IsExpired()
}

func (s *ZMailbox) BuildZimbraConnector() *zimbraConnector.Connector {
	return zimbraConnector.BuildConnector(s.url, s.tls, s.userAgent, nil, s.debug, s.timeout)
}

func (s *ZMailbox) BuildZimbraConnectorLogged() *zimbraConnector.Connector {
	connector := zimbraConnector.BuildConnector(s.url, s.tls, s.userAgent, nil, s.debug, s.timeout)
	connector.SetHeaderContext(s.GetToken(), "", "", "")
	return connector
}

func (s *ZMailbox) authRequestByNode() zimbraCommon.ByNode {
	if s.id != "" {
		return zimbraCommon.ByNode{
			By:    zimbraCommon.ID,
			Value: s.id,
		}
	}

	return zimbraCommon.ByNode{
		By:    zimbraCommon.NAME,
		Value: s.name,
	}
}
