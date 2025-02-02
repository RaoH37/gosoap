package zsoap_test

import (
	"fmt"
	"testing"
)

func TestGetAllCos(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	coses, err := zcs.GetAllCoses("", 1, "", 1, "")
	for _, cos := range coses {
		fmt.Printf("%v\n", cos)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}
