package zsoap

import "strings"

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
