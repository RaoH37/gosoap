package zsoap

type ZCos struct {
	Client                            *Client
	ID                                string
	Name                              string
	Description                       string
	ZimbraMailQuota                   int
	ZimbraFeatureMobileSyncEnabled    bool
	ZimbraFeatureMAPIConnectorEnabled bool
}

func NewCos(resp GenericResponse) *ZCos {
	cos := &ZCos{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &cos)

	return cos
}
