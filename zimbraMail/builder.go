package zimbraMail

const urnMail = "urn:zimbraMail"

func NewGetFolderRequest(view string) (*GetFolderRequest, GetFolderResponse) {
	content := GetFolderRequestContent{
		Visible:         true,
		NeedGranteeName: true,
		Tr:              true,
		Urn:             urnMail,
	}

	if view != "" {
		content.View = view
	}

	return &GetFolderRequest{
		Content: content,
	}, GetFolderResponse{}
}
