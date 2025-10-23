package zadmin

import (
	"log"

	"github.com/RaoH37/gosoap/zimbraAdmin"
)

func (s *ZAdmin) invokeWithoutResponse(req interface{}, serverId string, accountName string, userAgent string) error {
	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, serverId, accountName, userAgent)

	if err := connector.Invoke(req, nil); err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}

func (s *ZAdmin) AddAccountAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewAddAccountAliasRequest(id, alias)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) AddDistributionListAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewAddDistributionListAliasRequest(id, alias)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) AddDistributionListMemberRequest(id string, members []string) error {
	req := zimbraAdmin.NewAddDistributionListMemberRequest(id, members)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) AuthRequest() (*zimbraAdmin.AuthResponse, error) {
	req, resp := zimbraAdmin.NewAuthRequest(s.login, s.password)

	connector := s.buildZimbraConnector()
	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) BackupQueryRequest() (*zimbraAdmin.BackupQueryResponse, error) {
	req, resp := zimbraAdmin.NewBackupQueryRequest()

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CopyCosRequest(id string, name string, newName string) (*zimbraAdmin.CopyCosResponse, error) {
	by := s.byNode(id, name)
	req, resp := zimbraAdmin.NewCopyCosRequest(by, newName)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, nil); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return &resp, nil
	}
}

func (s *ZAdmin) CreateAccountRequest(name string, password string, attrs map[string]string) (*zimbraAdmin.CreateAccountResponse, error) {
	req, resp := zimbraAdmin.NewCreateAccountRequest(name, password, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateCalendarResourceRequest(name string, password string, attrs map[string]string) (*zimbraAdmin.CreateCalendarResourceResponse, error) {
	req, resp := zimbraAdmin.NewCreateCalendarResourceRequest(name, password, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateCosRequest(name string, attrs map[string]string) (*zimbraAdmin.CreateCosResponse, error) {
	req, resp := zimbraAdmin.NewCreateCosRequest(name, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateDistributionListRequest(name string, dynamic int, attrs map[string]string) (*zimbraAdmin.CreateDistributionListResponse, error) {
	req, resp := zimbraAdmin.NewCreateDistributionListRequest(name, dynamic, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateDomainRequest(name string, attrs map[string]string) (*zimbraAdmin.CreateDomainResponse, error) {
	req, resp := zimbraAdmin.NewCreateDomainRequest(name, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) DelegateAuthRequest(id string, name string) (*zimbraAdmin.DelegateAuthResponse, error) {
	by := s.byNode(id, name)
	req, resp := zimbraAdmin.NewDelegateAuthRequest(by)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, nil); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return &resp, nil
	}
}

func (s *ZAdmin) DeleteAccountRequest(id string) error {
	req := zimbraAdmin.NewDeleteAccountRequest(id)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) DeleteCalendarResourceRequest(id string) error {
	req := zimbraAdmin.NewDeleteCalendarResourceRequest(id)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) DeleteCosRequest(id string) error {
	req := zimbraAdmin.NewDeleteCosRequest(id)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) DeleteDistributionListRequest(id string, cascadeDelete bool) error {
	req := zimbraAdmin.NewDeleteDistributionListRequest(id, cascadeDelete)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) DeleteDomainRequest(id string) error {
	req := zimbraAdmin.NewDeleteDomainRequest(id)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) GetAccountRequest(id string, name string, attrs []string) (*zimbraAdmin.GetAccountResponse, error) {
	by := s.byNode(id, name)
	req, resp := zimbraAdmin.NewGetAccountRequest(by, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetAllServersRequest(service string) (*zimbraAdmin.GetAllServersResponse, error) {
	req, resp := zimbraAdmin.NewGetAllServersRequest(service)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetCalendarResourceRequest(id string, name string, attrs []string) (*zimbraAdmin.GetCalendarResourceResponse, error) {
	by := s.byNode(id, name)
	req, resp := zimbraAdmin.NewGetCalendarResourceRequest(by, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetDistributionListRequest(id string, name string, attrs []string) (*zimbraAdmin.GetDistributionListResponse, error) {
	by := s.byNode(id, name)
	req, resp := zimbraAdmin.NewGetDistributionListRequest(by, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetDomainRequest(id string, name string, attrs []string) (*zimbraAdmin.GetDomainResponse, error) {
	by := s.byNode(id, name)
	req, resp := zimbraAdmin.NewGetDomainRequest(by, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetLicenseRequest() (*zimbraAdmin.GetLicenseResponse, error) {
	req, resp := zimbraAdmin.NewLicenseRequest()

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetQuotaUsageRequest(serverId string, domain string, isAllServers bool) (*zimbraAdmin.GetQuotaUsageResponse, error) {
	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, serverId, "", "")

	allServers := 0

	if isAllServers {
		allServers = 1
	}

	req, resp := zimbraAdmin.NewGetQuotaUsageRequest(domain, allServers, 0, 0, "", 0, 0)

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetServerRequest(id string, name string, applyConfig int, attrs []string) (*zimbraAdmin.GetServerResponse, error) {
	by := s.byNode(id, name)
	req, resp := zimbraAdmin.NewGetServerRequest(by, applyConfig, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) ModifyAccountRequest(id string, attrs map[string]string) error {
	req, resp := zimbraAdmin.NewModifyAccountRequest(id, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyCalendarResourceRequest(id string, attrs map[string]string) error {
	req, resp := zimbraAdmin.NewModifyCalendarResourceRequest(id, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyCosRequest(id string, attrs map[string]string) error {
	req, resp := zimbraAdmin.NewModifyCosRequest(id, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyDistributionListRequest(id string, attrs map[string]string) error {
	req, resp := zimbraAdmin.NewModifyDistributionListRequest(id, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyDomainRequest(id string, attrs map[string]string) error {
	req, resp := zimbraAdmin.NewModifyDomainRequest(id, attrs)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ZAdmin) RemoveAccountAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewRemoveAccountAliasRequest(id, alias)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) RemoveDistributionListAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewRemoveDistributionListAliasRequest(id, alias)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) RemoveDistributionListMemberRequest(id string, members []string) error {
	req := zimbraAdmin.NewRemoveDistributionListMemberRequest(id, members)

	return s.invokeWithoutResponse(req, "", "", "")
}

func (s *ZAdmin) SearchDirectoryRequest(query string, maxResults int, limit int, offset int, domain string, applyCos int, applyConfig int, sortBy string, types string, sortAscending int, attrs string, countOnly int) (*zimbraAdmin.SearchDirectoryResponse, error) {
	params := zimbraAdmin.SearchDirectoryParams{
		Urn:        "urn:zimbraAdmin",
		Query:      query,
		MaxResults: maxResults,
		Domain:     domain,
		Types:      types,
		CountOnly:  countOnly,
	}

	if countOnly == 0 {
		params.Limit = limit
		params.Offset = offset
		params.ApplyCos = applyCos
		params.ApplyConfig = applyConfig
		params.SortBy = sortBy
		params.SortAscending = sortAscending
		params.Attrs = attrs
	}

	req, resp := zimbraAdmin.NewSearchDirectoryRequest(&params)

	connector := s.buildZimbraConnector()
	connector.SetHeaderContext(s.AuthToken, "", "", "")

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}
