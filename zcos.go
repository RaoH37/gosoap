package zsoap

import "reflect"
import "strings"

type ZCos struct {
	Client *Client
	ID string
	Name string
	// ZimbraCreateTimestamp string
	// ZimbraServiceEnabled []string
}

// func NewCos(resp CosResponse) *ZCos {
func NewCos(resp GenericResponse) *ZCos {
	cos := &ZCos{
		ID: resp.ID,
		Name: resp.Name,
	}
	
	for _, attr := range resp.Attrs {
		s := reflect.Indirect(reflect.ValueOf(&cos)).Elem()
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

	return cos
}