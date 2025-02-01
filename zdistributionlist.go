package zsoap

import "strings"

func (s *ZAdmin) GetAllDistributionLists(query string, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, err := s.SearchDirectoryAll(query, domain, applyCos, applyConfig, sortBy, "distributionlists", sortAscending, attrs)
	return dls, err
}

func (s *ZAdmin) GetDistributionLists(query string, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, applyCos, applyConfig, sortBy, "distributionlists", sortAscending, attrs)
	return dls, err
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
