package zimbraAdmin

// ************* COMMON ***************

type ContentString struct {
	Content string `json:"_content"`
}

type AttrResponse struct {
	Key   string `json:"n,omitempty"`
	Value string `json:"_content,omitempty"`
}

type AttrNamesResponse struct {
	Attrs []AttrNameResponse `json:"attr,omitempty"`
}

func (anr *AttrNamesResponse) ToAttrsResponse() []AttrResponse {
	collection := make([]AttrResponse, len(anr.Attrs))

	for i, attr := range anr.Attrs {
		collection[i] = attr.ToAttrResponse()
	}

	return collection
}

type AttrNameResponse struct {
	Key   string `json:"name"`
	Value string `json:"_content"`
}

func (anr *AttrNameResponse) ToAttrResponse() AttrResponse {
	return AttrResponse{
		Key:   anr.Key,
		Value: anr.Value,
	}
}

type ByNode struct {
	By    string `json:"by"`
	Value string `json:"_content"`
}

type GenericResponse struct {
	Name  string         `json:"name,omitempty"`
	ID    string         `json:"id,omitempty"`
	Attrs []AttrResponse `json:"a,omitempty"`
}

// ************* REQUEST ***************

type AddAccountAliasRequest struct {
	Content AccountAliasRequestContent `json:"AddAccountAliasRequest"`
}

type AccountAliasRequestContent struct {
	Urn   string `json:"_jsns"`
	ID    string `json:"id"`
	Alias string `json:"alias"`
}

type AuthRequest struct {
	Content AuthRequestContent `json:"AuthRequest"`
}

type AuthRequestContent struct {
	Name     string `json:"name,attr"`
	Password string `json:"password,attr"`
	Urn      string `json:"_jsns,attr"`
}

type AuthResponse struct {
	Content AuthResponseContent `json:"AuthResponse"`
}

type AuthResponseContent struct {
	TOKEN    []AuthResponseToken `json:"authToken"`
	Lifetime int                 `json:"lifetime"`
}

type AuthResponseToken struct {
	Content string `json:"_content"`
}

type DelegateAuthRequest struct {
	Content DelegateAuthRequestContent `json:"DelegateAuthRequest"`
}

type DelegateAuthRequestContent struct {
	Account ByNode `json:"account,attr"`
	Urn     string `json:"_jsns,attr"`
}

type DelegateAuthResponse struct {
	Content DelegateAuthResponseContent `json:"DelegateAuthResponse"`
}

type DelegateAuthResponseContent struct {
	TOKEN    string `json:"authToken"`
	LifeTime int    `json:"lifetime,omitempty"`
}

type DeleteAccountRequest struct {
	Content IdRequestContent `json:"DeleteAccountRequest"`
}

//type DeleteAccountRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//	ID  string `json:"id,omitempty"`
//}

type DeleteCalendarResourceRequest struct {
	Content IdRequestContent `json:"DeleteCalendarResourceRequest"`
}

//type DeleteCalendarResourceRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//	ID  string `json:"id,omitempty"`
//}

type DeleteCosRequest struct {
	Content IdRequestContent `json:"DeleteCosRequest"`
}

//type DeleteCosRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//	ID  string `json:"id"`
//}

type DeleteDistributionListRequest struct {
	Content DeleteDistributionListRequestContent `json:"DeleteDistributionListRequest"`
}

type DeleteDistributionListRequestContent struct {
	Urn           string `json:"_jsns,attr"`
	ID            string `json:"id"`
	CascadeDelete int    `json:"cascadeDelete,omitempty"`
}

type DeleteDomainRequest struct {
	Content IdRequestContent `json:"DeleteDomainRequest"`
}

type IdRequestContent struct {
	Urn string `json:"_jsns,attr"`
	ID  string `json:"id"`
}

type DeleteGalSyncAccountRequest struct {
	Content DeleteGalSyncAccountRequestContent `json:"DeleteGalSyncAccountRequest"`
}

type DeleteGalSyncAccountRequestContent struct {
	Account ByNode `json:"account,attr"`
	Urn     string `json:"_jsns,attr"`
}

type GetAccountRequest struct {
	Content GetAccountRequestContent `json:"GetAccountRequest"`
}

