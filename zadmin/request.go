package zadmin

import (
	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
	"github.com/RaoH37/gosoap/zimbraConnector"
)

func (s *ZAdmin) invokeWithoutResponse(req interface{}) error {

	if err := s.Connector.Invoke(req, nil); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) AddAccountAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewAddAccountAliasRequest(id, alias)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) AddDistributionListAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewAddDistributionListAliasRequest(id, alias)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) AddDistributionListMemberRequest(id string, members []string) error {
	req := zimbraAdmin.NewAddDistributionListMemberRequest(id, members)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) AuthRequest() (*zimbraAdmin.AuthResponse, error) {
	req, resp := zimbraAdmin.NewAuthRequest(s.login, s.password)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) BackupQueryRequest() (*zimbraAdmin.BackupQueryResponse, error) {
	req, resp := zimbraAdmin.NewBackupQueryRequest()

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CopyCosRequest(id string, name string, newName string) (*zimbraAdmin.CopyCosResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	req, resp := zimbraAdmin.NewCopyCosRequest(by, newName)

	if err := s.Connector.Invoke(req, nil); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateAccountRequest(name string, password string, attrs zimbraCommon.AttrsNode) (*zimbraAdmin.CreateAccountResponse, error) {
	req, resp := zimbraAdmin.NewCreateAccountRequest(name, password, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateCalendarResourceRequest(name string, password string, attrs zimbraCommon.AttrsNode) (*zimbraAdmin.CreateCalendarResourceResponse, error) {
	req, resp := zimbraAdmin.NewCreateCalendarResourceRequest(name, password, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateCosRequest(name string, attrs zimbraCommon.AttrsNode) (*zimbraAdmin.CreateCosResponse, error) {
	req, resp := zimbraAdmin.NewCreateCosRequest(name, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateDistributionListRequest(name string, isDynamic bool, attrs zimbraCommon.AttrsNode) (*zimbraAdmin.CreateDistributionListResponse, error) {
	req, resp := zimbraAdmin.NewCreateDistributionListRequest(name, isDynamic, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) CreateDomainRequest(name string, attrs zimbraCommon.AttrsNode) (*zimbraAdmin.CreateDomainResponse, error) {
	req, resp := zimbraAdmin.NewCreateDomainRequest(name, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) DelegateAuthRequest(id string, name string) (*zimbraAdmin.DelegateAuthResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	req, resp := zimbraAdmin.NewDelegateAuthRequest(by)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) DeleteAccountRequest(id string) error {
	req := zimbraAdmin.NewDeleteAccountRequest(id)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) DeleteCalendarResourceRequest(id string) error {
	req := zimbraAdmin.NewDeleteCalendarResourceRequest(id)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) DeleteCosRequest(id string) error {
	req := zimbraAdmin.NewDeleteCosRequest(id)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) DeleteDistributionListRequest(id string, cascadeDelete bool) error {
	req := zimbraAdmin.NewDeleteDistributionListRequest(id, cascadeDelete)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) DeleteDomainRequest(id string) error {
	req := zimbraAdmin.NewDeleteDomainRequest(id)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) GetAccountRequest(id string, name string, attrs zimbraCommon.StringList, applyCos bool) (*zimbraAdmin.GetAccountResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	req, resp := zimbraAdmin.NewGetAccountRequest(by, attrs, applyCos)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetAllServersRequest(service string) (*zimbraAdmin.GetAllServersResponse, error) {
	req, resp := zimbraAdmin.NewGetAllServersRequest(service)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetCalendarResourceRequest(id string, name string, attrs zimbraCommon.StringList, applyCos bool) (*zimbraAdmin.GetCalendarResourceResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	req, resp := zimbraAdmin.NewGetCalendarResourceRequest(by, attrs, applyCos)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetDistributionListRequest(id string, name string, attrs zimbraCommon.StringList) (*zimbraAdmin.GetDistributionListResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	req, resp := zimbraAdmin.NewGetDistributionListRequest(by, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetCosRequest(id string, name string, attrs zimbraCommon.StringList) (*zimbraAdmin.GetCosResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	req, resp := zimbraAdmin.NewGetCosRequest(by, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetDomainRequest(id string, name string, attrs zimbraCommon.StringList, applyConfig bool) (*zimbraAdmin.GetDomainResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)
	req, resp := zimbraAdmin.NewGetDomainRequest(by, attrs, applyConfig)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetLicenseRequest() (*zimbraAdmin.GetLicenseResponse, error) {
	req, resp := zimbraAdmin.NewLicenseRequest()

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetQuotaUsageRequest(serverId string, domain string, allServers bool) (*zimbraAdmin.GetQuotaUsageResponse, error) {
	var connector *zimbraConnector.Connector

	if serverId != s.ServerId {
		// Use another connector for this request
		connector = s.NewConnector()
		connector.SetHeaderContext(s.GetToken(), serverId, nil)
	} else {
		connector = s.Connector
	}

	req, resp := zimbraAdmin.NewGetQuotaUsageRequest(domain, allServers, 0, 0, "", true, false)

	if err := connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetAllMailboxesRequest(serverId string) (*zimbraAdmin.GetAllMailboxesResponse, error) {
	var connector *zimbraConnector.Connector

	if serverId != s.ServerId {
		// Use another connector for this request
		connector = s.NewConnector()
		connector.SetHeaderContext(s.GetToken(), serverId, nil)
	} else {
		connector = s.Connector
	}

	req, resp := zimbraAdmin.NewGetAllMailboxesRequest()

	if err := connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) GetServerRequest(id string, name string, applyConfig bool, attrs zimbraCommon.StringList) (*zimbraAdmin.GetServerResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(id, name)

	req, resp := zimbraAdmin.NewGetServerRequest(by, applyConfig, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) ModifyAccountRequest(id string, attrs zimbraCommon.AttrsNode) error {
	req, resp := zimbraAdmin.NewModifyAccountRequest(id, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyCalendarResourceRequest(id string, attrs zimbraCommon.AttrsNode) error {
	req, resp := zimbraAdmin.NewModifyCalendarResourceRequest(id, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyCosRequest(id string, attrs zimbraCommon.AttrsNode) error {
	req, resp := zimbraAdmin.NewModifyCosRequest(id, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyDistributionListRequest(id string, attrs zimbraCommon.AttrsNode) error {
	req, resp := zimbraAdmin.NewModifyDistributionListRequest(id, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyDomainRequest(id string, attrs zimbraCommon.AttrsNode) error {
	req, resp := zimbraAdmin.NewModifyDomainRequest(id, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) ModifyServerRequest(id string, attrs zimbraCommon.AttrsNode) error {
	req, resp := zimbraAdmin.NewModifyServerRequest(id, attrs)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) NoOpRequest() error {
	req := zimbraAdmin.NewNoOpRequest()

	if err := s.Connector.Invoke(req, nil); err != nil {
		return err
	}

	return nil
}

func (s *ZAdmin) RemoveAccountAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewRemoveAccountAliasRequest(id, alias)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) RemoveDistributionListAliasRequest(id string, alias string) error {
	req := zimbraAdmin.NewRemoveDistributionListAliasRequest(id, alias)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) RemoveDistributionListMemberRequest(id string, members []string) error {
	req := zimbraAdmin.NewRemoveDistributionListMemberRequest(id, members)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) RenameAccountRequest(id string, newName string) error {
	req := zimbraAdmin.NewRenameAccountRequest(id, newName)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) RenameCalendarResourceRequest(id string, newName string) error {
	req := zimbraAdmin.NewRenameCalendarResourceRequest(id, newName)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) RenameDistributionListRequest(id string, newName string) error {
	req := zimbraAdmin.NewRenameDistributionListRequest(id, newName)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) RenameCosRequest(id string, newName string) error {
	req := zimbraAdmin.NewRenameCosRequest(id, newName)

	return s.invokeWithoutResponse(req)
}

func (s *ZAdmin) SearchDirectoryRequest(
	query string,
	maxResults int,
	limit int,
	offset int,
	domain string,
	applyCos bool,
	applyConfig bool,
	sortBy string,
	types zimbraAdmin.SearchTypeList,
	sortAscending bool,
	attrs zimbraCommon.StringList,
	isCountOnly bool) (*zimbraAdmin.SearchDirectoryResponse, error) {
	params := zimbraAdmin.SearchDirectoryParams{
		Urn:        "urn:zimbraAdmin",
		Query:      query,
		MaxResults: maxResults,
		Domain:     domain,
		Types:      types,
		CountOnly:  zimbraCommon.ZBool(isCountOnly),
	}

	if !isCountOnly {
		params.Limit = limit
		params.Offset = offset
		params.ApplyCos = zimbraCommon.ZBool(applyCos)
		params.ApplyConfig = zimbraCommon.ZBool(applyConfig)
		params.SortBy = sortBy
		params.SortAscending = zimbraCommon.ZBool(sortAscending)
		params.Attrs = attrs
	}

	req, resp := zimbraAdmin.NewSearchDirectoryRequest(&params)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZAdmin) SetPassword(id string, newPassword string) error {
	req := zimbraAdmin.NewSetPasswordRequest(id, newPassword)
	return s.invokeWithoutResponse(req)
}
