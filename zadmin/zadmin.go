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
	url string,
	insecure bool,
	login string,
	password string,
	debug bool,
	retryWaitingDuration time.Duration,
	userAgent string,
	timeout time.Duration,
	serverId string) ZAdmin {
	return ZAdmin{
		url:                  buildUrl(url),
		insecure:             insecure,
		login:                login,
		password:             password,
		debug:                debug,
		RetryWaitingDuration: retryWaitingDuration,
		UserAgent:            userAgent,
		timeout:              timeout,
		ServerId:             serverId,
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
	Token                *zimbraCommon.Token
	url                  string
	insecure             bool
	login                string
	password             string
	debug                bool
	RetryWaitingDuration time.Duration
	UserAgent            string
	timeout              time.Duration
	ServerId             string
	accountContext       *zimbraCommon.ByNode
}

func (s *ZAdmin) GetToken() string {
	if s.Token == nil {
		return ""
	}

	return s.Token.String()
}

func (s *ZAdmin) SetToken(rawToken string) {
	s.Token = zimbraCommon.NewToken(rawToken)
}

func (s *ZAdmin) BuildConnector() *zimbraConnector.Connector {
	return zimbraConnector.NewConnector(s.url, s.insecure, s.UserAgent, s.debug, s.timeout)
}

func (s *ZAdmin) BuildConnectorWithContext() *zimbraConnector.Connector {
	connector := zimbraConnector.NewConnector(s.url, s.insecure, s.UserAgent, s.debug, s.timeout)
	connector.SetHeaderContext(s.GetToken(), s.ServerId, s.accountContext)
	return connector
}

func (s *ZAdmin) SetAccountContext(id string, name string) {
	by := s.byNode(id, name)
	s.accountContext = &by
}

func (s *ZAdmin) byNode(id string, name string) zimbraCommon.ByNode {
	if id != "" {
		return zimbraCommon.NewByNode(zimbraCommon.ID, id)
	}

	return zimbraCommon.NewByNode(zimbraCommon.NAME, name)
}
