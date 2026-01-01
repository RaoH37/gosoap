package zimbraMail

import "github.com/RaoH37/gosoap/zimbraCommon"

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
}
