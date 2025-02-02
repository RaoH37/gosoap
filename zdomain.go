package zsoap

import "log"

func (s *ZAdmin) GetAllDomains(query string, applyCos int, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZDomain, error) {
	_, _, domains, _, _, err := s.SearchDirectoryAll(query, "", applyCos, applyConfig, sortBy, "domains", sortAscending, attrs)
	return domains, err
}

func (s *ZAdmin) GetDomain(by ByRequest, attrs []string) (*ZDomain, error) {
	req, soapAction := NewGetDomainRequest(by, attrs)
	resp := GetDomainResponse{}

	if err := s.Client.Call(soapAction, req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	domain := NewDomain(resp.Content.Domain[0], s.Client)
	domain.Client = s.Client

	return domain, nil
}

func (s *ZAdmin) GetDomainByName(name string, attrs []string) (*ZDomain, error) {
	by := NewByRequest(NAME_STR, name)
	return s.GetDomain(by, attrs)
}

func (s *ZAdmin) GetDomainById(id string, attrs []string) (*ZDomain, error) {
	by := NewByRequest(ID_STR, id)
	return s.GetDomain(by, attrs)
}

type ZDomain struct {
	Client                                                     *Client
	ID                                                         string
	Name                                                       string
	ZimbraAdminConsoleCatchAllAddressEnabled                   bool
	ZimbraAdminConsoleDNSCheckEnabled                          bool
	ZimbraAdminConsoleLDAPAuthEnabled                          bool
	ZimbraAdminConsoleSkinEnabled                              bool
	ZimbraAggregateQuotaLastUsage                              int
	ZimbraAutoProvBatchSize                                    int
	ZimbraAutoProvNotificationBody                             string
	ZimbraAutoProvNotificationSubject                          string
	ZimbraBasicAuthRealm                                       string
	ZimbraChatConversationAuditEnabled                         bool
	ZimbraCommunityHomeURL                                     string
	ZimbraCommunityUsernameMapping                             string
	ZimbraCreateTimestamp                                      string
	ZimbraDomainAggregateQuota                                 int
	ZimbraDomainAggregateQuotaPolicy                           string
	ZimbraDomainAggregateQuotaWarnPercent                      int
	ZimbraDomainDefaultCOSId                                   string
	ZimbraDomainMandatoryMailSignatureEnabled                  bool
	ZimbraDomainName                                           string
	ZimbraDomainStatus                                         string
	ZimbraDomainType                                           string
	ZimbraExportMaxDays                                        int
	ZimbraExternalShareInvitationUrlExpiration                 int
	ZimbraFileUploadMaxSizePerFile                             int
	ZimbraFreebusyExchangeCachedInterval                       string
	ZimbraFreebusyExchangeCachedIntervalStart                  string
	ZimbraFreebusyExchangeServerType                           string
	ZimbraGalAccountId                                         string
	ZimbraGalAlwaysIncludeLocalCalendarResources               bool
	ZimbraGalAutoCompleteLdapFilter                            string
	ZimbraGalGroupIndicatorEnabled                             bool
	ZimbraGalInternalSearchBase                                string
	ZimbraGalLdapAttrMap                                       []string
	ZimbraGalLdapPageSize                                      int
	ZimbraGalLdapValueMap                                      []string
	ZimbraGalMaxResults                                        int
	ZimbraGalSyncLdapPageSize                                  int
	ZimbraGalSyncMaxConcurrentClients                          int
	ZimbraGalSyncSizeLimit                                     int
	ZimbraGalSyncTimestampFormat                               string
	ZimbraGalTokenizeAutoCompleteKey                           string
	ZimbraGalTokenizeSearchKey                                 string
	ZimbraId                                                   string
	ZimbraInternalSharingCrossDomainEnabled                    bool
	ZimbraLdapGalSyncDisabled                                  bool
	ZimbraMailDomainQuota                                      int
	ZimbraMailSSLClientCertPrincipalMap                        string
	ZimbraMailStatus                                           string
	ZimbraMobileMetadataMaxSizeEnabled                         bool
	ZimbraReverseProxyClientCertMode                           string
	ZimbraReverseProxyExternalRouteIncludeOriginalAuthusername bool
	ZimbraSkinLogoURL                                          string
	ZimbraWebClientMaxInputBufferLength                        int
	ZimbraWebClientStaySignedInDisabled                        bool
	ZimbraWebClientSupportedHelps                              []string
	ZimbraZimletDataSensitiveInMixedModeDisabled               bool
}

func NewDomain(resp GenericResponse, client *Client) *ZDomain {
	domain := &ZDomain{
		Client: client,
		ID:     resp.ID,
		Name:   resp.Name,
	}

	setResponseAttrs(resp.Attrs, &domain)

	return domain
}
