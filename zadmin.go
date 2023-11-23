package zsoap

import (
	"log"
)

const NAME_STR = "name"
const ID_STR = "id"

type ZAdmin struct {
	AuthToken string
	Client    *Client
}

func (s *ZAdmin) Init(url string, isTLS bool) {
	s.Client = NewClient(url, isTLS, nil)
}

func (s *ZAdmin) Login(name string, password string) {

	req, soapAction := NewAuthRequest(name, password)
	resp := AuthResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	s.Client.SetToken(resp.Content.TOKEN[0].Content)
}

func (s *ZAdmin) Debug() {
	s.Client.Debug = true
}

func (s *ZAdmin) GetAccount(byAccount ByRequest, attrs []string) *ZAccount {

	req, soapAction := NewGetAccountRequest(byAccount, attrs)
	resp := GetAccountResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	return NewAccount(resp.Content.Account[0], s.Client)
}

func (s *ZAdmin) GetAccountByName(name string, attrs []string) *ZAccount {
	by := NewByRequest(NAME_STR, name)
	return s.GetAccount(by, attrs)
}

func (s *ZAdmin) GetAccountById(id string, attrs []string) *ZAccount {
	by := NewByRequest(ID_STR, id)
	return s.GetAccount(by, attrs)
}

// func (s *ZAdmin) GetAllAccounts(server *ByRequest, domain *ByRequest) []ZAccount {

// 	req, soapAction := NewGetAllAccountsRequest(server, domain)
// 	resp := GetAllAccountsResponse{}

// 	if err := s.Client.Call(soapAction, req, &resp); err != nil {
// 		log.Fatal(err)
// 	}

// 	accounts := make([]ZAccount, len(resp.Content.Account))

// 	for index, account := range resp.Content.Account {
// 		accounts[index] = *NewAccount(account, s.Client)
// 	}

// 	return accounts
// }

func (s *ZAdmin) GetAllCos() []ZCos {
	req, soapAction := NewGetAllCosRequest()
	resp := GetAllCosResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	coses := make([]ZCos, len(resp.Content.Cos))

	for index, cos := range resp.Content.Cos {
		coses[index] = *NewCos(cos, s.Client)
	}

	return coses
}

func (s *ZAdmin) GetAllDomains() []ZDomain {
	req, soapAction := NewGetAllDomainsRequest()
	resp := GetAllDomainsResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	domains := make([]ZDomain, len(resp.Content.Domain))

	for index, domain := range resp.Content.Domain {
		domains[index] = *NewDomain(domain, s.Client)
	}

	return domains
}

func (s *ZAdmin) GetDomain(by ByRequest, attrs []string) *ZDomain {
	req, soapAction := NewGetDomainRequest(by, attrs)
	resp := GetDomainResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	domain := NewDomain(resp.Content.Domain[0], s.Client)
	domain.Client = s.Client

	return domain
}

func (s *ZAdmin) GetDomainByName(name string, attrs []string) *ZDomain {
	by := NewByRequest(NAME_STR, name)
	return s.GetDomain(by, attrs)
}

func (s *ZAdmin) GetDomainById(id string, attrs []string) *ZDomain {
	by := NewByRequest(ID_STR, id)
	return s.GetDomain(by, attrs)
}

func (s *ZAdmin) GetAllServers(service string) []ZServer {
	req, soapAction := NewGetAllServersRequest(service)
	resp := GetAllServersResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	servers := make([]ZServer, len(resp.Content.Server))

	for index, server := range resp.Content.Server {
		servers[index] = *NewServer(server, s.Client)
	}

	return servers
}

func (s *ZAdmin) GetQuotaUsage(serverId string, domain string, isAllServers bool) []ZAccount {
	allServers := 0

	if len(serverId) > 0 {
		s.Client.SetTargetServer(serverId)
	} else if isAllServers {
		allServers = 1
	}

	req, soapAction := NewGetQuotaUsageRequest(domain, allServers, 0, 0, "", 0, 0)
	resp := GetQuotaUsageResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	s.Client.RemoveTargetServer()

	accounts := make([]ZAccount, len(resp.Content.Account))

	for index, account := range resp.Content.Account {
		accounts[index] = *NewAccountQuota(account, s.Client)
	}

	return accounts
}

func (s *ZAdmin) GetAllBackups() []ZBackup {
	req, soapAction := NewBackupQueryRequest()
	resp := BackupQueryResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	return resp.Content.Backups
}

func (s *ZAdmin) SearchDirectoryCount(query string, domain string, types string) int {
	params := SearchDirectoryParams{
		Urn:       urnAdmin,
		Query:     query,
		Domain:    domain,
		Types:     types,
		CountOnly: 1,
	}

	req, soapAction := NewSearchDirectoryRequest(&params)
	resp := SearchDirectoryResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	return resp.Content.Count
}

func (s *ZAdmin) SearchDirectory(query string, maxResults int, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, types string, sortAscending int, attrs string) ([]ZAccount, []ZDomain, []ZCos) {
	params := SearchDirectoryParams{
		Urn:           urnAdmin,
		Query:         query,
		MaxResults:    maxResults,
		Limit:         limit,
		Offset:        offset,
		Domain:        domain,
		ApplyCos:      applyCos,
		ApplyConfig:   applyConfig,
		SortBy:        sortBy,
		Types:         types,
		SortAscending: sortAscending,
		CountOnly:     0,
		Attrs:         attrs,
	}

	req, soapAction := NewSearchDirectoryRequest(&params)
	resp := SearchDirectoryResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	accounts := make([]ZAccount, len(resp.Content.Accounts))

	for index, account := range resp.Content.Accounts {
		accounts[index] = *NewAccount(account, s.Client)
	}

	domains := make([]ZDomain, len(resp.Content.Domains))

	for index, domain := range resp.Content.Domains {
		domains[index] = *NewDomain(domain, s.Client)
	}

	coses := make([]ZCos, len(resp.Content.Coses))

	for index, cos := range resp.Content.Coses {
		coses[index] = *NewCos(cos, s.Client)
	}

	return accounts, domains, coses
}

func (s *ZAdmin) SearchDirectoryAll(query string, domain string, applyCos int, applyConfig int, sortBy string, types string, sortAscending int, attrs string) ([]ZAccount, []ZDomain, []ZCos) {
	accounts := make([]ZAccount, 0)
	domains := make([]ZDomain, 0)
	coses := make([]ZCos, 0)

	total := s.SearchDirectoryCount(query, domain, types)

	if total == 0 {
		return accounts, domains, coses
	}

	const maxResults = 1_000_000
	const limit = 500
	offset := 0

	for offset < total {
		_accounts, _domains, _coses := s.SearchDirectory(query, maxResults, limit, offset, domain, applyCos, applyConfig, sortBy, types, sortAscending, attrs)
		accounts = append(accounts, _accounts...)
		domains = append(domains, _domains...)
		coses = append(coses, _coses...)
		offset += limit
	}

	return accounts, domains, coses
}

func (s *ZAdmin) GetLicense() (*ZLicense, error) {
	req, soapAction := NewLicenseRequest()
	resp := GetLicenseResponse{}

	err := s.Client.Call(soapAction, req, &resp)

	if err == nil {
		return NewLicense(resp.Content), err
	} else {
		return nil, err
	}
}
