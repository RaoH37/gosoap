package zsoap_test

import (
	"fmt"
	"testing"
)

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
