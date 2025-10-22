package zsoap

import (
	"log"
	"time"

	"github.com/RaoH37/gosoap/zadmin"
)

type ZcsClient struct {
	ZAdminClient     zadmin.ZAdmin
	SearchMaxResults int
	SearchLimit      int
}

func NewZcsClient(urlAdmin string, tls bool, login string, password string, debug bool, retryWaitingDuration time.Duration, userAgent string, timeout time.Duration) ZcsClient {
	return ZcsClient{
		ZAdminClient:     zadmin.NewZAdmin(urlAdmin, tls, login, password, debug, retryWaitingDuration, userAgent, timeout),
		SearchMaxResults: 1_000_000,
		SearchLimit:      2_000,
	}
}

func (s *ZcsClient) setToken(token string) {
	s.ZAdminClient.AuthToken = token
}

func (s *ZcsClient) getToken() string {
	return s.ZAdminClient.AuthToken
}

func (s *ZcsClient) Login() error {
	resp, err := s.ZAdminClient.AuthRequest()

	if err != nil {
		log.Println(err)
		return err
	} else {
		s.setToken(resp.Content.TOKEN[0].Content)
		return nil
	}
}

func (s *ZcsClient) GetAccount(id string, name string, attrs []string) (*ZAccount, error) {
	resp, err := s.ZAdminClient.GetAccountRequest(id, name, attrs)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildAccount(resp.Content.Account[0]), nil
}

func (s *ZcsClient) GetAccounts(query string, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZAccount, error) {
	accounts, _, _, _, _, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, applyCos, applyConfig, sortBy, "accounts", sortAscending, attrs)
	return accounts, err
}

func (s *ZcsClient) GetAllAccounts(query string, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZAccount, error) {
	accounts, _, _, _, _, err := s.SearchDirectoryAll(query, domain, applyCos, applyConfig, sortBy, "accounts", sortAscending, attrs)
	return accounts, err
}

func (s *ZcsClient) GetAllCalendarResources(query string, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZCalendarResource, error) {
	_, _, _, _, calresources, err := s.SearchDirectoryAll(query, domain, applyCos, applyConfig, sortBy, "resources", sortAscending, attrs)
	return calresources, err
}

func (s *ZcsClient) GetAllCoses(query string, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZCos, error) {
	_, _, _, coses, _, err := s.SearchDirectoryAll(query, "", 0, applyConfig, sortBy, "coses", sortAscending, attrs)
	return coses, err
}

func (s *ZcsClient) GetCalendarResource(id string, name string, attrs []string) (*ZCalendarResource, error) {
	resp, err := s.ZAdminClient.GetCalendarResourceRequest(id, name, attrs)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildResource(resp.Content.CalResource[0]), nil
}

func (s *ZcsClient) GetCalendarResources(query string, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZCalendarResource, error) {
	_, _, _, _, calresources, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, applyCos, applyConfig, sortBy, "resources", sortAscending, attrs)
	return calresources, err
}

func (s *ZcsClient) GetQuotaUsage(serverId string, domain string, isAllServers bool) ([]ZAccount, error) {
	resp, err := s.ZAdminClient.GetQuotaUsageRequest(serverId, domain, isAllServers)

	if err != nil {
		return nil, err
	}

	accounts := make([]ZAccount, len(resp.Content.Account))

	for index, account := range resp.Content.Account {
		accounts[index] = *buildAccountQuota(account)
	}

	return accounts, nil
}

func (s *ZcsClient) GetAllDistributionLists(query string, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, _, err := s.SearchDirectoryAll(query, domain, applyCos, applyConfig, sortBy, "distributionlists", sortAscending, attrs)
	return dls, err
}

func (s *ZcsClient) GetDistributionLists(query string, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, _, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, applyCos, applyConfig, sortBy, "distributionlists", sortAscending, attrs)
	return dls, err
}

func (s *ZcsClient) GetDistributionList(id string, name string, attrs []string) (*ZDistributionList, error) {
	resp, err := s.ZAdminClient.GetDistributionListRequest(id, name, attrs)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildDistributionList(resp.Content.Dl[0]), nil
}

func (s *ZcsClient) GetAllDomains(query string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDomain, error) {
	_, _, domains, _, _, err := s.SearchDirectoryAll(query, "", applyCos, applyConfig, sortBy, "domains", sortAscending, attrs)
	return domains, err
}

