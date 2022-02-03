package zsoap

import "strings"

func NewAuthRequest(name string, password string) (*AuthRequest, string) {
	r := &AuthRequest{Content: AuthRequestContent{
		Name:     name,
		Password: password,
		Urn:      urnAdmin,
	},
	}
	return r, "urn:zimbraAdmin/Auth"
}

func NewByRequest(by string, value string) ByRequest {
	return ByRequest{
		By:    by,
		Value: value,
	}
}

func NewGetAccountRequest(by ByRequest, attrs []string) (*GetAccountRequest, string) {
	r := &GetAccountRequest{
		Content: GetAccountRequestContent{
			Urn:     urnAdmin,
			Account: by,
		},
	}

	if attrs != nil {
		r.Content.Attrs = strings.Join(attrs, ",")
	}

	return r, "urn:zimbraAdmin/GetAccount"
}

func NewGetAllAccountsRequest(server *ByRequest, domain *ByRequest) (*GetAllAccountsRequest, string) {
	r := &GetAllAccountsRequest{
		Content: GetAllAccountsRequestContent{
			Urn: urnAdmin,
		},
	}
	if server != nil {
		r.Content.Server = server
	}
	if domain != nil {
		r.Content.Domain = domain
	}
	return r, "urn:zimbraAdmin/GetAllAccounts"
}

func NewGetAllCosRequest() (*GetAllCosRequest, string) {
	r := &GetAllCosRequest{
		Content: GetAllCosRequestContent{
			Urn: urnAdmin,
		},
	}
	return r, "urn:zimbraAdmin/GetAllCos"
}

func NewGetAllDomainsRequest() (*GetAllDomainsRequest, string) {
	r := &GetAllDomainsRequest{
		Content: GetAllDomainsRequestContent{
			Urn: urnAdmin,
		},
	}
	return r, "urn:zimbraAdmin/GetAllDomains"
}

func NewGetAllServersRequest(service string) (*GetAllServersRequest, string) {
	r := &GetAllServersRequest{
		Content: GetAllServersRequestContent{
			Urn: urnAdmin,
		},
	}
	if service != "" {
		r.Content.Service = service
	}
	return r, "urn:zimbraAdmin/GetAllServers"
}

func NewGetQuotaUsageRequest(domain string, allServers int) (*GetQuotaUsageRequest, string) {
	r := &GetQuotaUsageRequest{
		Content: GetQuotaUsageRequestContent{
			Urn:     urnAdmin,
			Servers: allServers,
			Domain:  domain,
		},
	}
	return r, "urn:zimbraAdmin/GetQuotaUsageRequest"
}

func NewBackupQueryRequest() (*BackupQueryRequest, string) {
	r := &BackupQueryRequest{
		Content: BackupQueryRequestContent{
			Urn:   urnAdmin,
			Query: make(map[string]string),
		},
	}
	return r, "urn:zimbraAdmin/BackupQueryRequest"
}

func NewSearchDirectoryRequest(params *SearchDirectoryParams) (*SearchDirectoryRequest, string) {
	r := &SearchDirectoryRequest{
		Content: *params,
	}
	return r, "urn:zimbraAdmin/SearchDirectoryRequest"
}
