package zsoap

import "reflect"
import "strings"

type ZServer struct {
	Client *Client
	ID string
	Name string
	ZimbraCreateTimestamp string
	ZimbraServiceEnabled []string
}

// func NewServer(resp ServerResponse) *ZServer {
func NewServer(resp GenericResponse) *ZServer {
	server := &ZServer{
		ID: resp.ID,
		Name: resp.Name,
	}
	
	for _, attr := range resp.Attrs {
		s := reflect.Indirect(reflect.ValueOf(&server)).Elem()
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

	return server
}