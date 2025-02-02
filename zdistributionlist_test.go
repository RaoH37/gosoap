package zsoap_test

import (
	"fmt"
	"testing"
)

func TestGetAllDistributionLists(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	dls, err := zcs.GetAllDistributionLists("", "", 1, 1, "", 1, "")
	for _, dl := range dls {
		fmt.Printf("%v\n", dl)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}
