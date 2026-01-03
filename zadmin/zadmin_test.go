package zadmin_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/RaoH37/gosoap/zadmin"
	"github.com/RaoH37/gosoap/zimbraAdmin"
	"github.com/RaoH37/gosoap/zimbraCommon"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var token = os.Getenv("GOSOAP_TOKEN")
var account_name = os.Getenv("GOSOAP_ACCOUNT_NAME")
var account_id = os.Getenv("GOSOAP_ACCOUNT_ID")
var res_name = os.Getenv("GOSOAP_RES_NAME")
var dl_name = os.Getenv("GOSOAP_DL_NAME")
var dl_id = os.Getenv("GOSOAP_DL_ID")
var domain_name = os.Getenv("GOSOAP_DOMAIN_NAME")
var domain_id = os.Getenv("GOSOAP_DOMAIN_ID")
var server_name = os.Getenv("GOSOAP_SERVER_NAME")
var server_id = os.Getenv("GOSOAP_SERVER_ID")

var letters = os.Getenv("GOSOAP_LETTERS")

const prefix = "unittest."

func RandStringRunes(n int) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}

	newStr := string(b)

	fmt.Println(newStr)

	return newStr
}

func findRandomOjectId(objType string) string {
	zcs, _ := NewZAdmin()

	var query string

	var types zimbraAdmin.SearchTypeList

	switch objType {
	case "accounts":
		query = "(zimbraMailDeliveryAddress=" + prefix + "*)"
		types = zimbraAdmin.SearchTypeList{
			zimbraAdmin.SearchTypeAccounts,
		}
	case "resources":
		query = "(zimbraMailDeliveryAddress=" + prefix + "*)"
		types = zimbraAdmin.SearchTypeList{
			zimbraAdmin.SearchTypeResources,
		}
	case "distributionlists":
		query = "(mail=" + prefix + "*)"
		types = zimbraAdmin.SearchTypeList{
			zimbraAdmin.SearchTypeDistributionLists,
		}
	case "domains":
		query = "(zimbraDomainName=" + prefix + "*)"
		types = zimbraAdmin.SearchTypeList{
			zimbraAdmin.SearchTypeDomains,
		}
	case "coses":
		query = "(cn=" + prefix + "*)"
		types = zimbraAdmin.SearchTypeList{
			zimbraAdmin.SearchTypeCoses,
		}
	}

	resp, err := zcs.SearchDirectoryRequest(query, 1_000_000, 1, 0, "", true, true, "mail", types, true, zimbraCommon.StringList{}, false)

	if err != nil || (len(resp.Content.Accounts) == 0 && len(resp.Content.CalResources) == 0 && len(resp.Content.Dls) == 0 && len(resp.Content.Domains) == 0 && len(resp.Content.Coses) == 0) {
		return ""
	}

	switch objType {
	case "accounts":
		return resp.Content.Accounts[0].ID
	case "resources":
		return resp.Content.CalResources[0].ID
	case "distributionlists":
		return resp.Content.Dls[0].ID
	case "domains":
		return resp.Content.Domains[0].ID
	case "coses":
		return resp.Content.Coses[0].ID
	default:
		return ""
	}
}

func NewZAdmin() (zadmin.ZAdmin, error) {
	zcs := zadmin.NewZAdmin(url, true, login, pwd, testing.Verbose(), time.Second*5, "zsoap", time.Second*30)

	if len(token) > 0 {
		zcs.SetToken(token)
	} else {
		resp, err := zcs.AuthRequest()

		if err != nil {
			return zcs, err
		}

		zcs.SetToken(resp.Content.TOKEN[0].Content)
	}

	return zcs, nil
}

func TestAddAccountAliasRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newAlias := prefix + RandStringRunes(6) + "_" + account_name

	err = zcs.AddAccountAliasRequest(account_id, newAlias)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestAddDistributionListAliasRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newAlias := prefix + RandStringRunes(6) + "_" + dl_name

	err = zcs.AddDistributionListAliasRequest(dl_id, newAlias)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestAddDistributionListMemberRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newMembers := []string{
		prefix + RandStringRunes(6) + "@" + domain_name,
		prefix + RandStringRunes(6) + "@" + domain_name,
		prefix + RandStringRunes(6) + "@" + domain_name,
		prefix + RandStringRunes(6) + "@" + domain_name,
		prefix + RandStringRunes(6) + "@" + domain_name,
	}

	err = zcs.AddDistributionListMemberRequest(dl_id, newMembers)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestCopyCosRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newCosName := prefix + RandStringRunes(6)

	resp, err := zcs.CopyCosRequest("", "default", newCosName)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestCreateAccountRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newAccountName := prefix + RandStringRunes(10) + "@" + domain_name

	attrs := zimbraCommon.AttrsNode{
		zimbraCommon.AttrNode{
			Name:  "sn",
			Value: "Désécot",
		},
		zimbraCommon.AttrNode{
			Name:  "givenName",
			Value: "Maxime",
		},
	}

	resp, err := zcs.CreateAccountRequest(newAccountName, RandStringRunes(10), attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestCreateCalendarResourceRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newAccountName := prefix + RandStringRunes(10) + "@" + domain_name

	attrs := zimbraCommon.BuildAttrsNode(map[string]string{
		"displayName":      "resource test",
		"zimbraCalResType": zadmin.Equipment,
	})

	resp, err := zcs.CreateCalendarResourceRequest(newAccountName, RandStringRunes(10), attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestCreateCosRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newCosName := prefix + RandStringRunes(10)

	attrs := zimbraCommon.AttrsNode{}

	resp, err := zcs.CreateCosRequest(newCosName, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestCreateDistributionListRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newAccountName := prefix + RandStringRunes(10) + "@" + domain_name

	attrs := zimbraCommon.AttrsNode{}

	resp, err := zcs.CreateDistributionListRequest(newAccountName, false, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestCreateDomainRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newDomainName := prefix + RandStringRunes(10) + ".com"

	attrs := zimbraCommon.AttrsNode{}

	resp, err := zcs.CreateDomainRequest(newDomainName, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestDelegateAuthRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp, err := zcs.DelegateAuthRequest("", account_name)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestDeleteAccountRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("accounts")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.DeleteAccountRequest(id)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestDeleteCalendarResourceRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("resources")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.DeleteCalendarResourceRequest(id)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestDeleteCosRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("coses")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.DeleteCosRequest(id)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestDeleteDistributionListRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("distributionlists")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.DeleteDistributionListRequest(id, false)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestDeleteDomainRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("domains")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.DeleteDomainRequest(id)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetAllServersRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp, err := zcs.GetAllServersRequest("")

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestGetAccountRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{}

	resp, err := zcs.GetAccountRequest("", account_name, attrs, true)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp.Content.Account[0].Name)
	}
}

func TestGetCalendarResourceRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{}

	resp, err := zcs.GetCalendarResourceRequest("", res_name, attrs, true)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp.Content.CalResource[0])
	}
}

func TestGetDistributionListRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{}

	resp, err := zcs.GetDistributionListRequest("", dl_name, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp.Content.Dl[0])
	}
}

func TestGetCosRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{
		"zimbraMailHostPool",
	}

	resp, err := zcs.GetCosRequest("", "default", attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp.Content.Cos[0])
	}
}

func TestGetDomainRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{}

	resp, err := zcs.GetDomainRequest("", domain_name, attrs, true)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp.Content.Domains[0])
	}
}

func TestGetQuotaUsageRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp, err := zcs.GetQuotaUsageRequest("", "", true)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		for _, account := range resp.Content.Account {
			fmt.Printf("%v\n", account.Name)
		}
	}
}

func TestGetServerRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{}

	resp, err := zcs.GetServerRequest("", server_name, true, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp.Content.Server[0])
	}
}

func TestModifyAccountRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("accounts")

	attrs := zimbraCommon.BuildAttrsNode(map[string]string{
		"displayName": "modified",
		"sn":          "modified",
		"givenName":   "modified",
		"description": "modified",
	})

	err = zcs.ModifyAccountRequest(id, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestModifyCalendarResourceRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("resources")

	attrs := zimbraCommon.BuildAttrsNode(map[string]string{
		"displayName": "modified",
	})

	err = zcs.ModifyCalendarResourceRequest(id, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestModifyCosRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("coses")

	attrs := zimbraCommon.BuildAttrsNode(map[string]string{
		"description": "modified",
	})

	err = zcs.ModifyCosRequest(id, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestModifyDistributionListRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	id := findRandomOjectId("distributionlists")

	attrs := zimbraCommon.BuildAttrsNode(map[string]string{
		"displayName": "modified",
	})

	err = zcs.ModifyDistributionListRequest(id, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestModifyDomainRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := zimbraCommon.BuildAttrsNode(map[string]string{
		"description": "modified",
	})

	err = zcs.ModifyDomainRequest(domain_id, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestModifyServerRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := zimbraCommon.BuildAttrsNode(map[string]string{
		"description": "modified",
	})

	err = zcs.ModifyServerRequest(server_id, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestNoOpRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = zcs.NoOpRequest()

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestRenameAccountRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newAccountName := prefix + RandStringRunes(10) + "@" + domain_name

	id := findRandomOjectId("accounts")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.RenameAccountRequest(id, newAccountName)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestRenameCosRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newCosName := prefix + RandStringRunes(10)

	id := findRandomOjectId("coses")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.RenameCosRequest(id, newCosName)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestSearchDirectoryRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	types := zimbraAdmin.SearchTypeList{
		zimbraAdmin.SearchTypeAccounts,
	}

	resp, err := zcs.SearchDirectoryRequest("", 1_000_000, 100, 0, "", true, true, "mail", types, true, zimbraCommon.StringList{}, false)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		for _, account := range resp.Content.Accounts {
			fmt.Printf("%v\n", account.Name)
		}
	}
}

func TestSearchDirectoryRequestCount(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	types := zimbraAdmin.SearchTypeList{
		zimbraAdmin.SearchTypeAccounts,
	}

	resp, err := zcs.SearchDirectoryRequest("", 1_000_000, 1, 0, "", true, true, "mail", types, true, zimbraCommon.StringList{}, true)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp.Content.Count)
	}
}

func TestSetPasswordRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newPassword := RandStringRunes(12)

	id := findRandomOjectId("accounts")

	if id == "" {
		t.Errorf("id is empty")
		return
	}

	err = zcs.SetPassword(id, newPassword)

	if err != nil {
		t.Fatalf("%v", err)
	}
}
