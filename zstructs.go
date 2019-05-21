package zsoap

type AttrResponse struct {
	Key   string `json:"n,omitempty"`
	Value string `json:"_content,omitempty"`
}

type ByRequest struct {
	By    string `json:"by,omitempty"`
	Value string `json:"_content,omitempty"`
}

type GenericResponse struct {
	Name  string         `json:"name,omitempty"`
	ID    string         `json:"id,omitempty"`
	Attrs []AttrResponse `json:"a,omitempty"`
}

// ************* LOGIN ****************

type AuthRequest struct {
	Content AuthRequestContent `json:"AuthRequest,omitempty"`
}

type AuthRequestContent struct {
	Name string `json:"name,attr"`
	Password string `json:"password,attr"`
	Urn string `json:"_jsns,attr"`
}

type AuthResponse struct {
	Content AuthResponseContent `json:"AuthResponse,omitempty"`
}

type AuthResponseContent struct {
	TOKEN []AuthResponseToken `json:"authToken"`
	Lifetime int `json:"lifetime"`
}

type AuthResponseToken struct {
	Content string `json:"_content"`
}

// ************* ACCOUNT ****************

type GetAllAccountsRequest struct {
	Content GetAllAccountsRequestContent `json:"GetAllAccountsRequest,omitempty"`
}

type GetAllAccountsRequestContent struct {
	Urn string `json:"_jsns,attr"`
	Server interface{} `json:"server,omitempty"`
	Domain interface{} `json:"domain,omitempty"`
}

type GetAllAccountsResponse struct {
	Content GetAllAccountsResponseContent `json:"GetAllAccountsResponse,omitempty"`
}

type GetAllAccountsResponseContent struct {
	Account []GenericResponse `json:"account,omitempty"`
	// Account []AccountResponse `json:"account,omitempty"`
}

type GetAccountRequest struct {
	Content GetAccountRequestContent `json:"GetAccountRequest,omitempty"`
}

type GetAccountRequestContent struct {
	Account ByRequest `json:"account,attr"`
	Urn string `json:"_jsns,attr"`
	Attrs string `json:"attrs,omitempty"`
}

type GetAccountResponse struct {
	Content GetAccountResponseContent `json:"GetAccountResponse,omitempty"`
}

type GetAccountResponseContent struct {
	// Account []AccountResponse `json:"account,omitempty"`
	Account []GenericResponse `json:"account,omitempty"`
}

// type AccountResponse struct {
// 	Name string `json:"name,omitempty"`
// 	ID string `json:"id,omitempty"`
// 	Attrs []AttrResponse `json:"a,omitempty"`
// }

type ModifyAccountRequest struct {
	Content ModifyAccountRequestContent `json:"ModifyAccountRequest,omitempty"`
}

type ModifyAccountRequestContent struct {
	Urn string `json:"_jsns,attr"`
	ID string `json:"id,omitempty"`
	Attrs []AttrResponse `json:"a,omitempty"`
}

type ModifyAccountResponse struct {
	Content ModifyAccountResponseContent `json:"ModifyAccountResponse,omitempty"`
}

type ModifyAccountResponseContent struct {
	// Account []AccountResponse `json:"account,omitempty"`
	Account []GenericResponse `json:"account,omitempty"`
}

// ************* SERVER ****************

type GetAllServersRequest struct {
	Content GetAllServersRequestContent `json:"GetAllServersRequest,omitempty"`
}

type GetAllServersRequestContent struct {
	Urn string `json:"_jsns,attr"`
	Service string `json:"service,omitempty"`
}

type GetAllServersResponse struct {
	Content GetAllServersResponseContent `json:"GetAllServersResponse,omitempty"`
}

type GetAllServersResponseContent struct {
	// Server []ServerResponse `json:"server,omitempty"`
	Server []GenericResponse `json:"server,omitempty"`
}

// type ServerResponse struct {
// 	Name string `json:"name,omitempty"`
// 	ID string `json:"id,omitempty"`
// 	Attrs []AttrResponse `json:"a,omitempty"`
// }

// ************* DOMAIN ****************

type GetAllDomainsRequest struct {
	Content GetAllDomainsRequestContent `json:"GetAllDomainsRequest,omitempty"`
}

type GetAllDomainsRequestContent struct {
	Urn string `json:"_jsns,attr"`
}

type GetAllDomainsResponse struct {
	Content GetAllDomainsResponseContent `json:"GetAllDomainsResponse,omitempty"`
}

type GetAllDomainsResponseContent struct {
	// Domain []DomainResponse `json:"domain,omitempty"`
	Domain []GenericResponse `json:"domain,omitempty"`
}

// type DomainResponse struct {
// 	Name string `json:"name,omitempty"`
// 	ID string `json:"id,omitempty"`
// 	Attrs []AttrResponse `json:"a,omitempty"`
// }

// ************* COS ****************

type GetAllCosRequest struct {
	Content GetAllCosRequestContent `json:"GetAllCosRequest,omitempty"`
}

type GetAllCosRequestContent struct {
	Urn string `json:"_jsns,attr"`
}

type GetAllCosResponse struct {
	Content GetAllCosResponseContent `json:"GetAllCosResponse,omitempty"`
}

type GetAllCosResponseContent struct {
	Cos []GenericResponse `json:"cos,omitempty"`
}

// ************* QUOTA ****************

type GetQuotaUsageRequest struct {
	Content GetQuotaUsageRequestContent `json:"GetQuotaUsageRequest,omitempty"`
}

type GetQuotaUsageRequestContent struct {
	Urn     string `json:"_jsns,attr"`
	Servers int    `json:"allServers,omitempty"`
	Domain  string `json:"domain,omitempty"`
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