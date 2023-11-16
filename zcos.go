package zsoap

import (
	"reflect"
)

type ZCos struct {
	Client                            *Client
	ID                                string
	Name                              string
	Description                       string
	ZimbraMailQuota                   string
	ZimbraFeatureMobileSyncEnabled    string
	ZimbraFeatureMAPIConnectorEnabled string
}

// func NewCos(resp CosResponse) *ZCos {
func NewCos(resp GenericResponse) *ZCos {
	cos := &ZCos{
		ID:   resp.ID,
		Name: resp.Name,
	}

	for _, attr := range resp.Attrs {
		s := reflect.Indirect(reflect.ValueOf(&cos)).Elem()
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

	return cos
}
