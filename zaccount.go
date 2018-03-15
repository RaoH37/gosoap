package zsoap

import "log"
import "reflect"
import "strings"

type ZAccount struct {
	Client *Client
	ID string
	Name string
	Used int
	Limit int
	ZimbraMailHost string
	ZimbraMailTransport string
}

func (a *ZAccount) Modify(attrs map[string]string) {
	
	req, soapAction := NewModifyAccountRequest(a.ID, attrs)
	resp := ModifyAccountResponse{}

	if err := a.Client.Call(soapAction, req, &resp); err != nil {
	 	log.Fatal(err)
	}
}

func NewAccount(resp GenericResponse) *ZAccount {
	account := &ZAccount{
		ID: resp.ID,
		Name: resp.Name,
	}
	for _, attr := range resp.Attrs {
		s := reflect.Indirect(reflect.ValueOf(&account)).Elem()
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
	return account
}

func NewAccountQuota(resp QuotaResponse) *ZAccount {
	account := &ZAccount{
		ID:    resp.ID,
		Name:  resp.Name,
		Used:  resp.Used,
		Limit: resp.Limit,
	}
	return account
}

func NewModifyAccountRequest(id string, attrs map[string]string) (*ModifyAccountRequest, string){
	
	a := make([]AttrResponse, 0)

	for key, value := range attrs {
		a = append(a, AttrResponse{
			Key: key,
			Value: value,
		})
	}

	r := &ModifyAccountRequest{
		Content: ModifyAccountRequestContent{
			Urn: urnAdmin,
			ID: id,
			Attrs: a,
		},
	}
	return r, "urn:zimbraAdmin/ModifyAccount"
}
