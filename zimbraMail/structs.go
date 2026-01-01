package zimbraMail

import (
	"github.com/RaoH37/gosoap/zimbraCommon"
)

type GetFolderRequest struct {
	Content GetFolderRequestContent `json:"GetFolderRequest"`
}

type GetFolderRequestContent struct {
	Visible         zimbraCommon.ZBool `json:"visible,omitempty"`
	NeedGranteeName zimbraCommon.ZBool `json:"needGranteeName,omitempty"`
	View            string             `json:"view,omitempty"`
	Tr              zimbraCommon.ZBool `json:"tr,omitempty"`
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
	Grants []GrantNode `json:"grant"`
}

type GrantNode struct {
	Perm   string `json:"perm"`
	Gt     string `json:"gt"`
	ZID    string `json:"zid,omitempty"`
	Expiry int64  `json:"expiry,omitempty"`
	Gd     string `json:"d,omitempty"`
	Pw     string `json:"pw,omitempty"`
}

type LinkResponse struct {
	Owner         string             `json:"owner"`
	ZID           string             `json:"zid"`
	RID           int                `json:"rid"`
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
	Color    int                 `json:"color,omitempty"`
	RGB      string              `json:"rgb,omitempty"`
	Url      string              `json:"url,omitempty"`
	ParentID string              `json:"l"`
	Fie      *zimbraCommon.ZBool `json:"fie,omitempty"`
	Sync     *zimbraCommon.ZBool `json:"sync,omitempty"`
	Acls     *AclNode            `json:"acl,omitempty"`
	Urn      string              `json:"_jsns,attr"`
}

type FolderActionRequest struct {
	Content FolderActionRequestContent `json:"CreateFolderRequest"`
}

type FolderActionRequestContent struct {
	Action FolderActioNode `json:"action"`
}

type FolderActioNode struct {
	Recursive       *zimbraCommon.ZBool `json:"recursive,omitempty"`
	Url             string              `json:"url,omitempty"`
	ExcludeFreeBusy *zimbraCommon.ZBool `json:"excludeFreeBusy,omitempty"`
	ZID             string              `json:"zid,omitempty"`
	Gt              string              `json:"gt,omitempty"`
	View            string              `json:"view,omitempty"`
	ID              string              `json:"id,omitempty"`
	Operation       string              `json:"op"`
	ParentID        string              `json:"l,omitempty"`
	Flags           string              `json:"f,omitempty"`
	Color           int8                `json:"color,omitempty"`
	RGB             string              `json:"rgb,omitempty"`
	Name            string              `json:"name,omitempty"`
	Tn              string              `json:"tn,omitempty"`
}

type NoOpRequest struct {
	Content zimbraCommon.UrnRequestContent `json:"NoOpRequest"`
}
