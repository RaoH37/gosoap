package zimbraMail

const urnMail = "urn:zimbraMail"

func NewGetFolderRequest() (*GetFolderRequest, GetFolderResponse) {
	return &GetFolderRequest{
		Content: GetFolderRequestContent{
			Visible:         true,
			NeedGranteeName: true,
			View:            "message",
			Tr:              true,
			Urn:             urnMail,
		},
	}, GetFolderResponse{}
}
