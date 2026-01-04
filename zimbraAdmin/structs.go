package zimbraAdmin

import "github.com/RaoH37/gosoap/zimbraCommon"

type GenericResponse struct {
	Name  string                  `json:"name,omitempty"`
	ID    string                  `json:"id,omitempty"`
	Attrs []zimbraCommon.AttrNode `json:"a,omitempty"`
}

// ************* REQUEST ***************

type AddAccountAliasRequest struct {
	Content AliasRequestContent `json:"AddAccountAliasRequest"`
}

type AliasRequestContent struct {
	Urn   string `json:"_jsns"`
	ID    string `json:"id"`
	Alias string `json:"alias"`
}

type AuthRequest struct {
	Content AuthRequestContent `json:"AuthRequest"`
}

type AuthRequestContent struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Urn      string `json:"_jsns"`
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

type CopyCosRequest struct {
	Content CopyCosRequestContent `json:"CopyCosRequest"`
}

type CopyCosRequestContent struct {
	Name zimbraCommon.ContentString `json:"name"`
	By   zimbraCommon.ByNode        `json:"cos"`
	Urn  string                     `json:"_jsns"`
}

type CopyCosResponse struct {
	Content CopyCosResponseContent `json:"CopyCosResponse"`
}

type CopyCosResponseContent struct {
	Urn string            `json:"_jsns"`
	Cos []GenericResponse `json:"cos,omitempty"`
}

type DelegateAuthRequest struct {
	Content DelegateAuthRequestContent `json:"DelegateAuthRequest"`
}

type DelegateAuthRequestContent struct {
	Account zimbraCommon.ByNode `json:"account"`
	Urn     string              `json:"_jsns"`
}

type DelegateAuthResponse struct {
	Content AuthResponseContent `json:"DelegateAuthResponse"`
}

type DeleteAccountRequest struct {
	Content IdRequestContent `json:"DeleteAccountRequest"`
}

type DeleteCalendarResourceRequest struct {
	Content IdRequestContent `json:"DeleteCalendarResourceRequest"`
}

type DeleteCosRequest struct {
	Content DeleteCosRequestContent `json:"DeleteCosRequest"`
}

type DeleteCosRequestContent struct {
	Urn string                     `json:"_jsns"`
	ID  zimbraCommon.ContentString `json:"id"`
}

type DeleteDistributionListRequest struct {
	Content DeleteDistributionListRequestContent `json:"DeleteDistributionListRequest"`
}

type DeleteDistributionListRequestContent struct {
	Urn           string             `json:"_jsns"`
	ID            string             `json:"id"`
	CascadeDelete zimbraCommon.ZBool `json:"cascadeDelete,omitempty"`
}

type DeleteDomainRequest struct {
	Content IdRequestContent `json:"DeleteDomainRequest"`
}

type IdRequestContent struct {
	Urn string `json:"_jsns"`
	ID  string `json:"id"`
}

type DeleteGalSyncAccountRequest struct {
	Content DeleteGalSyncAccountRequestContent `json:"DeleteGalSyncAccountRequest"`
}

type DeleteGalSyncAccountRequestContent struct {
	Account zimbraCommon.ByNode `json:"account"`
	Urn     string              `json:"_jsns"`
}

type GetAccountRequest struct {
	Content GetAccountRequestContent `json:"GetAccountRequest"`
}

type GetAccountRequestContent struct {
	Account  zimbraCommon.ByNode     `json:"account"`
	Urn      string                  `json:"_jsns"`
	Attrs    zimbraCommon.StringList `json:"attrs,omitempty"`
	ApplyCos zimbraCommon.ZBool      `json:"applyCos,omitempty"`
}

type GetAccountResponse struct {
	Content GetAccountResponseContent `json:"GetAccountResponse"`
}

type GetAccountResponseContent struct {
	Account []GenericResponse `json:"account"`
}

type GetAllConfigRequest struct {
	Content zimbraCommon.UrnRequestContent `json:"GetAllConfigRequest"`
}

type GetAllConfigResponse struct {
	Content GetAllConfigResponseContent `json:"GetAllConfigResponse"`
}

type GetAllConfigResponseContent struct {
	Attrs []zimbraCommon.AttrNode `json:"a"`
}

