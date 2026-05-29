package zimbraMail

import (
	"github.com/RaoH37/gosoap/zimbraCommon"
)

const Urn = "urn:zimbraMail"

func NewGetFolderRequest(view FolderView, visible bool, needGranteeName bool, tr bool) (*GetFolderRequest, GetFolderResponse) {
	content := GetFolderRequestContent{
		Visible:         zimbraCommon.ZBool(visible),
		NeedGranteeName: zimbraCommon.ZBool(needGranteeName),
		Tr:              zimbraCommon.ZBool(tr),
		Urn:             Urn,
	}

	if view != "" {
		content.View = view
	}

	return &GetFolderRequest{
		Content: content,
	}, GetFolderResponse{}
}

func NewCreateFolderRequest(
	name string,
	view FolderView,
	parentId string) (*CreateFolderRequest, CreateFolderResponse) {

	folder := Folder{
		Name:     name,
		View:     view,
		ParentID: parentId,
	}

	return &CreateFolderRequest{
		Content: CreateFolderRequestContent{
			Folder: folder,
			Urn:    Urn,
		},
	}, CreateFolderResponse{}
}

func NewNoOpRequest() (*NoOpRequest, GetFolderResponse) {
	return &NoOpRequest{
		Content: zimbraCommon.UrnRequestContent{
			Urn: Urn,
		},
	}, GetFolderResponse{}
}

func NewFolderActionRequest(
	op ItemOperation,
	name string,
	view FolderView,
	id string,
	parentId string,
	recursive bool,
	url string,
	excludeFreeBusy bool,
	zid string,
	gt string,
	flags zimbraCommon.StringList,
	color int8,
	rgb string,
	tags zimbraCommon.StringList) (*ItemActionRequest, FolderActionResponse) {
	zrecursive := zimbraCommon.ZBool(recursive)
	zexcludeFreeBusy := zimbraCommon.ZBool(excludeFreeBusy)

	return &ItemActionRequest{
		Content: ItemActionRequestContent{
			Action: ActioNode{
				Recursive:       &zrecursive,
				Url:             url,
				ExcludeFreeBusy: &zexcludeFreeBusy,
				ZID:             zid,
				GranteeType:     gt,
				View:            view,
				ID:              id,
				Operation:       op,
				ParentID:        parentId,
				Flags:           flags,
				Color:           color,
				RGB:             rgb,
				Name:            name,
				Tags:            tags,
			},
		},
	}, FolderActionResponse{}
}
