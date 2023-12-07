package zsoap

import (
	"time"
)

type ZServer struct {
	Client                *Client
	ID                    string
	Name                  string
	ZimbraCreateTimestamp string
	ZimbraServiceEnabled  []string
	ZimbraSmtpHostname    []string
	ZimbraSmtpPort        int
}

func (server *ZServer) CreatedAt() int {
	t, err := time.Parse("20060102150405Z", server.ZimbraCreateTimestamp)

	if err != nil {
		return -1
	}

	return int(t.Unix())
}

func NewServer(resp GenericResponse, client *Client) *ZServer {
	server := &ZServer{
		Client: client,
		ID:     resp.ID,
		Name:   resp.Name,
	}

	setResponseAttrs(resp.Attrs, &server)

	return server
}
