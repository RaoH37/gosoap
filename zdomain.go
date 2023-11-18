package zsoap

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
	ZilbraDomainDefaultCOSId                                   string
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

func NewDomain(resp GenericResponse) *ZDomain {
	domain := &ZDomain{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &domain)

	return domain
}
