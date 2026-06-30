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

func NewZAdmin(url string, login string, password string, opts ...ZAdminOption) *ZAdmin {
	client := &ZAdmin{
		url:                  buildUrl(url),
		insecure:             false,
		login:                login,
		password:             password,
		debug:                false,
		retryWaitingDuration: time.Second * 5,
		userAgent:            "gosoap",
		timeout:              time.Second * 30,
	}

	for _, applyOpt := range opts {
		applyOpt(client)
	}

	client.buildConnector()

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

type ZAdminOption func(admin *ZAdmin)

func WithInsecureMode() ZAdminOption {
	return func(client *ZAdmin) {
		client.insecure = true
	}
}

func WithDebugMode() ZAdminOption {
	return func(client *ZAdmin) {
		client.debug = true
	}
}

func WithRetryWaitingDuration(duration time.Duration) ZAdminOption {
	return func(client *ZAdmin) {
		client.retryWaitingDuration = duration
	}
}

func WithUserAgent(userAgent string) ZAdminOption {
	return func(client *ZAdmin) {
		client.userAgent = userAgent
	}
}

func WithTimeout(timeout time.Duration) ZAdminOption {
	return func(client *ZAdmin) {
		client.timeout = timeout
	}
}

func WithServerId(serverId string) ZAdminOption {
	return func(client *ZAdmin) {
		client.serverId = serverId
	}
}

type ZAdmin struct {
	Token                *zimbraCommon.Token
	url                  string
	insecure             bool
	login                string
	password             string
	debug                bool
	retryWaitingDuration time.Duration
	userAgent            string
	timeout              time.Duration
	serverId             string
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
	s.Connector.SetHeaderContext(rawToken, s.serverId, s.accountContext)
}

func (s *ZAdmin) buildConnector() {
	s.Connector = s.newConnector()
}

func (s *ZAdmin) newConnector() *zimbraConnector.Connector {
	return zimbraConnector.NewConnector(s.url, s.insecure, s.userAgent, s.debug, s.timeout)
}

func (s *ZAdmin) setAccountContext(id string, name string) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	s.accountContext = &by
	s.Connector.SetHeaderContext(s.GetToken(), s.serverId, s.accountContext)
}

func (s *ZAdmin) setServerId(serverId string) {
	s.serverId = serverId
	s.Connector.SetHeaderContext(s.GetToken(), s.serverId, s.accountContext)
}
