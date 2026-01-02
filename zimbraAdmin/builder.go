package zimbraAdmin

import (
	"strings"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

const Urn = "urn:zimbraAdmin"

func NewAuthRequest(name string, password string) (*AuthRequest, AuthResponse) {
	return &AuthRequest{
		Content: AuthRequestContent{
			Name:     name,
			Password: password,
			Urn:      Urn,
		},
	}, AuthResponse{}
}

func NewGetAccountRequest(by zimbraCommon.ByNode, attrs []string, applyCos bool) (*GetAccountRequest, GetAccountResponse) {
	r := &GetAccountRequest{
		Content: GetAccountRequestContent{
			Urn:      Urn,
			Account:  by,
			ApplyCos: zimbraCommon.ZBool(applyCos),
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetAccountResponse{}
}

func NewGetAllConfigRequest(by zimbraCommon.ByNode, attrs []string) (*GetAllConfigRequest, GetAllConfigResponse) {
	return &GetAllConfigRequest{
		Content: newUrnRequestContent(),
	}, GetAllConfigResponse{}
}

func NewGetDistributionListRequest(by zimbraCommon.ByNode, attrs []string) (*GetDistributionListRequest, GetDistributionListResponse) {
	r := &GetDistributionListRequest{
		Content: GetDistributionListRequestContent{
			Urn: Urn,
			Dl:  by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetDistributionListResponse{}
}

func NewGetCalendarResourceRequest(by zimbraCommon.ByNode, attrs []string, applyCos bool) (*GetCalendarResourceRequest, GetCalendarResourceResponse) {
	r := &GetCalendarResourceRequest{
		Content: GetCalendarResourceRequestContent{
			Urn:         Urn,
			CalResource: by,
			ApplyCos:    zimbraCommon.ZBool(applyCos),
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetCalendarResourceResponse{}
}

func NewGetCosRequest(by zimbraCommon.ByNode, attrs []string) (*GetCosRequest, GetCosResponse) {
	r := &GetCosRequest{
		Content: GetCosRequestContent{
			Urn: Urn,
			Cos: by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetCosResponse{}
}

func NewGetDomainRequest(by zimbraCommon.ByNode, attrs []string, applyConfig bool) (*GetDomainRequest, GetDomainResponse) {
	r := &GetDomainRequest{
		Content: GetDomainRequestContent{
			Urn:         Urn,
			Domain:      by,
			ApplyConfig: zimbraCommon.ZBool(applyConfig),
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetDomainResponse{}
}

func NewGetAllServersRequest(service string) (*GetAllServersRequest, GetAllServersResponse) {
	r := &GetAllServersRequest{
		Content: GetAllServersRequestContent{
			Urn: Urn,
		},
	}

	if service != "" {
		r.Content.Service = service
	}

	return r, GetAllServersResponse{}
}

func NewGetServerRequest(by zimbraCommon.ByNode, applyConfig bool, attrs []string) (*GetServerRequest, GetServerResponse) {
	r := &GetServerRequest{
		Content: GetServerRequestContent{
			Urn:         Urn,
			ApplyConfig: zimbraCommon.ZBool(applyConfig),
			Server:      by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetServerResponse{}
}

func NewGetQuotaUsageRequest(domain string, allServers bool, limit int, offset int, sortBy string, sortAscending bool, refresh bool) (*GetQuotaUsageRequest, GetQuotaUsageResponse) {
	return &GetQuotaUsageRequest{
		Content: GetQuotaUsageRequestContent{
			Urn:           Urn,
			Servers:       zimbraCommon.ZBool(allServers),
			Domain:        domain,
			Limit:         limit,
			Offset:        offset,
			SortBy:        sortBy,
			SortAscending: zimbraCommon.ZBool(sortAscending),
			Refresh:       zimbraCommon.ZBool(refresh),
		},
	}, GetQuotaUsageResponse{}
}

func NewBackupQueryRequest() (*BackupQueryRequest, BackupQueryResponse) {
	return &BackupQueryRequest{
		Content: BackupQueryRequestContent{
			Urn:   Urn,
			Query: make(map[string]string),
		},
	}, BackupQueryResponse{}
}

func NewCopyCosRequest(by zimbraCommon.ByNode, newName string) (*CopyCosRequest, CopyCosResponse) {
	return &CopyCosRequest{
		Content: CopyCosRequestContent{
			Urn: Urn,
			By:  by,
			Name: zimbraCommon.ContentString{
				Content: newName,
			},
		},
	}, CopyCosResponse{}
}

func NewCreateAccountRequest(name string, password string, attrs map[string]string) (*CreateAccountRequest, CreateAccountResponse) {
	return &CreateAccountRequest{
		Content: CreateAccountRequestContent{
			Urn:      Urn,
			Name:     name,
			Password: password,
			Attrs:    buildAttrResponses(attrs),
		},
	}, CreateAccountResponse{}
}

func NewCreateCalendarResourceRequest(name string, password string, attrs map[string]string) (*CreateCalendarResourceRequest, CreateCalendarResourceResponse) {
	return &CreateCalendarResourceRequest{
		Content: CreateCalendarResourceRequestContent{
			Urn:      Urn,
			Name:     name,
			Password: password,
			Attrs:    buildAttrResponses(attrs),
		},
	}, CreateCalendarResourceResponse{}
}

func NewCreateCosRequest(name string, attrs map[string]string) (*CreateCosRequest, CreateCosResponse) {
	return &CreateCosRequest{
		Content: CreateCosRequestContent{
			Urn: Urn,
			Name: zimbraCommon.ContentString{
				Content: name,
			},
			Attrs: buildAttrResponses(attrs),
		},
	}, CreateCosResponse{}
}

func NewCreateDistributionListRequest(name string, dynamic bool, attrs map[string]string) (*CreateDistributionListRequest, CreateDistributionListResponse) {
	return &CreateDistributionListRequest{
		Content: CreateDistributionListRequestContent{
			Urn:     Urn,
			Name:    name,
			Dynamic: zimbraCommon.ZBool(dynamic),
			Attrs:   buildAttrResponses(attrs),
		},
	}, CreateDistributionListResponse{}
}

func NewCreateDomainRequest(name string, attrs map[string]string) (*CreateDomainRequest, CreateDomainResponse) {
	return &CreateDomainRequest{
		Content: CreateDomainRequestContent{
			Urn:   Urn,
			Name:  name,
			Attrs: buildAttrResponses(attrs),
		},
	}, CreateDomainResponse{}
}

func NewDeleteCalendarResourceRequest(id string) *DeleteCalendarResourceRequest {
	return &DeleteCalendarResourceRequest{
		Content: newIdRequestContent(id),
	}
}

func NewDeleteCosRequest(id string) *DeleteCosRequest {
	return &DeleteCosRequest{
		Content: DeleteCosRequestContent{
			Urn: Urn,
			ID: zimbraCommon.ContentString{
				Content: id,
			},
		},
	}
}

func NewDeleteDistributionListRequest(id string, cascadeDelete bool) *DeleteDistributionListRequest {
	return &DeleteDistributionListRequest{
		Content: DeleteDistributionListRequestContent{
			Urn:           Urn,
			ID:            id,
			CascadeDelete: zimbraCommon.ZBool(cascadeDelete),
		},
	}
}

func NewDeleteDomainRequest(id string) *DeleteDomainRequest {
	return &DeleteDomainRequest{
		Content: newIdRequestContent(id),
	}
}

func NewDeleteAccountRequest(id string) *DeleteAccountRequest {
	return &DeleteAccountRequest{
		Content: newIdRequestContent(id),
	}
}

func newIdRequestContent(id string) IdRequestContent {
	return IdRequestContent{
		Urn: Urn,
		ID:  id,
	}
}

func NewDelegateAuthRequest(by zimbraCommon.ByNode) (*DelegateAuthRequest, DelegateAuthResponse) {
	return &DelegateAuthRequest{
		Content: DelegateAuthRequestContent{
			Urn:     Urn,
			Account: by,
		},
	}, DelegateAuthResponse{}
}

func NewDeleteGalSyncAccountRequest(by zimbraCommon.ByNode) *DeleteGalSyncAccountRequest {
	return &DeleteGalSyncAccountRequest{
		Content: DeleteGalSyncAccountRequestContent{
			Urn:     Urn,
			Account: by,
		},
	}
}

func NewSearchDirectoryRequest(params *SearchDirectoryParams) (*SearchDirectoryRequest, SearchDirectoryResponse) {
	return &SearchDirectoryRequest{
		Content: *params,
	}, SearchDirectoryResponse{}
}

func NewLicenseRequest() (*GetLicenseRequest, GetLicenseResponse) {
	return &GetLicenseRequest{
		Content: newUrnRequestContent(),
	}, GetLicenseResponse{}
}

func newUrnRequestContent() zimbraCommon.UrnRequestContent {
	return zimbraCommon.UrnRequestContent{
		Urn: Urn,
	}
}

func newModifyRequestContent(id string, a []zimbraCommon.AttrNode) ModifyRequestContent {
	return ModifyRequestContent{
		Urn:   Urn,
		ID:    id,
		Attrs: a,
	}
}

func NewModifyAccountRequest(id string, attrs map[string]string) (*ModifyAccountRequest, ModifyAccountResponse) {
	return &ModifyAccountRequest{
		Content: newModifyRequestContent(id, buildAttrResponses(attrs)),
	}, ModifyAccountResponse{}
}

func NewAddAccountAliasRequest(id string, alias string) *AddAccountAliasRequest {
	return &AddAccountAliasRequest{
		Content: newAccountAliasRequestContent(id, alias),
	}
}

func NewRemoveAccountAliasRequest(id string, alias string) *RemoveAccountAliasRequest {
	return &RemoveAccountAliasRequest{
		Content: newAccountAliasRequestContent(id, alias),
	}
}

func newAccountAliasRequestContent(id string, alias string) AliasRequestContent {
	return AliasRequestContent{
		Urn:   Urn,
		ID:    id,
		Alias: alias,
	}
}

func NewModifyCalendarResourceRequest(id string, attrs map[string]string) (*ModifyCalendarResourceRequest, ModifyCalendarResourceResponse) {
	return &ModifyCalendarResourceRequest{
		Content: newModifyRequestContent(id, buildAttrResponses(attrs)),
	}, ModifyCalendarResourceResponse{}
}

func NewAddDistributionListAliasRequest(id string, alias string) *AddDistributionListAliasRequest {
	return &AddDistributionListAliasRequest{
		Content: newDistributionListAliasRequestContent(id, alias),
	}
}

func NewRemoveDistributionListAliasRequest(id string, alias string) *RemoveDistributionListAliasRequest {
	return &RemoveDistributionListAliasRequest{
		Content: newDistributionListAliasRequestContent(id, alias),
	}
}

func newDistributionListAliasRequestContent(id string, alias string) AliasRequestContent {
	return AliasRequestContent{
		Urn:   Urn,
		ID:    id,
		Alias: alias,
	}
}

func NewAddDistributionListMemberRequest(id string, members []string) *AddDistributionListMemberRequest {
	return &AddDistributionListMemberRequest{
		Content: newDistributionListMemberRequestContent(id, members),
	}
}

func NewRemoveDistributionListMemberRequest(id string, members []string) *RemoveDistributionListMemberRequest {
	return &RemoveDistributionListMemberRequest{
		Content: newDistributionListMemberRequestContent(id, members),
	}
}

func newDistributionListMemberRequestContent(id string, members []string) DistributionListMemberRequestContent {
	return DistributionListMemberRequestContent{
		Urn:     Urn,
		ID:      id,
		Members: convertToContentStrings(members),
	}
}

func convertToContentStrings(arr []string) []zimbraCommon.ContentString {
	contentStrings := make([]zimbraCommon.ContentString, len(arr))

	for i, a := range arr {
		contentStrings[i] = zimbraCommon.ContentString{Content: a}
	}

	return contentStrings
}

func NewGetLicenseRequest() (*GetLicenseRequest, GetLicenseResponse) {
	return &GetLicenseRequest{
		Content: newUrnRequestContent(),
	}, GetLicenseResponse{}
}

func NewModifyCosRequest(id string, attrs map[string]string) (*ModifyCosRequest, ModifyCosResponse) {
	return &ModifyCosRequest{
		Content: ModifyCosRequestContent{
			Urn:   Urn,
			ID:    zimbraCommon.ContentString{Content: id},
			Attrs: buildAttrResponses(attrs),
		},
	}, ModifyCosResponse{}
}

func NewModifyDistributionListRequest(id string, attrs map[string]string) (*ModifyDistributionListRequest, ModifyDistributionListResponse) {
	return &ModifyDistributionListRequest{
		Content: newModifyRequestContent(id, buildAttrResponses(attrs)),
	}, ModifyDistributionListResponse{}
}

func NewModifyDomainRequest(id string, attrs map[string]string) (*ModifyDomainRequest, ModifyDomainResponse) {
	return &ModifyDomainRequest{
		Content: newModifyRequestContent(id, buildAttrResponses(attrs)),
	}, ModifyDomainResponse{}
}

func NewModifyServerRequest(id string, attrs map[string]string) (*ModifyServerRequest, ModifyServerResponse) {
	return &ModifyServerRequest{
		Content: newModifyRequestContent(id, buildAttrResponses(attrs)),
	}, ModifyServerResponse{}
}

func buildAttrResponses(attrs map[string]string) []zimbraCommon.AttrNode {
	a := make([]zimbraCommon.AttrNode, len(attrs))

	i := 0
	for name, value := range attrs {
		a[i] = zimbraCommon.AttrNode{
			Name:  name,
			Value: value,
		}
		i++
	}

	return a
}

func NewNoOpRequest() *NoOpRequest {
	return &NoOpRequest{
		Content: NoOpRequestContent{
			Urn: Urn,
		},
	}
}

func NewRenameAccountRequest(id string, newName string) *RenameAccountRequest {
	return &RenameAccountRequest{
		Content: newRenameRequestContent(id, newName),
	}
}

func NewRenameCalendarResourceRequest(id string, newName string) *RenameCalendarResourceRequest {
	return &RenameCalendarResourceRequest{
		Content: newRenameRequestContent(id, newName),
	}
}

func NewRenameCosRequest(id string, newName string) *RenameCosRequest {
	return &RenameCosRequest{
		Content: RenameCosRequestContent{
			Urn:     Urn,
			ID:      zimbraCommon.ContentString{Content: id},
			NewName: zimbraCommon.ContentString{Content: newName},
		},
	}
}

func NewRenameDistributionListRequest(id string, newName string) *RenameDistributionListRequest {
	return &RenameDistributionListRequest{
		Content: newRenameRequestContent(id, newName),
	}
}

func newRenameRequestContent(id string, newName string) RenameRequestContent {
	return RenameRequestContent{
		Urn:     Urn,
		ID:      id,
		NewName: newName,
	}
}

func NewSetPasswordRequest(id string, newPassword string) *SetPasswordRequest {
	return &SetPasswordRequest{
		Content: SetPasswordRequestContent{
			Urn:         Urn,
			ID:          id,
			NewPassword: newPassword,
		},
	}
}
