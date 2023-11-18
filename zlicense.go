package zsoap

type ZLicense struct {
	AccountsLimit                        int
	ArchivingAccountsLimit               int
	AttachmentConversionEnabled          bool
	AttachmentIndexingAccountsLimit      int
	BackupEnabled                        bool
	CrossMailboxSearchEnabled            bool
	EwsAccountsLimit                     int
	HierarchicalStorageManagementEnabled bool
	ISyncAccountsLimit                   int
	InstallType                          string
	IssuedOn                             string
	IssuedToEmail                        string
	IssuedToName                         string
	LicenseId                            string
	MAPIConnectorAccountsLimit           int
	MobileSyncAccountsLimit              int
	MobileSyncEnabled                    bool
	ResellerName                         string
	SMIMEAccountsLimit                   int
	TouchClientsAccountsLimit            int
	TwoFactorAuthAccountsLimit           int
	ValidFrom                            string
	ValidUntil                           string
	VoiceAccountsLimit                   int
	ZSSAccountsLimit                     int
	ZTalkAccountsLimit                   int
	ZXAccountsLimit                      int
	ZXDesktopAccountsLimit               int
	ZXWebAccountsLimit                   int
	ActivationId                         string
	Fingerprint                          string
	LastUpdate                           string
	Version                              string
	Status                               string
	TotalAccounts                        int
	ArchivingAccounts                    int
	ServerTime                           int
}

func NewLicense(resp GetLicenseResponseContent) *ZLicense {
	license := &ZLicense{}

	for _, attrName := range resp.License {
		setResponseAttrs(attrName.ToAttrsResponse(), &license)
	}

	for _, attrName := range resp.Activation {
		setResponseAttrs(attrName.ToAttrsResponse(), &license)
	}

	for _, attrName := range resp.Info {
		setResponseAttrs(attrName.ToAttrsResponse(), &license)
	}

	return license
}
