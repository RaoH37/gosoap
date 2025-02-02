package zsoap

func (s *ZAdmin) GetAllCoses(query string, applyConfig int, sortBy string, sortAscending int, attrs string) ([]ZCos, error) {
	_, _, _, coses, _, err := s.SearchDirectoryAll(query, "", 0, applyConfig, sortBy, "coses", sortAscending, attrs)
	return coses, err
}

type ZCos struct {
	Client                            *Client
	ID                                string
	Name                              string
	Description                       string
	ZimbraMailQuota                   int
	ZimbraFeatureMobileSyncEnabled    bool
	ZimbraFeatureMAPIConnectorEnabled bool
}

func NewCos(resp GenericResponse, client *Client) *ZCos {
	cos := &ZCos{
		Client: client,
		ID:     resp.ID,
		Name:   resp.Name,
	}

	setResponseAttrs(resp.Attrs, &cos)

	return cos
}
