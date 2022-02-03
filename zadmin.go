package zsoap

import (
	"log"
)

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

	s.Client.TOKEN = resp.Content.TOKEN[0].Content
	s.Client.SetHeader()
}

func (s *ZAdmin) GetAccount(byAccount ByRequest, attrs []string) *ZAccount {

	req, soapAction := NewGetAccountRequest(byAccount, attrs)
	resp := GetAccountResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	account := NewAccount(resp.Content.Account[0])
	account.Client = s.Client

	return account
}

func (s *ZAdmin) GetAllAccounts(server *ByRequest, domain *ByRequest) []ZAccount {

	req, soapAction := NewGetAllAccountsRequest(server, domain)
	resp := GetAllAccountsResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	accounts := make([]ZAccount, 0)

	for _, account := range resp.Content.Account {
		accounts = append(accounts, *NewAccount(account))
	}

	return accounts
}

func (s *ZAdmin) GetAllCos() []ZCos {
	req, soapAction := NewGetAllCosRequest()
	resp := GetAllCosResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	coses := make([]ZCos, 0)

	for _, cos := range resp.Content.Cos {
		coses = append(coses, *NewCos(cos))
	}

	return coses
}

func (s *ZAdmin) GetAllDomains() []ZDomain {
	req, soapAction := NewGetAllDomainsRequest()
	resp := GetAllDomainsResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	domains := make([]ZDomain, 0)

	for _, domain := range resp.Content.Domain {
		domains = append(domains, *NewDomain(domain))
	}

	return domains
}

func (s *ZAdmin) GetAllServers(service string) []ZServer {
	req, soapAction := NewGetAllServersRequest(service)
	resp := GetAllServersResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	servers := make([]ZServer, 0)

	for _, server := range resp.Content.Server {
		servers = append(servers, *NewServer(server))
	}

	return servers
}

func (s *ZAdmin) GetQuotaUsage(domain string, isAllServers bool) []ZAccount {
	allServers := 0

	if isAllServers {
		allServers = 1
	}

	req, soapAction := NewGetQuotaUsageRequest(domain, allServers)
	resp := GetQuotaUsageResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	accounts := make([]ZAccount, 0)

	for _, account := range resp.Content.Account {
		accounts = append(accounts, *NewAccountQuota(account))
	}

	return accounts
}

func (s *ZAdmin) GetAllBackups() []ZBackup {
	req, soapAction := NewBackupQueryRequest()
	resp := BackupQueryResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}

	// backups := make([]ZBackup, 0)

	// for _, backup := range resp.Content.Backups {
	// 	backups = append(backups, *NewBackup(backup))
	// }

	// return backups
	return resp.Content.Backups
}

func (s *ZAdmin) Search(query string, maxResults int, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, types string, sortAscending int, countOnly int, attrs string) []ZAccount {
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
		CountOnly:     countOnly,
		Attrs:         attrs,
	}
	// fmt.Println(time.Now())
	req, soapAction := NewSearchDirectoryRequest(&params)
	resp := SearchDirectoryResponse{}
	// fmt.Println(time.Now())
	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(time.Now())
	accounts := make([]ZAccount, 0)

	for _, account := range resp.Content.Account {
		accounts = append(accounts, *NewAccount(account))
	}
	// fmt.Println(time.Now())
	return accounts
}

func (s *ZAdmin) SearchAccounts(query string, maxResults int, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, types string, sortAscending int, countOnly int, attrs string) []ZAccount {
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
		CountOnly:     countOnly,
		Attrs:         attrs,
	}
	// fmt.Println(time.Now())
	req, soapAction := NewSearchDirectoryRequest(&params)
	resp := SearchDirectoryResponse{}
	// fmt.Println(time.Now())
	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(time.Now())
	accounts := make([]ZAccount, 0)

	for _, account := range resp.Content.Account {
		accounts = append(accounts, *NewAccount(account))
	}
	// fmt.Println(time.Now())
	return accounts
}
