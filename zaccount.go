package zsoap

import (
	"log"
	"strings"
	"time"
)

func (s *ZAdmin) GetAllAccounts(query string, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZAccount, error) {
	accounts, _, _, _, _, err := s.SearchDirectoryAll(query, domain, applyCos, applyConfig, sortBy, "accounts", sortAscending, attrs)
	return accounts, err
}

func (s *ZAdmin) GetAccounts(query string, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZAccount, error) {
	accounts, _, _, _, _, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, applyCos, applyConfig, sortBy, "accounts", sortAscending, attrs)
	return accounts, err
}

func (s *ZAdmin) GetAccount(byAccount ByRequest, attrs []string) (*ZAccount, error) {

	req, soapAction := NewGetAccountRequest(byAccount, attrs)
	resp := GetAccountResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return NewAccount(resp.Content.Account[0], s.Client), nil
}

func (s *ZAdmin) GetAccountByName(name string, attrs []string) (*ZAccount, error) {
	by := NewByRequest(NAME_STR, name)
	return s.GetAccount(by, attrs)
}

func (s *ZAdmin) GetAccountById(id string, attrs []string) (*ZAccount, error) {
	by := NewByRequest(ID_STR, id)
	return s.GetAccount(by, attrs)
}

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
