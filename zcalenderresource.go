package zsoap

import (
	"strings"
	"time"
)

type ZCalendarResource struct {
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

func (a *ZCalendarResource) LastLogon() (time.Time, error) {
	return time.Parse("20060102150405Z", a.ZimbraLastLogonTimestamp)
}

func (a *ZCalendarResource) DomainName() string {
	return strings.Split(a.Name, "@")[1]
}
