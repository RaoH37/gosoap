package zsoap

import "log"

type ZAdmin struct {
	AuthToken string
	Client *Client
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

	s.Client.SetHeader(resp.Content.TOKEN[0].Content)
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

func (s *ZAdmin) GetQuotaUsage(domain string) []ZAccount {
	req, soapAction := NewGetQuotaUsageRequest(domain)
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