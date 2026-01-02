package zimbraMail

import (
	"github.com/RaoH37/gosoap/zimbraCommon"
)

type GetFolderRequest struct {
	Content GetFolderRequestContent `json:"GetFolderRequest"`
}

type GetFolderRequestContent struct {
	Visible         zimbraCommon.ZBool `json:"visible,omitzero"`
	NeedGranteeName zimbraCommon.ZBool `json:"needGranteeName,omitzero"`
	View            string             `json:"view,omitempty"`
	Tr              zimbraCommon.ZBool `json:"tr,omitzero"`
	Urn             string             `json:"_jsns,attr"`
}

type GetFolderResponse struct {
	Content GetFolderResponseContent `json:"GetFolderResponse"`
}

type GetFolderResponseContent struct {
	Folders []FolderResponse `json:"folder"`
}

type FolderResponse struct {
	ID            string           `json:"id"`
	UUID          string           `json:"uuid"`
	Name          string           `json:"name"`
	AbsFolderPath string           `json:"absFolderPath"`
	ParentID      string           `json:"l"`
	Flags         string           `json:"f"`
	Color         int              `json:"color"`
	RGB           string           `json:"rgb"`
	UnRead        int              `json:"u,omitempty"`
	View          string           `json:"view,omitempty"`
	Size          int              `json:"s,omitempty"`
	Folders       []FolderResponse `json:"folder"`
	Acls          AclNode          `json:"acl"`
	Links         []LinkResponse   `json:"link"`
}

type AclNode struct {
	Grants []zimbraCommon.GrantNode `json:"grant"`
}

type LinkResponse struct {
	Owner         string             `json:"owner"`
	ZID           string             `json:"zid"`
	RemoteID      int                `json:"rid"`
	Broken        zimbraCommon.ZBool `json:"broken"`
	ID            string             `json:"id"`
	UUID          string             `json:"uuid"`
	Name          string             `json:"name"`
	AbsFolderPath string             `json:"absFolderPath"`
	ParentID      string             `json:"l"`
	Flags         string             `json:"f"`
	Color         int                `json:"color"`
	RGB           string             `json:"rgb"`
	UnRead        int                `json:"u,omitempty"`
	View          string             `json:"view,omitempty"`
	Size          int                `json:"s,omitempty"`
}

type CreateFolderRequest struct {
	Content CreateFolderRequestContent `json:"CreateFolderRequest"`
}

type CreateFolderRequestContent struct {
	Name     string              `json:"name"`
	View     string              `json:"view,omitempty"`
	Flags    string              `json:"f,omitempty"`
	Color    int                 `json:"color,omitzero"`
	RGB      string              `json:"rgb,omitempty"`
	Url      string              `json:"url,omitempty"`
	ParentID string              `json:"l"`
	Fie      *zimbraCommon.ZBool `json:"fie,omitempty"`
	Sync     *zimbraCommon.ZBool `json:"sync,omitempty"`
	Acls     *AclNode            `json:"acl,omitempty"`
	Urn      string              `json:"_jsns,attr"`
}

type ItemActionRequest struct {
	Content ItemActionRequestContent `json:"ItemActionRequest"`
}

type ItemActionRequestContent struct {
	Action ActioNode `json:"action"`
}

type ActioNode struct {
	Recursive       *zimbraCommon.ZBool `json:"recursive,omitempty"`
	Url             string              `json:"url,omitempty"`
	ExcludeFreeBusy *zimbraCommon.ZBool `json:"excludeFreeBusy,omitempty"`
	ZID             string              `json:"zid,omitempty"`
	GranteeType     string              `json:"gt,omitempty"`
	View            string              `json:"view,omitempty"`
	ID              string              `json:"id,omitempty"`
	Operation       string              `json:"op"`
	ParentID        string              `json:"l,omitempty"`
	Flags           string              `json:"f,omitempty"`
	Color           int8                `json:"color,omitzero"`
	RGB             string              `json:"rgb,omitempty"`
	Name            string              `json:"name,omitempty"`
	Tags            string              `json:"tn,omitempty"`
}

type FolderActionResponse struct {
	Content struct{} `json:"FolderActioResponse"`
}

type NoOpRequest struct {
	Content zimbraCommon.UrnRequestContent `json:"NoOpRequest"`
}

type NoOpResponse struct {
	Content struct{} `json:"NoOpResponse"`
}

type CreateMountpointRequest struct {
	Content CreateMountpointRequestContent `json:"CreateMountpointRequest"`
}

type CreateMountpointRequestContent struct {
	Urn  string         `json:"_jsns,attr"`
	Link MountpointNode `json:"link"`
}

type MountpointNode struct {
	Name     string `json:"n,attr"`
	View     string `json:"view,attr,omitzero"`
	Owner    string `json:"owner,attr"`
	RemoteID string `json:"rid,attr,omitzero"`
	Color    int8   `json:"color,attr,omitzero"`
}

type CreateMountpointResponse struct {
	Content struct {
		Link struct {
			ID string `json:"id,attr"`
		} `json:"link"`
	} `json:"CreateMountpointResponse"`
}

type CreateTagRequest struct {
	Content CreateTagRequestContent `json:"CreateTagRequest"`
}

type CreateTagRequestContent struct {
	Urn string  `json:"_jsns,attr"`
	Tag TagNode `json:"tag"`
}

type TagNode struct {
	Name  string `json:"name,attr"`
	Color int8   `json:"color,attr,omitzero"` // 0=none, 1=blue, 2=cyan, etc.
}

type CreateTagResponse struct {
	Content struct {
		Tag struct {
			ID    string `json:"id,attr"`
			Name  string `json:"n,attr"`
			Color int    `json:"color,attr"`
		} `json:"tag"`
	} `json:"CreateTagResponse"`
}

type AddMsgRequest struct {
	Content AddMsgRequestContent `json:"AddMsgRequest"`
}

type AddMsgRequestContent struct {
	Urn     string      `json:"_jsns,attr"`
	Message MessageNode `json:"m"`
}

type MessageNode struct {
	ParentID string                     `json:"l,attr"`
	Flags    string                     `json:"f,attr,omitzero"`
	Tags     string                     `json:"tn,attr,omitzero"`
	Date     int64                      `json:"d,attr,omitzero"`
	Content  zimbraCommon.ContentString `json:"content"`
}

type AddMsgResponse struct {
	Content struct {
		Message struct {
			ID string `json:"id,attr"`
		} `json:"m"`
	} `json:"AddMsgResponse"`
}
