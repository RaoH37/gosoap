package zsoap_test

import (
	"os"
)

var dlName = os.Getenv("GOSOAP_DL_NAME")

//func TestGetAllDistributionLists(t *testing.T) {
//	zcs, err := NewZcsClient()
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//
//	dls, err := zcs.GetAllDistributionLists("", "", 1, 1, "", 1, "")
//	for _, dl := range dls {
//		fmt.Printf("%v\n", dl)
//	}
//
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//}
//
//func TestDistributionListMember(t *testing.T) {
//	zcs, err := NewZcsClient()
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//
//	dl, err1 := zcs.GetDistributionList("", dlName, []string{})
//	fmt.Printf("%v\n", dl)
//
//	if err1 != nil {
//		t.Fatalf("%v", err)
//	}
//
//	newMembers := []string{"test1@toto.fr", "test2@toto.fr", "test3@toto.fr"}
//
//	err2 := zcs.AddDistributionListMemberRequest(dl.ID, newMembers)
//	if err2 != nil {
//		t.Fatalf("%v", err)
//	}
//}
