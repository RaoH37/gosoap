package zclient

import "strings"

type ZDistributionList struct {
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
