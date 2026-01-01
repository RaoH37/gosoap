package zclient

import (
	"log"
	"time"

	"github.com/RaoH37/gosoap/zadmin"
	"github.com/RaoH37/gosoap/zimbraAdmin"
)

type ZClient struct {
	ZAdminClient     zadmin.ZAdmin
	SearchMaxResults int
	SearchLimit      int
}

func NewZClient(
	urlAdmin string,
	tls bool,
	login string,
	password string,
	debug bool,
	retryWaitingDuration time.Duration,
	userAgent string,
	timeout time.Duration) ZClient {
	return ZClient{
		ZAdminClient:     zadmin.NewZAdmin(urlAdmin, tls, login, password, debug, retryWaitingDuration, userAgent, timeout),
		SearchMaxResults: 1_000_000,
		SearchLimit:      2_000,
	}
}

func (s *ZClient) setToken(token string) {
	s.ZAdminClient.AuthToken = token
}

func (s *ZClient) getToken() string {
	return s.ZAdminClient.AuthToken
}

func (s *ZClient) Login() error {
	resp, err := s.ZAdminClient.AuthRequest()

	if err != nil {
		log.Println(err)
		return err
	}

	s.setToken(resp.Content.TOKEN[0].Content)
	return nil
}

func (s *ZClient) GetAccount(id string, name string, attrs []string, isApplyCos bool) (*ZAccount, error) {
	resp, err := s.ZAdminClient.GetAccountRequest(id, name, attrs, isApplyCos)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildAccount(resp.Content.Account[0]), nil
}

func (s *ZClient) GetAccounts(query string, limit int, offset int, domain string, isApplyCos bool, isApplyConfig bool, sortBy string, isSortAscending bool, attrs string) ([]ZAccount, error) {
	accounts, _, _, _, _, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, isApplyCos, isApplyConfig, sortBy, "accounts", isSortAscending, attrs)
	return accounts, err
}

func (s *ZClient) GetAllAccounts(query string, domain string, isApplyCos bool, isApplyConfig bool, sortBy string, isSortAscending bool, attrs string) ([]ZAccount, error) {
	accounts, _, _, _, _, err := s.SearchDirectoryAll(query, domain, isApplyCos, isApplyConfig, sortBy, "accounts", isSortAscending, attrs)
	return accounts, err
}

func (s *ZClient) GetAllBackups() ([]zimbraAdmin.ZBackup, error) {
	resp, err := s.ZAdminClient.BackupQueryRequest()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp.Content.Backups, nil
}

func (s *ZClient) GetAllCalendarResources(query string, domain string, isApplyCos bool, isApplyConfig bool, sortBy string, isSortAscending bool, attrs string) ([]ZCalendarResource, error) {
	_, _, _, _, calresources, err := s.SearchDirectoryAll(query, domain, isApplyCos, isApplyConfig, sortBy, "resources", isSortAscending, attrs)
	return calresources, err
}

func (s *ZClient) GetAllCoses(query string, isApplyConfig bool, sortBy string, isSortAscending bool, attrs string) ([]ZCos, error) {
	_, _, _, coses, _, err := s.SearchDirectoryAll(query, "", true, isApplyConfig, sortBy, "coses", isSortAscending, attrs)
	return coses, err
}

func (s *ZClient) GetCalendarResource(id string, name string, attrs []string, isApplyCos bool) (*ZCalendarResource, error) {
	resp, err := s.ZAdminClient.GetCalendarResourceRequest(id, name, attrs, isApplyCos)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildResource(resp.Content.CalResource[0]), nil
}

func (s *ZClient) GetCalendarResources(query string, limit int, offset int, domain string, isApplyCos bool, isApplyConfig bool, sortBy string, isSortAscending bool, attrs string) ([]ZCalendarResource, error) {
	_, _, _, _, calresources, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, isApplyCos, isApplyConfig, sortBy, "resources", isSortAscending, attrs)
	return calresources, err
}