type GetDistributionListRequest struct {
	Content GetDistributionListRequestContent `json:"GetDistributionListRequest"`
}

type GetDistributionListRequestContent struct {
	Dl    zimbraCommon.ByNode     `json:"dl"`
	Urn   string                  `json:"_jsns"`
	Attrs zimbraCommon.StringList `json:"attrs,omitempty"`
}

type GetDistributionListResponse struct {
	Content GetDistributionListResponseContent `json:"GetDistributionListResponse"`
}

type GetDistributionListResponseContent struct {
	Dl []GenericResponse `json:"dl"`
}

type GetLicenseRequest struct {
	Content zimbraCommon.UrnRequestContent `json:"GetLicenseRequest"`
}

type GetLicenseResponse struct {
	Content GetLicenseResponseContent `json:"GetLicenseResponse"`
}

type GetLicenseResponseContent struct {
	License    []zimbraCommon.AttrNamesNode `json:"license"`
	Activation []zimbraCommon.AttrNamesNode `json:"activation"`
	Info       []zimbraCommon.AttrNamesNode `json:"info"`
}

type CreateAccountRequest struct {
	Content CreateAccountRequestContent `json:"CreateAccountRequest"`
}

type CreateAccountRequestContent struct {
	Urn      string                  `json:"_jsns"`
	Name     string                  `json:"name"`
	Password string                  `json:"password"`
	Attrs    []zimbraCommon.AttrNode `json:"a"`
}

type CreateAccountResponse struct {
	Content CreateAccountResponseContent `json:"CreateAccountResponse"`
}

type CreateAccountResponseContent struct {
	Account []GenericResponse `json:"account"`
}

type CreateCalendarResourceRequest struct {
	Content CreateCalendarResourceRequestContent `json:"CreateCalendarResourceRequest"`
}

type CreateCalendarResourceRequestContent struct {
	Urn      string                  `json:"_jsns"`
	Name     string                  `json:"name"`
	Password string                  `json:"password"`
	Attrs    []zimbraCommon.AttrNode `json:"a"`
}

type CreateCalendarResourceResponse struct {
	Content CreateCalendarResourceResponseContent `json:"CreateCalendarResourceResponse"`
}

type CreateCalendarResourceResponseContent struct {
	Calresource []GenericResponse `json:"calresource"`
}

type CreateCosRequest struct {
	Content CreateCosRequestContent `json:"CreateCosRequest"`
}

type CreateCosRequestContent struct {
	Urn   string                     `json:"_jsns"`
	Name  zimbraCommon.ContentString `json:"name"`
	Attrs []zimbraCommon.AttrNode    `json:"a"`
}

type CreateCosResponse struct {
	Content CreateCosResponseContent `json:"CreateCosResponse"`
}

type CreateCosResponseContent struct {
	Cos []GenericResponse `json:"cos"`
}

type CreateDistributionListRequest struct {
	Content CreateDistributionListRequestContent `json:"CreateDistributionListRequest"`
}

type CreateDistributionListRequestContent struct {
	Urn     string                  `json:"_jsns"`
	Name    string                  `json:"name"`
	Dynamic zimbraCommon.ZBool      `json:"dynamic"`
	Attrs   []zimbraCommon.AttrNode `json:"a"`
}

type CreateDistributionListResponse struct {
	Content CreateDistributionListResponseContent `json:"CreateDistributionListResponse"`
}

type CreateDistributionListResponseContent struct {
	Dl []GenericResponse `json:"dl"`
}

type CreateDomainRequest struct {
	Content CreateDomainRequestContent `json:"CreateDomainRequest"`
}

type CreateDomainRequestContent struct {
	Urn   string                  `json:"_jsns"`
	Name  string                  `json:"name"`
	Attrs []zimbraCommon.AttrNode `json:"a"`
}

type CreateDomainResponse struct {
	Content CreateDomainResponseContent `json:"CreateDomainResponse"`
}

type CreateDomainResponseContent struct {
	Domain []GenericResponse `json:"domain"`
}

type ModifyAccountRequest struct {
	Content ModifyRequestContent `json:"ModifyAccountRequest"`
}