type GetAccountRequestContent struct {
	Account ByNode `json:"account,attr"`
	Urn     string `json:"_jsns,attr"`
	Attrs   string `json:"attrs,omitempty"`
}

type GetAccountResponse struct {
	Content GetAccountResponseContent `json:"GetAccountResponse"`
}

type GetAccountResponseContent struct {
	Account []GenericResponse `json:"account"`
}

type GetAllConfigRequest struct {
	Content UrnRequestContent `json:"GetAllConfigRequest"`
}

//type GetAllConfigRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//}

type GetAllConfigResponse struct {
	Content GetAllConfigResponseContent `json:"GetAllConfigResponse"`
}

type GetAllConfigResponseContent struct {
	Attrs []AttrResponse `json:"a"`
}

type GetDistributionListRequest struct {
	Content GetDistributionListRequestContent `json:"GetDistributionListRequest"`
}

type GetDistributionListRequestContent struct {
	Dl    ByNode `json:"dl,attr"`
	Urn   string `json:"_jsns,attr"`
	Attrs string `json:"attrs,omitempty"`
}

type GetDistributionListResponse struct {
	Content GetDistributionListResponseContent `json:"GetDistributionListResponse"`
}

type GetDistributionListResponseContent struct {
	Dl []GenericResponse `json:"dl"`
}

type GetLicenseRequest struct {
	Content UrnRequestContent `json:"GetLicenseRequest"`
}

type UrnRequestContent struct {
	Urn string `json:"_jsns,attr"`
}

type GetLicenseResponse struct {
	Content GetLicenseResponseContent `json:"GetLicenseResponse"`
}

type GetLicenseResponseContent struct {
	License    []AttrNamesResponse `json:"license"`
	Activation []AttrNamesResponse `json:"activation"`
	Info       []AttrNamesResponse `json:"info"`
}

type ModifyAccountRequest struct {
	Content ModifyAccountRequestContent `json:"ModifyAccountRequest"`
}

type ModifyAccountRequestContent struct {
	Urn   string         `json:"_jsns,attr"`
	ID    string         `json:"id"`
	Attrs []AttrResponse `json:"a"`
}

type ModifyAccountResponse struct {
	Content ModifyAccountResponseContent `json:"ModifyAccountResponse"`
}

type ModifyAccountResponseContent struct {
	Account []GenericResponse `json:"account"`
}

type ModifyCalendarResourceRequest struct {
	Content ModifyCalendarResourceRequestContent `json:"ModifyCalendarResourceRequest"`
}

type ModifyCalendarResourceRequestContent struct {
	Urn   string         `json:"_jsns,attr"`
	ID    string         `json:"id"`
	Attrs []AttrResponse `json:"a"`
}

type ModifyCalendarResourceResponse struct {
	Content ModifyCalendarResourceResponseContent `json:"ModifyCalendarResourceResponse"`
}

type ModifyCalendarResourceResponseContent struct {
	Calresource []GenericResponse `json:"calresource"`
}

type ModifyCosRequest struct {
	Content ModifyCosRequestContent `json:"ModifyCosRequest"`
}

type ModifyCosRequestContent struct {
	Urn   string         `json:"_jsns,attr"`
	ID    string         `json:"id"`
	Attrs []AttrResponse `json:"a"`
}

type ModifyCosResponse struct {
	Content ModifyCosResponseContent `json:"ModifyCosResponse"`
}

type ModifyCosResponseContent struct {
	Cos []GenericResponse `json:"cos"`
}

type ModifyDistributionListRequest struct {
	Content ModifyDistributionListRequestContent `json:"ModifyDistributionListRequest"`
}

type ModifyDistributionListRequestContent struct {
	Urn   string         `json:"_jsns,attr"`
	ID    string         `json:"id"`
	Attrs []AttrResponse `json:"a"`
}

type ModifyDistributionListResponse struct {
	Content ModifyDistributionListResponseContent `json:"ModifyDistributionListResponse"`
}

type ModifyDistributionListResponseContent struct {
	Dl []GenericResponse `json:"cos"`
}