func (s *ZClient) GetQuotaUsage(serverId string, domain string, isAllServers bool) ([]ZAccount, error) {
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

func (s *ZClient) GetAllDistributionLists(query string, domain string, isApplyCos bool, isApplyConfig bool, sortBy string, isSortAscending bool, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, _, err := s.SearchDirectoryAll(query, domain, isApplyCos, isApplyConfig, sortBy, "distributionlists", isSortAscending, attrs)
	return dls, err
}

func (s *ZClient) GetDistributionLists(query string, limit int, offset int, domain string, isApplyCos bool, isApplyConfig bool, sortBy string, isSortAscending bool, attrs string) ([]ZDistributionList, error) {
	_, dls, _, _, _, err := s.SearchDirectory(query, 1_000_000, limit, offset, domain, isApplyCos, isApplyConfig, sortBy, "distributionlists", isSortAscending, attrs)
	return dls, err
}

func (s *ZClient) GetDistributionList(id string, name string, attrs []string) (*ZDistributionList, error) {
	resp, err := s.ZAdminClient.GetDistributionListRequest(id, name, attrs)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildDistributionList(resp.Content.Dl[0]), nil
}

func (s *ZClient) GetAllDomains(
	query string,
	isApplyCos bool,
	isApplyConfig bool,
	sortBy string,
	isSortAscending bool,
	attrs string) ([]ZDomain, error) {
	_, _, domains, _, _, err := s.SearchDirectoryAll(query, "", isApplyCos, isApplyConfig, sortBy, "domains", isSortAscending, attrs)
	return domains, err
}

func (s *ZClient) GetDomain(id string, name string, attrs []string, isApplyConfig bool) (*ZDomain, error) {
	resp, err := s.ZAdminClient.GetDomainRequest(id, name, attrs, isApplyConfig)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildDomain(resp.Content.Domains[0]), nil
}

func (s *ZClient) GetLicense() (*ZLicense, error) {
	resp, err := s.ZAdminClient.GetLicenseRequest()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildLicense(resp.Content), nil
}

func (s *ZClient) GetAllServers(service string) ([]ZServer, error) {
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

func (s *ZClient) GetServer(id string, name string, isApplyConfig bool, attrs []string) (*ZServer, error) {
	resp, err := s.ZAdminClient.GetServerRequest(id, name, isApplyConfig, attrs)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildServer(resp.Content.Server[0]), nil
}

func (s *ZClient) SearchDirectory(
	query string,
	maxResults int,
	limit int,
	offset int,
	domain string,
	isApplyCos bool,
	isApplyConfig bool,
	sortBy string,
	types string,
	isSortAscending bool,
	attrs string) ([]ZAccount, []ZDistributionList, []ZDomain, []ZCos, []ZCalendarResource, error) {

	resp, err := s.ZAdminClient.SearchDirectoryRequest(query, maxResults, limit, offset, domain, isApplyCos, isApplyConfig, sortBy, types, isSortAscending, attrs, false)
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

func (s *ZClient) SearchDirectoryAll(
	query string,
	domain string,
	isApplyCos bool,
	isApplyConfig bool,
	sortBy string,
	types string,
	isSortAscending bool,
	attrs string) ([]ZAccount, []ZDistributionList, []ZDomain, []ZCos, []ZCalendarResource, error) {
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
		_accounts, _dls, _domains, _coses, _calresources, err := s.SearchDirectory(query, s.SearchMaxResults, s.SearchLimit, offset, domain, isApplyCos, isApplyConfig, sortBy, types, isSortAscending, attrs)
		if err != nil {
			log.Println(err)
			retries -= 1

			if retries <= 0 {
				return nil, nil, nil, nil, nil, err
			}

			time.Sleep(s.ZAdminClient.RetryWaitingDuration)
			continue
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

func (s *ZClient) SearchDirectoryCount(query string, domain string, types string) (int, error) {
	resp, err := s.ZAdminClient.SearchDirectoryRequest(query, 1_000_000, 1, 0, domain, false, false, "", types, false, "", true)
	if err != nil {
		return 0, err
	}

	return resp.Content.Count, nil
}
