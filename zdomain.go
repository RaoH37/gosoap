package zsoap

type ZDomain struct {
	Client                   *Client
	ID                       string
	Name                     string
	ZimbraCreateTimestamp    string
	ZimbraMailStatus         string
	ZimbraDomainDefaultCOSId string
}

func NewDomain(resp GenericResponse) *ZDomain {
	domain := &ZDomain{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &domain)

	return domain
}
