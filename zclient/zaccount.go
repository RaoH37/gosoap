package zclient

import (
	"strings"
	"time"
)

type ZAccount struct {
	ID                                  string
	Name                                string
	Used                                int
	Limit                               int
	ZimbraMailHost                      string
	ZimbraMailTransport                 string
	ZimbraCOSId                         string
	ZimbraMailStatus                    string
	ZimbraMailQuota                     string
	ZimbraAccountStatus                 string
	ZimbraFeatureMobileSyncEnabled      bool
	ZimbraFeatureMAPIConnectorEnabled   bool
	ZimbraLastLogonTimestamp            string
	ZimbraPrefMailForwardingAddress     []string
	ZimbraMailForwardingAddress         []string
	ZimbraPrefMailLocalDeliveryDisabled bool
	ZimbraMailAlias                     []string
}

func (a *ZAccount) LastLogon() (time.Time, error) {
	return time.Parse("20060102150405Z", a.ZimbraLastLogonTimestamp)
}

func (a *ZAccount) DomainName() string {
	return strings.Split(a.Name, "@")[1]
}
