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
	insecure bool,
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
		insecure:             insecure,
		id:                   id,
		name:                 name,
		password:             password,
		domainKey:            domainKey,
		debug:                debug,
		RetryWaitingDuration: retryWaitingDuration,
		UserAgent:            userAgent,
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
	insecure             bool
	id                   string
	name                 string
	password             string
	domainKey            string
	debug                bool
	RetryWaitingDuration time.Duration
	UserAgent            string
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

func (s *ZMailbox) BuildConnector() *zimbraConnector.Connector {
	return zimbraConnector.NewConnector(s.url, s.insecure, s.UserAgent, s.debug, s.timeout)
}

func (s *ZMailbox) BuildConnectorWithContext() *zimbraConnector.Connector {
	connector := zimbraConnector.NewConnector(s.url, s.insecure, s.UserAgent, s.debug, s.timeout)
	connector.SetHeaderContext(s.GetToken(), "", nil)
	return connector
}

func (s *ZMailbox) userAgentContext() *zimbraCommon.NameNode {
	if s.UserAgent == "" {
		return nil
	}

	return &zimbraCommon.NameNode{Name: s.UserAgent}
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