type ModifyDomainRequest struct {
	Content ModifyDomainRequestContent `json:"ModifyDomainRequest"`
}

type ModifyDomainRequestContent struct {
	Urn   string         `json:"_jsns,attr"`
	ID    string         `json:"id"`
	Attrs []AttrResponse `json:"a"`
}

type ModifyDomainResponse struct {
	Content ModifyDomainResponseContent `json:"ModifyDomainResponse"`
}

type ModifyDomainResponseContent struct {
	Domain []GenericResponse `json:"domain"`
}

type SearchDirectoryParams struct {
	Urn           string `json:"_jsns,attr"`
	Query         string `json:"query,omitempty"`
	MaxResults    int    `json:"maxResults,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Domain        string `json:"domain,omitempty"`
	ApplyCos      int    `json:"applyCos,omitempty"`
	ApplyConfig   int    `json:"applyConfig,omitempty"`
	SortBy        string `json:"sortBy,omitempty"`
	Types         string `json:"types,omitempty"`
	SortAscending int    `json:"sortAscending,omitempty"`
	CountOnly     int    `json:"countOnly,omitempty"`
	Attrs         string `json:"attrs,omitempty"`
}

type SearchDirectoryRequest struct {
	Content SearchDirectoryParams `json:"SearchDirectoryRequest"`
}

type SearchDirectoryResponse struct {
	Content SearchDirectoryResponseContent `json:"SearchDirectoryResponse"`
}

type SearchDirectoryResponseContent struct {
	Count        int               `json:"num,omitempty"`
	Accounts     []GenericResponse `json:"account,omitempty"`
	Dls          []GenericResponse `json:"dl,omitempty"`
	Domains      []GenericResponse `json:"domain,omitempty"`
	CalResources []GenericResponse `json:"calresource,omitempty"`
	Coses        []GenericResponse `json:"cos,omitempty"`
}

type RemoveAccountAliasRequest struct {
	Content AccountAliasRequestContent `json:"RemoveAccountAliasRequest"`
}

//type RemoveAccountAliasRequestContent struct {
//	Urn   string `json:"_jsns"`
//	ID    string `json:"id"`
//	Alias string `json:"alias"`
//}

// ************* DISTRIBUTION LIST ****************

type AddDistributionListAliasRequest struct {
	Content DistributionListAliasRequestContent `json:"AddDistributionListAliasRequest"`
}

type RemoveDistributionListAliasRequest struct {
	Content DistributionListAliasRequestContent `json:"RemoveDistributionListAliasRequest"`
}

type DistributionListAliasRequestContent struct {
	Urn   string `json:"_jsns"`
	ID    string `json:"id"`
	Alias string `json:"alias"`
}

type AddDistributionListMemberRequest struct {
	Content DistributionListMemberRequestContent `json:"AddDistributionListMemberRequest"`
}

type RemoveDistributionListMemberRequest struct {
	Content DistributionListMemberRequestContent `json:"removeDistributionListMemberRequest"`
}

type DistributionListMemberRequestContent struct {
	Urn     string          `json:"_jsns"`
	ID      string          `json:"id"`
	Members []ContentString `json:"dlm"`
}

// ************* RESOURCE ****************

type GetCalendarResourceRequest struct {
	Content GetCalendarResourceRequestContent `json:"GetCalendarResourceRequest"`
}

type GetCalendarResourceRequestContent struct {
	CalResource ByNode `json:"calresource,attr"`
	Urn         string `json:"_jsns,attr"`
	Attrs       string `json:"attrs,omitempty"`
}

type GetCalendarResourceResponse struct {
	Content GetCalendarResourceResponseContent `json:"GetCalendarResourceResponse"`
}

type GetCalendarResourceResponseContent struct {
	CalResource []GenericResponse `json:"calresource"`
}

// ************* SERVER ****************

type GetAllServersRequest struct {
	Content GetAllServersRequestContent `json:"GetAllServersRequest"`
}

type GetAllServersRequestContent struct {
	Urn     string `json:"_jsns,attr"`
	Service string `json:"service,omitempty"`
}

type GetAllServersResponse struct {
	Content GetAllServersResponseContent `json:"GetAllServersResponse"`
}

type GetAllServersResponseContent struct {
	Servers []GenericResponse `json:"servers"`
}

type GetServerRequest struct {
	Content GetServerRequestContent `json:"GetServerRequest"`
}

type GetServerRequestContent struct {
	Server      ByNode `json:"server,attr"`
	Urn         string `json:"_jsns,attr"`
	ApplyConfig int    `json:"applyConfig,omitempty"`
	Attrs       string `json:"attrs,omitempty"`
}

type GetServerResponse struct {
	Content GetServerResponseContent `json:"GetServerResponse"`
}

type GetServerResponseContent struct {
	Server []GenericResponse `json:"server"`
}

// ************* DOMAIN ****************

type GetDomainRequest struct {
	Content GetDomainRequestContent `json:"GetDomainRequest"`
}

type GetDomainRequestContent struct {
	Domain ByNode `json:"domain,attr"`
	Urn    string `json:"_jsns,attr"`
	Attrs  string `json:"attrs,omitempty"`
}

type GetDomainResponse struct {
	Content GetDomainResponseContent `json:"GetDomainResponse"`
}

type GetDomainResponseContent struct {
	Domains []GenericResponse `json:"domain"`
}

type GetQuotaUsageRequest struct {
	Content GetQuotaUsageRequestContent `json:"GetQuotaUsageRequest"`
}

type GetQuotaUsageRequestContent struct {
	Urn           string `json:"_jsns,attr"`
	Servers       int    `json:"allServers,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	SortBy        string `json:"sortBy,omitempty"`
	SortAscending int    `json:"sortAscending,omitempty"`
	Refresh       int    `json:"refresh,omitempty"`
}

