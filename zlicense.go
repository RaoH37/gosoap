package zsoap

import (
	"reflect"
	"strconv"
	"strings"
)

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

	SetLicenseAttrsByReflect(license, resp.License)
	SetLicenseAttrsByReflect(license, resp.Activation)
	SetLicenseAttrsByReflect(license, resp.Info)

	return license
}

func SetLicenseAttrsByReflect(license *ZLicense, attrNames []AttrNamesResponse) {
	for _, attrName := range attrNames {
		// fmt.Println("%v", attrName)
		// fmt.Println("===================")
		for _, attr := range attrName.Attrs {
			// fmt.Println("%v", attr)
			s := reflect.Indirect(reflect.ValueOf(&license)).Elem()
			metric := s.FieldByName(strings.Title(attr.Key))
			if metric.IsValid() {
				switch metric.Interface().(type) {
				case string:
					metric.SetString(attr.Value)
				case []string:
					metric.Set(reflect.Append(metric, reflect.ValueOf(attr.Value)))
				case int:
					int_vale, _ := strconv.ParseInt(attr.Value, 10, 32)
					metric.SetInt(int_vale)
				}
			}
		}
	}
}
