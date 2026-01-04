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
	client := ZAdmin{
		url:                  buildUrl(url),
		insecure:             insecure,
		login:                login,
		password:             password,
		Debug:                debug,
		RetryWaitingDuration: retryWaitingDuration,
		UserAgent:            userAgent,
		Timeout:              timeout,
		ServerId:             serverId,
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
	Debug                bool
	RetryWaitingDuration time.Duration
	UserAgent            string
	Timeout              time.Duration
	ServerId             string
	accountContext       *zimbraCommon.ByNode
	Connector            *zimbraConnector.Connector
}

func (s *ZAdmin) GetToken() string {
	if s.Token == nil {
		return ""
	}

	return s.Token.String()
}

func (s *ZAdmin) SetToken(rawToken string) {
	s.Token = zimbraCommon.NewToken(rawToken)
	s.Connector.SetHeaderContext(rawToken, s.ServerId, s.accountContext)
}

func (s *ZAdmin) BuildConnector() {
	s.Connector = s.NewConnector()
}

func (s *ZAdmin) NewConnector() *zimbraConnector.Connector {
	return zimbraConnector.NewConnector(s.url, s.insecure, s.UserAgent, s.Debug, s.Timeout)
}

func (s *ZAdmin) SetAccountContext(id string, name string) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	s.accountContext = &by
	s.Connector.SetHeaderContext(s.GetToken(), s.ServerId, s.accountContext)
}

func (s *ZAdmin) SetServerId(serverId string) {
	s.ServerId = serverId
	s.Connector.SetHeaderContext(s.GetToken(), s.ServerId, s.accountContext)
}
