package zsoap

import "reflect"
import "strings"

type ZDomain struct {
	Client *Client
	ID string
	Name string
	ZimbraCreateTimestamp string
	ZimbraMailStatus string
}

func NewDomain(resp GenericResponse) *ZDomain {
	domain := &ZDomain{
		ID: resp.ID,
		Name: resp.Name,
	}
	
	for _, attr := range resp.Attrs {
		s := reflect.Indirect(reflect.ValueOf(&domain)).Elem()
		metric := s.FieldByName(strings.Title(attr.Key))
		if metric.IsValid() {
			switch metric.Interface().(type){
			case string:
				metric.SetString(attr.Value)
			case []string:
				metric.Set(reflect.Append(metric, reflect.ValueOf(attr.Value)))
			}
		}
	}

	return domain
}