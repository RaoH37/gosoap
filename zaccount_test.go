package zsoap_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var accountName = os.Getenv("GOSOAP_ACCOUNT_NAME")

func TestGetAllAccounts(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
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
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	accounts, err := zcs.GetAccounts("", 10, 0, "", 1, 1, "", 1, "")
	for _, account := range accounts {
		fmt.Printf("%v\n", account)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetAccount(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	account, err := zcs.GetAccountByName(accountName, []string{})
	fmt.Printf("%v\n", account)

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestAliasAccountAlias(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	account, err1 := zcs.GetAccountByName(accountName, []string{})
	fmt.Printf("%v\n", account)

	if err1 != nil {
		t.Fatalf("%v", err)
	}

	newAlias := fmt.Sprintf("%d.%s", rand.Intn(100), account.Name)
	fmt.Printf("Add newAlias=%s\n", newAlias)

	err2 := zcs.AddAccountAlias(account.ID, newAlias)
	if err2 != nil {
		t.Fatalf("%v", err)
	}

	err3 := zcs.RemoveAccountAlias(account.ID, newAlias)
	if err3 != nil {
		t.Fatalf("%v", err)
	}

	fmt.Printf("Remove newAlias=%s\n", newAlias)
}