type ModifyRequestContent struct {
	Urn   string                  `json:"_jsns"`
	ID    string                  `json:"id"`
	Attrs []zimbraCommon.AttrNode `json:"a"`
}

type ModifyAccountResponse struct {
	Content ModifyAccountResponseContent `json:"ModifyAccountResponse"`
}

type ModifyAccountResponseContent struct {
	Accounts []GenericResponse `json:"account"`
}

type ModifyCalendarResourceRequest struct {
	Content ModifyRequestContent `json:"ModifyCalendarResourceRequest"`
}

type ModifyCalendarResourceResponse struct {
	Content ModifyCalendarResourceResponseContent `json:"ModifyCalendarResourceResponse"`
}

type ModifyCalendarResourceResponseContent struct {
	Calresources []GenericResponse `json:"calresource"`
}

type ModifyCosRequest struct {
	Content ModifyCosRequestContent `json:"ModifyCosRequest"`
}

type ModifyCosRequestContent struct {
	Urn   string                     `json:"_jsns"`
	ID    zimbraCommon.ContentString `json:"id"`
	Attrs []zimbraCommon.AttrNode    `json:"a"`
}

type ModifyCosResponse struct {
	Content ModifyCosResponseContent `json:"ModifyCosResponse"`
}

type ModifyCosResponseContent struct {
	Coses []GenericResponse `json:"cos"`
}

type ModifyDistributionListRequest struct {
	Content ModifyRequestContent `json:"ModifyDistributionListRequest"`
}

type ModifyDistributionListResponse struct {
	Content ModifyDistributionListResponseContent `json:"ModifyDistributionListResponse"`
}

type ModifyDistributionListResponseContent struct {
	Dls []GenericResponse `json:"cos"`
}

type ModifyDomainRequest struct {
	Content ModifyRequestContent `json:"ModifyDomainRequest"`
}

type ModifyDomainResponse struct {
	Content ModifyDomainResponseContent `json:"ModifyDomainResponse"`
}

type ModifyDomainResponseContent struct {
	Domains []GenericResponse `json:"domain"`
}

type ModifyServerRequest struct {
	Content ModifyRequestContent `json:"ModifyServerRequest"`
}

type ModifyServerResponse struct {
	Content ModifyServerResponseContent `json:"ModifyServerResponse"`
}

type ModifyServerResponseContent struct {
	Servers []GenericResponse `json:"server"`
}

type NoOpRequest struct {
	Content NoOpRequestContent `json:"NoOpRequest"`
}

type NoOpRequestContent struct {
	Urn string `json:"_jsns"`
}

type SearchDirectoryParams struct {
	Urn           string                  `json:"_jsns"`
	Query         string                  `json:"query,omitempty"`
	MaxResults    int                     `json:"maxResults,omitempty"`
	Limit         int                     `json:"limit,omitempty"`
	Offset        int                     `json:"offset,omitempty"`
	Domain        string                  `json:"domain,omitempty"`
	ApplyCos      zimbraCommon.ZBool      `json:"applyCos,omitempty"`
	ApplyConfig   zimbraCommon.ZBool      `json:"applyConfig,omitempty"`
	SortBy        string                  `json:"sortBy,omitempty"`
	Types         SearchTypeList          `json:"types,omitempty"`
	SortAscending zimbraCommon.ZBool      `json:"sortAscending,omitempty"`
	CountOnly     zimbraCommon.ZBool      `json:"countOnly,omitempty"`
	Attrs         zimbraCommon.StringList `json:"attrs,omitempty"`
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
	Content AliasRequestContent `json:"RemoveAccountAliasRequest"`
}

type AddDistributionListAliasRequest struct {
	Content AliasRequestContent `json:"AddDistributionListAliasRequest"`
}

type RemoveDistributionListAliasRequest struct {
	Content AliasRequestContent `json:"RemoveDistributionListAliasRequest"`
}

type AddDistributionListMemberRequest struct {
	Content DistributionListMemberRequestContent `json:"AddDistributionListMemberRequest"`
}

type RemoveDistributionListMemberRequest struct {
	Content DistributionListMemberRequestContent `json:"removeDistributionListMemberRequest"`
}

type DistributionListMemberRequestContent struct {
	Urn     string                       `json:"_jsns"`
	ID      string                       `json:"id"`
	Members []zimbraCommon.ContentString `json:"dlm"`
}