func (s *ZcsClient) GetDomain(id string, name string, attrs []string) (*ZDomain, error) {
	resp, err := s.ZAdminClient.GetDomainRequest(id, name, attrs)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	domain := buildDomain(resp.Content.Domains[0])

	return domain, nil
}

func (s *ZcsClient) GetLicense() (*ZLicense, error) {
	resp, err := s.ZAdminClient.GetLicenseRequest()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildLicense(resp.Content), nil
}

func (s *ZcsClient) GetAllServers(service string) ([]ZServer, error) {
	resp, err := s.ZAdminClient.GetAllServersRequest(service)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	servers := make([]ZServer, len(resp.Content.Servers))

	for index, server := range resp.Content.Servers {
		servers[index] = *buildServer(server)
	}

	return servers, nil
}

func (s *ZcsClient) GetServer(id string, name string, applyConfig int, attrs []string) (*ZServer, error) {
	resp, err := s.ZAdminClient.GetServerRequest(id, name, applyConfig, attrs)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	server := buildServer(resp.Content.Server[0])

	return server, nil
}

func (s *ZcsClient) SearchDirectory(
	query string,
	maxResults int,
	limit int,
	offset int,
	domain string,
	applyCos int,
	applyConfig int,
	sortBy string,
	types string,
	sortAscending int,
	attrs string) ([]ZAccount, []ZDistributionList, []ZDomain, []ZCos, []ZCalendarResource, error) {

	resp, err := s.ZAdminClient.SearchDirectoryRequest(query, maxResults, limit, offset, domain, applyCos, applyConfig, sortBy, types, sortAscending, attrs, 0)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	accounts := make([]ZAccount, len(resp.Content.Accounts))

	for index, account := range resp.Content.Accounts {
		accounts[index] = *buildAccount(account)
	}

	dls := make([]ZDistributionList, len(resp.Content.Dls))

	for index, dl := range resp.Content.Dls {
		dls[index] = *buildDistributionList(dl)
	}

	domains := make([]ZDomain, len(resp.Content.Domains))

	for index, domain := range resp.Content.Domains {
		domains[index] = *buildDomain(domain)
	}

	coses := make([]ZCos, len(resp.Content.Coses))

	for index, cos := range resp.Content.Coses {
		coses[index] = *buildCos(cos)
	}

	calresources := make([]ZCalendarResource, len(resp.Content.CalResources))

	for index, calresource := range resp.Content.CalResources {
		calresources[index] = *buildResource(calresource)
	}

	return accounts, dls, domains, coses, calresources, nil
}

func (s *ZcsClient) SearchDirectoryAll(query string, domain string, applyCos int, applyConfig int, sortBy string, types string, sortAscending int, attrs string) ([]ZAccount, []ZDistributionList, []ZDomain, []ZCos, []ZCalendarResource, error) {
	accounts := make([]ZAccount, 0)
	dls := make([]ZDistributionList, 0)
	domains := make([]ZDomain, 0)
	coses := make([]ZCos, 0)
	calresources := make([]ZCalendarResource, 0)

	total, err := s.SearchDirectoryCount(query, domain, types)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, nil, nil, err
	}

	if total == 0 {
		return accounts, dls, domains, coses, calresources, nil
	}

	offset := 0
	retries := 3

	for offset < total {
		_accounts, _dls, _domains, _coses, _calresources, err := s.SearchDirectory(query, s.SearchMaxResults, s.SearchLimit, offset, domain, applyCos, applyConfig, sortBy, types, sortAscending, attrs)
		if err != nil {
			log.Println(err)
			retries -= 1

			if retries <= 0 {
				return nil, nil, nil, nil, nil, err
			} else {
				time.Sleep(s.ZAdminClient.RetryWaitingDuration)
				continue
			}
		}

		accounts = append(accounts, _accounts...)
		dls = append(dls, _dls...)
		domains = append(domains, _domains...)
		coses = append(coses, _coses...)
		calresources = append(calresources, _calresources...)
		offset += s.SearchLimit
	}

	return accounts, dls, domains, coses, calresources, nil
}

func (s *ZcsClient) SearchDirectoryCount(query string, domain string, types string) (int, error) {
	resp, err := s.ZAdminClient.SearchDirectoryRequest(query, 1_000_000, 1, 0, domain, 0, 0, "", types, 0, "", 1)
	if err != nil {
		return 0, err
	}

	return resp.Content.Count, nil
}
