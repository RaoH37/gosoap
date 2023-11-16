package zsoap

import (
	"reflect"
)

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

	for _, attr := range resp.Attrs {
		s := reflect.Indirect(reflect.ValueOf(&domain)).Elem()
		metric := s.FieldByName(attr.Key)
		if metric.IsValid() {
			switch metric.Interface().(type) {
			case string:
				metric.SetString(attr.Value)
			case []string:
				metric.Set(reflect.Append(metric, reflect.ValueOf(attr.Value)))
			}
		}
	}

	return domain
}