type GetCalendarResourceRequest struct {
	Content GetCalendarResourceRequestContent `json:"GetCalendarResourceRequest"`
}

type GetCalendarResourceRequestContent struct {
	CalResource zimbraCommon.ByNode     `json:"calresource"`
	Urn         string                  `json:"_jsns"`
	Attrs       zimbraCommon.StringList `json:"attrs,omitempty"`
	ApplyCos    zimbraCommon.ZBool      `json:"applyCos,omitempty"`
}

type GetCalendarResourceResponse struct {
	Content GetCalendarResourceResponseContent `json:"GetCalendarResourceResponse"`
}

type GetCalendarResourceResponseContent struct {
	CalResource []GenericResponse `json:"calresource"`
}

type GetCosRequest struct {
	Content GetCosRequestContent `json:"GetCosRequest"`
}

type GetCosRequestContent struct {
	Cos   zimbraCommon.ByNode     `json:"cos"`
	Urn   string                  `json:"_jsns"`
	Attrs zimbraCommon.StringList `json:"attrs,omitempty"`
}

type GetCosResponse struct {
	Content GetCosResponseContent `json:"GetCosResponse"`
}

type GetCosResponseContent struct {
	Cos []GenericResponse `json:"cos"`
}

type GetAllServersRequest struct {
	Content GetAllServersRequestContent `json:"GetAllServersRequest"`
}

type GetAllServersRequestContent struct {
	Urn     string `json:"_jsns"`
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
	Server      zimbraCommon.ByNode     `json:"server"`
	Urn         string                  `json:"_jsns"`
	ApplyConfig zimbraCommon.ZBool      `json:"applyConfig,omitempty"`
	Attrs       zimbraCommon.StringList `json:"attrs,omitempty"`
}

type GetServerResponse struct {
	Content GetServerResponseContent `json:"GetServerResponse"`
}

type GetServerResponseContent struct {
	Server []GenericResponse `json:"server"`
}

type GetDomainRequest struct {
	Content GetDomainRequestContent `json:"GetDomainRequest"`
}

type GetDomainRequestContent struct {
	Domain      zimbraCommon.ByNode     `json:"domain"`
	Urn         string                  `json:"_jsns"`
	Attrs       zimbraCommon.StringList `json:"attrs,omitempty"`
	ApplyConfig zimbraCommon.ZBool      `json:"applyConfig,omitempty"`
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
	Urn           string             `json:"_jsns"`
	Servers       zimbraCommon.ZBool `json:"allServers,omitempty"`
	Domain        string             `json:"domain,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	SortBy        string             `json:"sortBy,omitempty"`
	SortAscending zimbraCommon.ZBool `json:"sortAscending,omitempty"`
	Refresh       zimbraCommon.ZBool `json:"refresh,omitempty"`
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

type BackupQueryRequest struct {
	Content BackupQueryRequestContent `json:"BackupQueryRequest"`
}

type BackupQueryRequestContent struct {
	Urn   string            `json:"_jsns"`
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

func (b *ZBackup) Account() *ZBackupAccount {
	return &b.Accounts[0]
}

func (b *ZBackup) Date() string {
	runes := []rune(b.Label)
	return string(runes[5:13])
}

func (a *ZBackupAccount) DiffTotalCompletionCount() int {
	return (a.Total - a.CompletionCount)
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
	Content RenameCosRequestContent `json:"RenameCosRequest"`
}

type RenameCosRequestContent struct {
	Urn     string                     `json:"_jsns"`
	ID      zimbraCommon.ContentString `json:"id"`
	NewName zimbraCommon.ContentString `json:"newName"`
}

type RenameDistributionListRequest struct {
	Content RenameRequestContent `json:"RenameDistributionListRequest"`
}

type RenameRequestContent struct {
	Urn     string `json:"_jsns"`
	ID      string `json:"id"`
	NewName string `json:"newName"`
}

type SetPasswordRequest struct {
	Content SetPasswordRequestContent `json:"SetPasswordRequest"`
}

type SetPasswordRequestContent struct {
	Urn         string `json:"_jsns"`
	ID          string `json:"id"`
	NewPassword string `json:"newPassword"`
}
