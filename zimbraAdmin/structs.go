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
	Key   string `json:"name,omitempty"`
	Value string `json:"_content,omitempty"`
}

func (anr *AttrNameResponse) ToAttrResponse() AttrResponse {
	return AttrResponse{
		Key:   anr.Key,
		Value: anr.Value,
	}
}

type ByNode struct {
	By    string `json:"by,omitempty"`
	Value string `json:"_content,omitempty"`
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
	Content AuthRequestContent `json:"AuthRequest,omitempty"`
}

type AuthRequestContent struct {
	Name     string `json:"name,attr"`
	Password string `json:"password,attr"`
	Urn      string `json:"_jsns,attr"`
}

type AuthResponse struct {
	Content AuthResponseContent `json:"AuthResponse,omitempty"`
}

type AuthResponseContent struct {
	TOKEN    []AuthResponseToken `json:"authToken"`
	Lifetime int                 `json:"lifetime"`
}

type AuthResponseToken struct {
	Content string `json:"_content"`
}

type DelegateAuthRequest struct {
	Content DelegateAuthRequestContent `json:"DelegateAuthRequest,omitempty"`
}

type DelegateAuthRequestContent struct {
	Account ByNode `json:"account,attr"`
	Urn     string `json:"_jsns,attr"`
}

type DelegateAuthResponse struct {
	Content DelegateAuthResponseContent `json:"DelegateAuthResponse,omitempty"`
}

type DelegateAuthResponseContent struct {
	TOKEN    string `json:"authToken,omitempty"`
	LifeTime int    `json:"lifetime,omitempty"`
}

type DeleteAccountRequest struct {
	Content IdRequestContent `json:"DeleteAccountRequest,omitempty"`
}

//type DeleteAccountRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//	ID  string `json:"id,omitempty"`
//}

type DeleteCalendarResourceRequest struct {
	Content IdRequestContent `json:"DeleteCalendarResourceRequest,omitempty"`
}

//type DeleteCalendarResourceRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//	ID  string `json:"id,omitempty"`
//}

type DeleteCosRequest struct {
	Content IdRequestContent `json:"DeleteCosRequest,omitempty"`
}

//type DeleteCosRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//	ID  string `json:"id,omitempty"`
//}

type DeleteDistributionListRequest struct {
	Content DeleteDistributionListRequestContent `json:"DeleteDistributionListRequest,omitempty"`
}

type DeleteDistributionListRequestContent struct {
	Urn           string `json:"_jsns,attr"`
	ID            string `json:"id,omitempty"`
	CascadeDelete int    `json:"cascadeDelete,omitempty"`
}

type DeleteDomainRequest struct {
	Content IdRequestContent `json:"DeleteDomainRequest,omitempty"`
}

type IdRequestContent struct {
	Urn string `json:"_jsns,attr"`
	ID  string `json:"id,omitempty"`
}

type DeleteGalSyncAccountRequest struct {
	Content DeleteGalSyncAccountRequestContent `json:"DeleteGalSyncAccountRequest,omitempty"`
}

type DeleteGalSyncAccountRequestContent struct {
	Account ByNode `json:"account,attr"`
	Urn     string `json:"_jsns,attr"`
}

type GetAccountRequest struct {
	Content GetAccountRequestContent `json:"GetAccountRequest,omitempty"`
}

type GetAccountRequestContent struct {
	Account ByNode `json:"account,attr"`
	Urn     string `json:"_jsns,attr"`
	Attrs   string `json:"attrs,omitempty"`
}

type GetAccountResponse struct {
	Content GetAccountResponseContent `json:"GetAccountResponse,omitempty"`
}

type GetAccountResponseContent struct {
	Account []GenericResponse `json:"account,omitempty"`
}

type GetAllConfigRequest struct {
	Content UrnRequestContent `json:"GetAllConfigRequest,omitempty"`
}

//type GetAllConfigRequestContent struct {
//	Urn string `json:"_jsns,attr"`
//}

type GetAllConfigResponse struct {
	Content GetAllConfigResponseContent `json:"GetAllConfigResponse,omitempty"`
}

type GetAllConfigResponseContent struct {
	Attrs []AttrResponse `json:"a,omitempty"`
}

type GetDistributionListRequest struct {
	Content GetDistributionListRequestContent `json:"GetDistributionListRequest,omitempty"`
}

type GetDistributionListRequestContent struct {
	Dl    ByNode `json:"dl,attr"`
	Urn   string `json:"_jsns,attr"`
	Attrs string `json:"attrs,omitempty"`
}

type GetDistributionListResponse struct {
	Content GetDistributionListResponseContent `json:"GetDistributionListResponse,omitempty"`
}

type GetDistributionListResponseContent struct {
	Dl []GenericResponse `json:"dl,omitempty"`
}

type GetLicenseRequest struct {
	Content UrnRequestContent `json:"GetLicenseRequest,omitempty"`
}

type UrnRequestContent struct {
	Urn string `json:"_jsns,attr"`
}

type GetLicenseResponse struct {
	Content GetLicenseResponseContent `json:"GetLicenseResponse,omitempty"`
}

