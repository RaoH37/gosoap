package zsoap_test

import (
	"testing"
)

func TestLogin(t *testing.T) {
	_, err := NewZcsClient()
	if err != nil {
		t.Fatalf("%v", err)
	}
}