type GetQuotaUsageResponse struct {
	Content GetQuotaUsageResponseContent `json:"GetQuotaUsageResponse"`
}

type GetQuotaUsageResponseContent struct {
	Account []QuotaResponse `json:"account"`
}

type QuotaResponse struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Used  int    `json:"used"`
	Limit int    `json:"limit"`
}

// ************* BACKUP ****************

type BackupQueryRequest struct {
	Content BackupQueryRequestContent `json:"BackupQueryRequest"`
}

type BackupQueryRequestContent struct {
	Urn   string            `json:"_jsns,attr"`
	Query map[string]string `json:"query"`
}

type BackupQueryResponse struct {
	Content BackupQueryResponseContent `json:"BackupQueryResponse"`
}

type BackupQueryResponseContent struct {
	TotalSpace  int       `json:"totalSpace"`
	FreeSpace   int       `json:"freeSpace"`
	NameSpace   string    `json:"_jsns"`
	Backups     []ZBackup `json:"backup"`
	MaxAccounts int
}

type ZBackup struct {
	Label      string           `json:"label,omitempty"`
	Type       string           `json:"type,omitempty"`
	Aborted    bool             `json:"aborted,omitempty"`
	Start      int              `json:"start,omitempty"`
	End        int              `json:"end,omitempty"`
	MinRedoSeq int              `json:"minRedoSeq,omitempty"`
	MaxRedoSeq int              `json:"maxRedoSeq,omitempty"`
	Live       bool             `json:"live,omitempty"`
	Accounts   []ZBackupAccount `json:"accounts,omitempty"`
}

type ZBackupAccount struct {
	Total           int `json:"total,omitempty"`
	CompletionCount int `json:"completionCount,omitempty"`
}

type RenameAccountRequest struct {
	Content RenameRequestContent `json:"RenameAccountRequest"`
}

type RenameCalendarResourceRequest struct {
	Content RenameRequestContent `json:"RenameCalendarResourceRequest"`
}

type RenameCosRequest struct {
	Content RenameRequestContent `json:"RenameCosRequest"`
}

type RenameDistributionListRequest struct {
	Content RenameRequestContent `json:"RenameDistributionListRequest"`
}

type RenameRequestContent struct {
	ID      string `json:"id"`
	NewName string `json:"newName"`
}

type SetPasswordRequest struct {
	Content SetPasswordRequestContent `json:"SetPasswordRequest"`
}

type SetPasswordRequestContent struct {
	ID          string `json:"id"`
	NewPassword string `json:"newPassword"`
}
