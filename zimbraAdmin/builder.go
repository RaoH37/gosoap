package zimbraAdmin

import "strings"

const urnAdmin = "urn:zimbraAdmin"

func NewAuthRequest(name string, password string) (*AuthRequest, AuthResponse) {
	r := &AuthRequest{
		Content: AuthRequestContent{
			Name:     name,
			Password: password,
			Urn:      urnAdmin,
		},
	}

	return r, AuthResponse{}
}

func NewByNode(by string, value string) ByNode {
	return ByNode{
		By:    by,
		Value: value,
	}
}

func NewGetAccountRequest(by ByNode, attrs []string) (*GetAccountRequest, GetAccountResponse) {
	r := &GetAccountRequest{
		Content: GetAccountRequestContent{
			Urn:     urnAdmin,
			Account: by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, GetAccountResponse{}
}

func NewGetAllConfigRequest(by ByNode, attrs []string) (*GetAllConfigRequest, GetAllConfigResponse) {
	r := &GetAllConfigRequest{
		Content: newUrnRequestContent(),
	}

	return r, GetAllConfigResponse{}
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

func NewGetCalendarResourceRequest(by ByNode, attrs []string) (*GetCalendarResourceRequest, GetCalendarResourceResponse) {
	r := &GetCalendarResourceRequest{
		Content: GetCalendarResourceRequestContent{
			Urn:         urnAdmin,
			CalResource: by,
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

func NewGetDomainRequest(by ByNode, attrs []string) (*GetDomainRequest, GetDomainResponse) {
	r := &GetDomainRequest{
		Content: GetDomainRequestContent{
			Urn:    urnAdmin,
			Domain: by,
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
	r := &GetQuotaUsageRequest{
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
	}

	return r, GetQuotaUsageResponse{}
}

func NewBackupQueryRequest() (*BackupQueryRequest, BackupQueryResponse) {
	r := &BackupQueryRequest{
		Content: BackupQueryRequestContent{
			Urn:   urnAdmin,
			Query: make(map[string]string),
		},
	}

	return r, BackupQueryResponse{}
}

func NewCopyCosRequest(by ByNode, newName string) (*CopyCosRequest, CopyCosResponse) {
	r := &CopyCosRequest{
		Content: CopyCosRequestContent{
			Urn: urnAdmin,
			By:  by,
			Name: ContentString{
				Content: newName,
			},
		},
	}

	return r, CopyCosResponse{}
}

func NewCreateAccountRequest(name string, password string, attrs map[string]string) (*CreateAccountRequest, CreateAccountResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &CreateAccountRequest{
		Content: CreateAccountRequestContent{
			Urn:      urnAdmin,
			Name:     name,
			Password: password,
			Attrs:    a,
		},
	}

	return r, CreateAccountResponse{}
}

func NewCreateCalendarResourceRequest(name string, password string, attrs map[string]string) (*CreateCalendarResourceRequest, CreateCalendarResourceResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &CreateCalendarResourceRequest{
		Content: CreateCalendarResourceRequestContent{
			Urn:      urnAdmin,
			Name:     name,
			Password: password,
			Attrs:    a,
		},
	}

	return r, CreateCalendarResourceResponse{}
}

func NewCreateCosRequest(name string, attrs map[string]string) (*CreateCosRequest, CreateCosResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &CreateCosRequest{
		Content: CreateCosRequestContent{
			Urn: urnAdmin,
			Name: ContentString{
				Content: name,
			},
			Attrs: a,
		},
	}

	return r, CreateCosResponse{}
}

func NewCreateDistributionListRequest(name string, dynamic int8, attrs map[string]string) (*CreateDistributionListRequest, CreateDistributionListResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &CreateDistributionListRequest{
		Content: CreateDistributionListRequestContent{
			Urn:     urnAdmin,
			Name:    name,
			Dynamic: dynamic,
			Attrs:   a,
		},
	}

	return r, CreateDistributionListResponse{}
}

func NewCreateDomainRequest(name string, attrs map[string]string) (*CreateDomainRequest, CreateDomainResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &CreateDomainRequest{
		Content: CreateDomainRequestContent{
			Urn:   urnAdmin,
			Name:  name,
			Attrs: a,
		},
	}

	return r, CreateDomainResponse{}
}

func NewDeleteCalendarResourceRequest(id string) *DeleteCalendarResourceRequest {
	r := &DeleteCalendarResourceRequest{
		Content: newIdRequestContent(id),
	}

	return r
}

func NewDeleteCosRequest(id string) *DeleteCosRequest {
	r := &DeleteCosRequest{
		Content: DeleteCosRequestContent{
			Urn: urnAdmin,
			ID: ContentString{
				Content: id,
			},
		},
	}

	return r
}

func NewDeleteDistributionListRequest(id string, cascadeDelete int8) *DeleteDistributionListRequest {
	content := DeleteDistributionListRequestContent{
		Urn:           urnAdmin,
		ID:            id,
		CascadeDelete: cascadeDelete,
	}

	r := &DeleteDistributionListRequest{
		Content: content,
	}

	return r
}

func NewDeleteDomainRequest(id string) *DeleteDomainRequest {
	r := &DeleteDomainRequest{
		Content: newIdRequestContent(id),
	}

	return r
}

func NewDeleteAccountRequest(id string) *DeleteAccountRequest {
	r := &DeleteAccountRequest{
		Content: newIdRequestContent(id),
	}

	return r
}

func newIdRequestContent(id string) IdRequestContent {
	return IdRequestContent{
		Urn: urnAdmin,
		ID:  id,
	}
}

func NewDelegateAuthRequest(by ByNode) (*DelegateAuthRequest, DelegateAuthResponse) {
	r := &DelegateAuthRequest{
		Content: DelegateAuthRequestContent{
			Urn:     urnAdmin,
			Account: by,
		},
	}

	return r, DelegateAuthResponse{}
}

func NewDeleteGalSyncAccountRequest(by ByNode) *DeleteGalSyncAccountRequest {
	r := &DeleteGalSyncAccountRequest{
		Content: DeleteGalSyncAccountRequestContent{
			Urn:     urnAdmin,
			Account: by,
		},
	}

	return r
}

func NewSearchDirectoryRequest(params *SearchDirectoryParams) (*SearchDirectoryRequest, SearchDirectoryResponse) {
	r := &SearchDirectoryRequest{
		Content: *params,
	}

	return r, SearchDirectoryResponse{}
}

func NewLicenseRequest() (*GetLicenseRequest, GetLicenseResponse) {
	r := &GetLicenseRequest{
		Content: newUrnRequestContent(),
	}

	return r, GetLicenseResponse{}
}

func newUrnRequestContent() UrnRequestContent {
	return UrnRequestContent{
		Urn: urnAdmin,
	}
}

func NewModifyAccountRequest(id string, attrs map[string]string) (*ModifyAccountRequest, ModifyAccountResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &ModifyAccountRequest{
		Content: ModifyRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Attrs: a,
		},
	}

	return r, ModifyAccountResponse{}
}

func NewAddAccountAliasRequest(id string, alias string) *AddAccountAliasRequest {
	r := &AddAccountAliasRequest{
		Content: newAccountAliasRequestContent(id, alias),
	}

	return r
}

func NewRemoveAccountAliasRequest(id string, alias string) *RemoveAccountAliasRequest {
	r := &RemoveAccountAliasRequest{
		Content: newAccountAliasRequestContent(id, alias),
	}

	return r
}

func newAccountAliasRequestContent(id string, alias string) AliasRequestContent {
	return AliasRequestContent{
		Urn:   urnAdmin,
		ID:    id,
		Alias: alias,
	}
}

func NewModifyCalendarResourceRequest(id string, attrs map[string]string) (*ModifyCalendarResourceRequest, ModifyCalendarResourceResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &ModifyCalendarResourceRequest{
		Content: ModifyRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Attrs: a,
		},
	}

	return r, ModifyCalendarResourceResponse{}
}

func NewAddDistributionListAliasRequest(id string, alias string) *AddDistributionListAliasRequest {
	r := &AddDistributionListAliasRequest{
		Content: newDistributionListAliasRequestContent(id, alias),
	}

	return r
}

func NewRemoveDistributionListAliasRequest(id string, alias string) *RemoveDistributionListAliasRequest {
	r := &RemoveDistributionListAliasRequest{
		Content: newDistributionListAliasRequestContent(id, alias),
	}

	return r
}

func newDistributionListAliasRequestContent(id string, alias string) AliasRequestContent {
	return AliasRequestContent{
		Urn:   urnAdmin,
		ID:    id,
		Alias: alias,
	}
}

func NewAddDistributionListMemberRequest(id string, members []string) *AddDistributionListMemberRequest {
	r := &AddDistributionListMemberRequest{
		Content: newDistributionListMemberRequestContent(id, members),
	}

	return r
}

func NewRemoveDistributionListMemberRequest(id string, members []string) *RemoveDistributionListMemberRequest {
	r := &RemoveDistributionListMemberRequest{
		Content: newDistributionListMemberRequestContent(id, members),
	}

	return r
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
	r := &GetLicenseRequest{
		Content: newUrnRequestContent(),
	}

	return r, GetLicenseResponse{}
}

func NewModifyCosRequest(id string, attrs map[string]string) (*ModifyCosRequest, ModifyCosResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &ModifyCosRequest{
		Content: ModifyRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Attrs: a,
		},
	}

	return r, ModifyCosResponse{}
}

func NewModifyDistributionListRequest(id string, attrs map[string]string) (*ModifyDistributionListRequest, ModifyDistributionListResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &ModifyDistributionListRequest{
		Content: ModifyRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Attrs: a,
		},
	}

	return r, ModifyDistributionListResponse{}
}

func NewModifyDomainRequest(id string, attrs map[string]string) (*ModifyDomainRequest, ModifyDomainResponse) {
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key:   key,
			Value: value,
		})
	}

	r := &ModifyDomainRequest{
		Content: ModifyRequestContent{
			Urn:   urnAdmin,
			ID:    id,
			Attrs: a,
		},
	}

	return r, ModifyDomainResponse{}
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
		Content: newRenameRequestContent(id, newName),
	}
}

func NewRenameDistributionListRequest(id string, newName string) *RenameDistributionListRequest {
	return &RenameDistributionListRequest{
		Content: newRenameRequestContent(id, newName),
	}
}

func newRenameRequestContent(id string, newName string) RenameRequestContent {
	return RenameRequestContent{
		ID:      id,
		NewName: newName,
	}
}

func NewSetPasswordRequest(id string, newPassword string) *SetPasswordRequest {
	return &SetPasswordRequest{
		Content: SetPasswordRequestContent{
			ID:          id,
			NewPassword: newPassword,
		},
	}
}
