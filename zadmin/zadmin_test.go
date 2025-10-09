package zadmin_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/RaoH37/gosoap/zadmin"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var token = os.Getenv("GOSOAP_TOKEN")
var account_name = os.Getenv("GOSOAP_ACCOUNT_NAME")
var dl_name = os.Getenv("GOSOAP_DL_NAME")
var domain_name = os.Getenv("GOSOAP_DOMAIN_NAME")

func NewZAdmin() (zadmin.ZAdmin, error) {
	zcs := zadmin.NewZAdmin(url, true, login, pwd, false, time.Second*5, "zsoap", time.Second*30)

	if len(token) > 0 {
		zcs.AuthToken = token
	} else {
		resp, err := zcs.AuthRequest()

		if err != nil {
			return zcs, err
		}

		zcs.AuthToken = resp.Content.TOKEN[0].Content
	}

	return zcs, nil
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

	for _, account := range resp.Content.Account {
		fmt.Printf("%v\n", account.Name)
	}
}

func TestSearchDirectoryRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp, err := zcs.SearchDirectoryRequest("", 1_000_000, 100, 0, "", 1, 1, "mail", "accounts", 1, "", 0)

	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, account := range resp.Content.Accounts {
		fmt.Printf("%v\n", account.Name)
	}
}

func TestSearchDirectoryRequestCount(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp, err := zcs.SearchDirectoryRequest("", 1_000_000, 1, 0, "", 1, 1, "mail", "accounts", 1, "", 1)

	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Printf("%v\n", resp.Content.Count)
}

func TestGetAccountRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{}

	resp, err := zcs.GetAccountRequest("", account_name, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Printf("%v\n", resp.Content.Account[0].Name)
}

func TestGelAllServersRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp, err := zcs.GetAllServersRequest("")

	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, server := range resp.Content.Servers {
		fmt.Printf("%v\n", server.Name)
	}
}

func TestGelDomainRequest(t *testing.T) {
	zcs, err := NewZAdmin()
	if err != nil {
		t.Fatalf("%v", err)
	}

	attrs := []string{}

	resp, err := zcs.GetDomainRequest("", domain_name, attrs)

	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, domain := range resp.Content.Domains {
		fmt.Printf("%v\n", domain.Name)
	}
}
