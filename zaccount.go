package zsoap

import (
	"log"
	"strings"
	"time"
)

type ZAccount struct {
	Client                              *Client
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
}

func (a *ZAccount) LastLogon() (time.Time, error) {
	return time.Parse("20060102150405Z", a.ZimbraLastLogonTimestamp)
}

func (a *ZAccount) Modify(attrs map[string]string) error {

	req, soapAction := NewModifyAccountRequest(a.ID, attrs)
	resp := ModifyAccountResponse{}

	if err := a.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *ZAccount) DomainName() string {
	return strings.Split(a.Name, "@")[1]
}

func NewAccount(resp GenericResponse, client *Client) *ZAccount {
	account := &ZAccount{
		Client: client,
		ID:     resp.ID,
		Name:   resp.Name,
	}

	setResponseAttrs(resp.Attrs, &account)

	return account
}

func NewAccountQuota(resp QuotaResponse, client *Client) *ZAccount {
	account := &ZAccount{
		Client: client,
		ID:     resp.ID,
		Name:   resp.Name,
		Used:   resp.Used,
		Limit:  resp.Limit,
	}

	return account
}

func NewModifyAccountRequest(id string, attrs map[string]string) (*ModifyAccountRequest, string) {

	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &ModifyAccountRequest{
		Content: ModifyAccountRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Attrs: a,
		},
	}
	return r, "urn:zimbraAdmin/ModifyAccount"
}
