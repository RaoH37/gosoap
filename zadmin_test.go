package zsoap_test

import (
	"fmt"
	"github.com/RaoH37/gosoap"
	"os"
	"testing"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var token = os.Getenv("GOSOAP_TOKEN")

func TestLogin(t *testing.T) {
	zcs := zsoap.ZAdmin{}
	zcs.Init(url, true)

	if len(token) > 0 {
		zcs.Client.SetToken(token)
	} else {
		err := zcs.Login(login, pwd)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}

	fmt.Println(zcs.Client.TOKEN)
}

func TestGetAllAccounts(t *testing.T) {
	zcs := zsoap.ZAdmin{}
	zcs.Init(url, true)

	if len(token) > 0 {
		zcs.Client.SetToken(token)
	} else {
		err := zcs.Login(login, pwd)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}

	accounts, err := zcs.GetAllAccounts("", "", 1, 1, "", 1, "")
	for _, account := range accounts {
		fmt.Printf("%v\n", account)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetAccounts(t *testing.T) {
	zcs := zsoap.ZAdmin{}
	zcs.Init(url, true)

	if len(token) > 0 {
		zcs.Client.SetToken(token)
	} else {
		err := zcs.Login(login, pwd)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}

	accounts, err := zcs.GetAccounts("", 10, 0, "", 1, 1, "", 1, "")
	for _, account := range accounts {
		fmt.Printf("%v\n", account)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetAllCos(t *testing.T) {
	zcs := zsoap.ZAdmin{}
	zcs.Init(url, true)

	if len(token) > 0 {
		zcs.Client.SetToken(token)
	} else {
		err := zcs.Login(login, pwd)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}

	coses, err := zcs.GetAllCos("", 1, "", 1, "")
	for _, cos := range coses {
		fmt.Printf("%v\n", cos)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetAllDomains(t *testing.T) {
	zcs := zsoap.ZAdmin{}
	zcs.Init(url, true)

	if len(token) > 0 {
		zcs.Client.SetToken(token)
	} else {
		err := zcs.Login(login, pwd)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}

	domains, err := zcs.GetAllDomains("", 1, 1, "", 1, "")
	for _, domain := range domains {
		fmt.Printf("%v\n", domain)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetAllDistributionLists(t *testing.T) {
	zcs := zsoap.ZAdmin{}
	zcs.Init(url, true)

	if len(token) > 0 {
		zcs.Client.SetToken(token)
	} else {
		err := zcs.Login(login, pwd)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}

	dls, err := zcs.GetAllDistributionLists("", "", 1, 1, "", 1, "")
	for _, dl := range dls {
		fmt.Printf("%v\n", dl)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}
