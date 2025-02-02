package zsoap

import (
	"log"
	"strings"
)

func (s *ZAdmin) GetAllDistributionLists(query string, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, _, err := s.SearchDirectoryAll(query, domain, applyCos, applyConfig, sortBy, "distributionlists", sortAscending, attrs)
	return dls, err
}

func (s *ZAdmin) GetDistributionLists(query string, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, _, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, applyCos, applyConfig, sortBy, "distributionlists", sortAscending, attrs)
	return dls, err
}

func (s *ZAdmin) GetDistributionList(by ByRequest, attrs []string) (*ZDistributionList, error) {

	req, soapAction := NewGetDistributionListRequest(by, attrs)
	resp := GetDistributionListResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return NewDistributionList(resp.Content.Dl[0], s.Client), nil
}

func (s *ZAdmin) GetDistributionListByName(name string, attrs []string) (*ZDistributionList, error) {
	by := NewByRequest(NAME_STR, name)
	return s.GetDistributionList(by, attrs)
}

func (s *ZAdmin) GetDistributionListById(id string, attrs []string) (*ZDistributionList, error) {
	by := NewByRequest(ID_STR, id)
	return s.GetDistributionList(by, attrs)
}

type ZDistributionList struct {
	Client                                             *Client
	ID                                                 string
	Name                                               string
	ZimbraACE                                          []string
	ZimbraCreateTimestamp                              string
	ZimbraMailStatus                                   string
	ZimbraDistributionListSendShareMessageToNewMembers bool
	ZimbraMailHost                                     string
	ZimbraMailForwardingAddress                        []string
	ZimbraMailAlias                                    []string
	ZimbraHideInGal                                    bool
}

func (a *ZDistributionList) DomainName() string {
	return strings.Split(a.Name, "@")[1]
}

func NewDistributionList(resp GenericResponse, client *Client) *ZDistributionList {
	dl := &ZDistributionList{
		Client: client,
		ID:     resp.ID,
		Name:   resp.Name,
	}

	setResponseAttrs(resp.Attrs, &dl)

	return dl
}

func NewAddDistributionListAliasRequest(id string, alias string) (*AddDistributionListAliasRequest, string) {
	r := &AddDistributionListAliasRequest{
		Content: DistributionListAliasRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Alias: alias,
		},
	}
	return r, "urn:zimbraAdmin/AddDistributionListAlias"
}

func (s *ZAdmin) AddDistributionListAlias(id string, alias string) error {
	req, soapAction := NewAddDistributionListAliasRequest(id, alias)

	if err := s.Client.Call(soapAction, req, nil); err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}

func NewRemoveDistributionListAliasRequest(id string, alias string) (*RemoveDistributionListAliasRequest, string) {
	r := &RemoveDistributionListAliasRequest{
		Content: DistributionListAliasRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Alias: alias,
		},
	}
	return r, "urn:zimbraAdmin/RemoveDistributionListAlias"
}

func (s *ZAdmin) RemoveDistributionListAlias(id string, alias string) error {
	req, soapAction := NewRemoveDistributionListAliasRequest(id, alias)

	if err := s.Client.Call(soapAction, req, nil); err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}

func NewAddDistributionListMemberRequest(id string, members []string) (*AddDistributionListMemberRequest, string) {
	r := &AddDistributionListMemberRequest{
		Content: DistributionListMemberRequestContent{
			Urn:     urnAdmin,
			ID:      id,
			Members: convertToContentStrings(members),
		},
	}
	return r, "urn:zimbraAdmin/AddDistributionListMember"
}

func (s *ZAdmin) AddDistributionListMember(id string, members []string) error {
	req, soapAction := NewAddDistributionListMemberRequest(id, members)

	if err := s.Client.Call(soapAction, req, nil); err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}

func NewRemoveDistributionListMemberRequest(id string, members []string) (*RemoveDistributionListMemberRequest, string) {
	r := &RemoveDistributionListMemberRequest{
		Content: DistributionListMemberRequestContent{
			Urn:     urnAdmin,
			ID:      id,
			Members: convertToContentStrings(members),
		},
	}
	return r, "urn:zimbraAdmin/RemoveDistributionListMember"
}

func (s *ZAdmin) RemoveDistributionListMember(id string, members []string) error {
	req, soapAction := NewRemoveDistributionListMemberRequest(id, members)

	if err := s.Client.Call(soapAction, req, nil); err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}
