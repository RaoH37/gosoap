package zimbraAdmin

import "strings"

const urnAdmin = "urn:zimbraAdmin"

func NewAuthRequest(name string, password string) (*AuthRequest, AuthResponse) {
	return &AuthRequest{
		Content: AuthRequestContent{
			Name:     name,
			Password: password,
			Urn:      urnAdmin,
		},
	}, AuthResponse{}
}

func NewByNode(by string, value string) ByNode {
	return ByNode{
		By:    by,
		Value: value,
	}
}

func NewGetAccountRequest(by ByNode, attrs []string, applyCos int8) (*GetAccountRequest, GetAccountResponse) {
	r := &GetAccountRequest{
		Content: GetAccountRequestContent{
			Urn:      urnAdmin,
			Account:  by,
			ApplyCos: applyCos,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetAccountResponse{}
}

func NewGetAllConfigRequest(by ByNode, attrs []string) (*GetAllConfigRequest, GetAllConfigResponse) {
	return &GetAllConfigRequest{
		Content: newUrnRequestContent(),
	}, GetAllConfigResponse{}
}

func NewGetDistributionListRequest(by ByNode, attrs []string) (*GetDistributionListRequest, GetDistributionListResponse) {
	r := &GetDistributionListRequest{
		Content: GetDistributionListRequestContent{
			Urn: urnAdmin,
			Dl:  by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetDistributionListResponse{}
}

func NewGetCalendarResourceRequest(by ByNode, attrs []string, applyCos int8) (*GetCalendarResourceRequest, GetCalendarResourceResponse) {
	r := &GetCalendarResourceRequest{
		Content: GetCalendarResourceRequestContent{
			Urn:         urnAdmin,
			CalResource: by,
			ApplyCos:    applyCos,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetCalendarResourceResponse{}
}

func NewGetCosRequest(by ByNode, attrs []string) (*GetCosRequest, GetCosResponse) {
	r := &GetCosRequest{
		Content: GetCosRequestContent{
			Urn: urnAdmin,
			Cos: by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetCosResponse{}
}

func NewGetDomainRequest(by ByNode, attrs []string, applyConfig int8) (*GetDomainRequest, GetDomainResponse) {
	r := &GetDomainRequest{
		Content: GetDomainRequestContent{
			Urn:         urnAdmin,
			Domain:      by,
			ApplyConfig: applyConfig,
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
			Urn: urnAdmin,
		},
	}

	if service != "" {
		r.Content.Service = service
	}

	return r, GetAllServersResponse{}
}

func NewGetServerRequest(by ByNode, applyConfig int8, attrs []string) (*GetServerRequest, GetServerResponse) {
	r := &GetServerRequest{
		Content: GetServerRequestContent{
			Urn:         urnAdmin,
			ApplyConfig: applyConfig,
			Server:      by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetServerResponse{}
}

func NewGetQuotaUsageRequest(domain string, allServers int8, limit int, offset int, sortBy string, sortAscending int8, refresh int8) (*GetQuotaUsageRequest, GetQuotaUsageResponse) {
	return &GetQuotaUsageRequest{
		Content: GetQuotaUsageRequestContent{
			Urn:           urnAdmin,
			Servers:       allServers,
			Domain:        domain,
			Limit:         limit,
			Offset:        offset,
			SortBy:        sortBy,
			SortAscending: sortAscending,
			Refresh:       refresh,
		},
	}, GetQuotaUsageResponse{}
}

func NewBackupQueryRequest() (*BackupQueryRequest, BackupQueryResponse) {
	return &BackupQueryRequest{
		Content: BackupQueryRequestContent{
			Urn:   urnAdmin,
			Query: make(map[string]string),
		},
	}, BackupQueryResponse{}
}

func NewCopyCosRequest(by ByNode, newName string) (*CopyCosRequest, CopyCosResponse) {
	return &CopyCosRequest{
		Content: CopyCosRequestContent{
			Urn: urnAdmin,
			By:  by,
			Name: ContentString{
				Content: newName,
			},
		},
	}, CopyCosResponse{}
}

func NewCreateAccountRequest(name string, password string, attrs map[string]string) (*CreateAccountRequest, CreateAccountResponse) {
	return &CreateAccountRequest{
		Content: CreateAccountRequestContent{
			Urn:      urnAdmin,
			Name:     name,
			Password: password,
			Attrs:    buildAttrResponses(attrs),
		},
	}, CreateAccountResponse{}
}

func NewCreateCalendarResourceRequest(name string, password string, attrs map[string]string) (*CreateCalendarResourceRequest, CreateCalendarResourceResponse) {
	return &CreateCalendarResourceRequest{
		Content: CreateCalendarResourceRequestContent{
			Urn:      urnAdmin,
			Name:     name,
			Password: password,
			Attrs:    buildAttrResponses(attrs),
		},
	}, CreateCalendarResourceResponse{}
}

func NewCreateCosRequest(name string, attrs map[string]string) (*CreateCosRequest, CreateCosResponse) {
	return &CreateCosRequest{
		Content: CreateCosRequestContent{
			Urn: urnAdmin,
			Name: ContentString{
				Content: name,
			},
			Attrs: buildAttrResponses(attrs),
		},
	}, CreateCosResponse{}
}

func NewCreateDistributionListRequest(name string, dynamic int8, attrs map[string]string) (*CreateDistributionListRequest, CreateDistributionListResponse) {
	return &CreateDistributionListRequest{
		Content: CreateDistributionListRequestContent{
			Urn:     urnAdmin,
			Name:    name,
			Dynamic: dynamic,
			Attrs:   buildAttrResponses(attrs),
		},
	}, CreateDistributionListResponse{}
}

func NewCreateDomainRequest(name string, attrs map[string]string) (*CreateDomainRequest, CreateDomainResponse) {
	return &CreateDomainRequest{
		Content: CreateDomainRequestContent{
			Urn:   urnAdmin,
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
			Urn: urnAdmin,
			ID: ContentString{
				Content: id,
			},
		},
	}
}

func NewDeleteDistributionListRequest(id string, cascadeDelete int8) *DeleteDistributionListRequest {
	return &DeleteDistributionListRequest{
		Content: DeleteDistributionListRequestContent{
			Urn:           urnAdmin,
			ID:            id,
			CascadeDelete: cascadeDelete,
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
		Urn: urnAdmin,
		ID:  id,
	}
}

func NewDelegateAuthRequest(by ByNode) (*DelegateAuthRequest, DelegateAuthResponse) {
	return &DelegateAuthRequest{
		Content: DelegateAuthRequestContent{
			Urn:     urnAdmin,
			Account: by,
		},
	}, DelegateAuthResponse{}
}

func NewDeleteGalSyncAccountRequest(by ByNode) *DeleteGalSyncAccountRequest {
	return &DeleteGalSyncAccountRequest{
		Content: DeleteGalSyncAccountRequestContent{
			Urn:     urnAdmin,
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

func newUrnRequestContent() UrnRequestContent {
	return UrnRequestContent{
		Urn: urnAdmin,
	}
}

func newModifyRequestContent(id string, a []AttrResponse) ModifyRequestContent {
	return ModifyRequestContent{
		Urn:   urnAdmin,
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
		Urn:   urnAdmin,
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
		Urn:   urnAdmin,
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
		Urn:     urnAdmin,
		ID:      id,
		Members: convertToContentStrings(members),
	}
}

func convertToContentStrings(arr []string) []ContentString {
	contentStrings := make([]ContentString, len(arr))

	for i, a := range arr {
		contentStrings[i] = ContentString{Content: a}
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
			Urn:   urnAdmin,
			ID:    ContentString{Content: id},
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

func buildAttrResponses(attrs map[string]string) []AttrResponse {
	a := make([]AttrResponse, len(attrs))

	i := 0
	for key, value := range attrs {
		a[i] = AttrResponse{
			Key:   key,
			Value: value,
		}
		i++
	}

	return a
}

func NewNoOpRequest() *NoOpRequest {
	return &NoOpRequest{
		Content: NoOpRequestContent{
			Urn: urnAdmin,
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
			Urn:     urnAdmin,
			ID:      ContentString{Content: id},
			NewName: ContentString{Content: newName},
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
		Urn:     urnAdmin,
		ID:      id,
		NewName: newName,
	}
}

func NewSetPasswordRequest(id string, newPassword string) *SetPasswordRequest {
	return &SetPasswordRequest{
		Content: SetPasswordRequestContent{
			Urn:         urnAdmin,
			ID:          id,
			NewPassword: newPassword,
		},
	}
}
