package zsoap_test

import (
	"fmt"
	"os"
	"testing"
)

var domainName = os.Getenv("GOSOAP_DOMAIN_NAME")

func TestGetAllDomains(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	domains, err := zcs.GetAllDomains("", 1, 1, "", 1, "")
	for _, domain := range domains {
		fmt.Printf("%v\n", domain)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetDomain(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	domain, err := zcs.GetDomainByName(domainName, []string{})
	fmt.Printf("%v\n", domain)

	if err != nil {
		t.Fatalf("%v", err)
	}
}
