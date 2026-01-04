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
	client := ZMailbox{
		url:                  buildUrl(url),
		insecure:             insecure,
		id:                   id,
		name:                 name,
		password:             password,
		domainKey:            domainKey,
		Debug:                debug,
		RetryWaitingDuration: retryWaitingDuration,
		UserAgent:            userAgent,
		Timeout:              timeout,
	}

	client.BuildConnector()

	return client
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
	Debug                bool
	RetryWaitingDuration time.Duration
	UserAgent            string
	Timeout              time.Duration
	Connector            *zimbraConnector.Connector
}

func (s *ZMailbox) GetToken() string {
	if s.Token == nil {
		return ""
	}

	return s.Token.String()
}

func (s *ZMailbox) SetToken(rawToken string) {
	s.Token = zimbraCommon.NewToken(rawToken)
	s.Connector.SetHeaderContext(rawToken, "", nil)
}

func (s *ZMailbox) IsTokenValid() bool {
	return s.Token != nil && !s.Token.IsExpired()
}

func (s *ZMailbox) BuildConnector() {
	s.Connector = zimbraConnector.NewConnector(s.url, s.insecure, s.UserAgent, s.Debug, s.Timeout)
}
