package zsoap_test

import (
	"fmt"
	"testing"
)

func TestGetAllResources(t *testing.T) {
	zcs, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resources, err := zcs.GetAllResources("", "", 1, 1, "", 1, "")
	for _, resource := range resources {
		fmt.Printf("%v\n", resource)
	}

	if err != nil {
		t.Fatalf("%v", err)
	}
}
