package zsoap

import (
	"log"
	"strings"
	"time"
)

func (s *ZAdmin) GetAllResources(query string, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZResource, error) {
	_, _, _, _, calresources, err := s.SearchDirectoryAll(query, domain, applyCos, applyConfig, sortBy, "resources", sortAscending, attrs)
	return calresources, err
}

func (s *ZAdmin) GetResources(query string, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZResource, error) {
	_, _, _, _, calresources, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, applyCos, applyConfig, sortBy, "resources", sortAscending, attrs)
	return calresources, err
}

func (s *ZAdmin) GetResource(by ByRequest, attrs []string) (*ZResource, error) {

	req, soapAction := NewGetCalendarResourceRequest(by, attrs)
	resp := GetCalendarResourceResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return NewResource(resp.Content.CalResource[0], s.Client), nil
}

func (s *ZAdmin) GetResourceByName(name string, attrs []string) (*ZResource, error) {
	by := NewByRequest(NAME_STR, name)
	return s.GetResource(by, attrs)
}

func (s *ZAdmin) GetResourceById(id string, attrs []string) (*ZResource, error) {
	by := NewByRequest(ID_STR, id)
	return s.GetResource(by, attrs)
}

type ZResource struct {
	Client                           *Client
	ID                               string
	Name                             string
	Used                             int
	Limit                            int
	ZimbraMailHost                   string
	ZimbraMailTransport              string
	ZimbraCOSId                      string
	ZimbraMailStatus                 string
	ZimbraMailQuota                  string
	ZimbraAcccountStatus             string
	ZimbraCalResAutoAcceptDecline    bool
	ZimbraCalResAutoDeclineIfBusy    bool
	ZimbraCalResAutoDeclineRecurring bool
	ZimbraLastLogonTimestamp         string
}

func (a *ZResource) LastLogon() (time.Time, error) {
	return time.Parse("20060102150405Z", a.ZimbraLastLogonTimestamp)
}

func (a *ZResource) Modify(attrs map[string]string) error {

	req, soapAction := NewModifyResourceRequest(a.ID, attrs)
	resp := ModifyCalendarResourceResponse{}

	if err := a.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *ZResource) DomainName() string {
	return strings.Split(a.Name, "@")[1]
}

func NewResource(resp GenericResponse, client *Client) *ZResource {
	resource := &ZResource{
		Client: client,
		ID:     resp.ID,
		Name:   resp.Name,
	}

	setResponseAttrs(resp.Attrs, &resource)

	return resource
}

func NewModifyResourceRequest(id string, attrs map[string]string) (*ModifyCalendarResourceRequest, string) {

	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &ModifyCalendarResourceRequest{
		Content: ModifyCalendarResourceRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Attrs: a,
		},
	}
	return r, "urn:zimbraAdmin/ModifyCalendarResource"
}