type GetLicenseResponseContent struct {
	License    []AttrNamesResponse `json:"license,omitempty"`
	Activation []AttrNamesResponse `json:"activation,omitempty"`
	Info       []AttrNamesResponse `json:"info,omitempty"`
}

type ModifyAccountRequest struct {
	Content ModifyAccountRequestContent `json:"ModifyAccountRequest,omitempty"`
}

type ModifyAccountRequestContent struct {
	Urn   string         `json:"_jsns,attr"`
	ID    string         `json:"id,omitempty"`
	Attrs []AttrResponse `json:"a,omitempty"`
}

type ModifyAccountResponse struct {
	Content ModifyAccountResponseContent `json:"ModifyAccountResponse,omitempty"`
}

type ModifyAccountResponseContent struct {
	Account []GenericResponse `json:"account,omitempty"`
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
	Content SearchDirectoryParams `json:"SearchDirectoryRequest,omitempty"`
}

type SearchDirectoryResponse struct {
	Content SearchDirectoryResponseContent `json:"SearchDirectoryResponse,omitempty"`
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
	Content GetCalendarResourceRequestContent `json:"GetCalendarResourceRequest,omitempty"`
}

type GetCalendarResourceRequestContent struct {
	CalResource ByNode `json:"calresource,attr"`
	Urn         string `json:"_jsns,attr"`
	Attrs       string `json:"attrs,omitempty"`
}

type GetCalendarResourceResponse struct {
	Content GetCalendarResourceResponseContent `json:"GetCalendarResourceResponse,omitempty"`
}

type GetCalendarResourceResponseContent struct {
	CalResource []GenericResponse `json:"calresource,omitempty"`
}

type ModifyCalendarResourceRequest struct {
	Content ModifyCalendarResourceRequestContent `json:"ModifyCalendarResourceRequest,omitempty"`
}

type ModifyCalendarResourceRequestContent struct {
	Urn   string         `json:"_jsns,attr"`
	ID    string         `json:"id,omitempty"`
	Attrs []AttrResponse `json:"a,omitempty"`
}

type ModifyCalendarResourceResponse struct {
	Content ModifyCalendarResourceResponseContent `json:"ModifyCalendarResourceResponse,omitempty"`
}

type ModifyCalendarResourceResponseContent struct {
	CalResource []GenericResponse `json:"calresource,omitempty"`
}

// ************* SERVER ****************

type GetAllServersRequest struct {
	Content GetAllServersRequestContent `json:"GetAllServersRequest,omitempty"`
}

type GetAllServersRequestContent struct {
	Urn     string `json:"_jsns,attr"`
	Service string `json:"service,omitempty"`
}

type GetAllServersResponse struct {
	Content GetAllServersResponseContent `json:"GetAllServersResponse,omitempty"`
}

type GetAllServersResponseContent struct {
	Servers []GenericResponse `json:"servers,omitempty"`
}

type GetServerRequest struct {
	Content GetServerRequestContent `json:"GetServerRequest,omitempty"`
}

type GetServerRequestContent struct {
	Server      ByNode `json:"server,attr"`
	Urn         string `json:"_jsns,attr"`
	ApplyConfig int    `json:"applyConfig,omitempty"`
	Attrs       string `json:"attrs,omitempty"`
}

type GetServerResponse struct {
	Content GetServerResponseContent `json:"GetServerResponse,omitempty"`
}

type GetServerResponseContent struct {
	Server []GenericResponse `json:"server,omitempty"`
}

// ************* DOMAIN ****************

type GetDomainRequest struct {
	Content GetDomainRequestContent `json:"GetDomainRequest,omitempty"`
}

type GetDomainRequestContent struct {
	Domain ByNode `json:"domain,attr"`
	Urn    string `json:"_jsns,attr"`
	Attrs  string `json:"attrs,omitempty"`
}

type GetDomainResponse struct {
	Content GetDomainResponseContent `json:"GetDomainResponse,omitempty"`
}

type GetDomainResponseContent struct {
	Domains []GenericResponse `json:"domain,omitempty"`
}

type GetQuotaUsageRequest struct {
	Content GetQuotaUsageRequestContent `json:"GetQuotaUsageRequest,omitempty"`
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
	Content GetQuotaUsageResponseContent `json:"GetQuotaUsageResponse,omitempty"`
}

type GetQuotaUsageResponseContent struct {
	Account []QuotaResponse `json:"account,omitempty"`
}

type QuotaResponse struct {
	Name  string `json:"name,omitempty"`
	ID    string `json:"id,omitempty"`
	Used  int    `json:"used,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

// ************* BACKUP ****************

type BackupQueryRequest struct {
	Content BackupQueryRequestContent `json:"BackupQueryRequest,omitempty"`
}

type BackupQueryRequestContent struct {
	Urn   string            `json:"_jsns,attr"`
	Query map[string]string `json:"query"`
}

type BackupQueryResponse struct {
	Content BackupQueryResponseContent `json:"BackupQueryResponse,omitempty"`
}

type BackupQueryResponseContent struct {
	TotalSpace  int       `json:"totalSpace,omitempty"`
	FreeSpace   int       `json:"freeSpace,omitempty"`
	NameSpace   string    `json:"_jsns,omitempty"`
	Backups     []ZBackup `json:"backup,omitempty"`
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
