package zsoap

import (
	"log"
	"time"
)

func (s *ZAdmin) GetAllServers(service string) ([]ZServer, error) {
	req, soapAction := NewGetAllServersRequest(service)
	resp := GetAllServersResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	servers := make([]ZServer, len(resp.Content.Server))

	for index, server := range resp.Content.Server {
		servers[index] = *NewServer(server, s.Client)
	}

	return servers, nil
}

func (s *ZAdmin) GetServer(by ByRequest, applyConfig int, attrs []string) (*ZServer, error) {
	req, soapAction := NewGetServerRequest(by, applyConfig, attrs)
	resp := GetServerResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	server := NewServer(resp.Content.Server[0], s.Client)
	server.Client = s.Client

	return server, nil
}

func (s *ZAdmin) GetServerByName(name string, applyConfig int, attrs []string) (*ZServer, error) {
	by := NewByRequest(NAME_STR, name)
	return s.GetServer(by, applyConfig, attrs)
}

func (s *ZAdmin) GetServerById(id string, applyConfig int, attrs []string) (*ZServer, error) {
	by := NewByRequest(ID_STR, id)
	return s.GetServer(by, applyConfig, attrs)
}

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
