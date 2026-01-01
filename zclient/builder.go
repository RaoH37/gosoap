package zclient

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/RaoH37/gosoap/zimbraAdmin"
)

var regexpStartByUpperChar = regexp.MustCompile(`^[[:upper:]]`)

func capitalizeByteSlice(str string) string {
	if regexpStartByUpperChar.MatchString(str) {
		return str
	}

	bs := []byte(str)
	if len(bs) == 0 {
		return ""
	}
	bs[0] = byte(bs[0] - 32)

	return string(bs)
}

func setResponseAttrs(attrs []zimbraAdmin.AttrResponse, object interface{}) {
	for _, attr := range attrs {
		s := reflect.Indirect(reflect.ValueOf(object)).Elem()
		metric := s.FieldByName(capitalizeByteSlice(attr.Key))

		// fmt.Printf("key=%s upkey=%s :: value=%s (%T) valid=%s\n", attr.Key, capitalizeByteSlice(attr.Key), attr.Value, attr.Value, metric.IsValid())

		if metric.IsValid() {
			switch metric.Interface().(type) {
			case bool:
				metric.SetBool(strings.ToLower(attr.Value) == "true")
			case int:
				vint, _ := strconv.ParseInt(attr.Value, 10, 64)
				metric.SetInt(vint)
			case string:
				metric.SetString(attr.Value)
			case []string:
				elements := strings.Split(attr.Value, ",")

				for _, element := range elements {
					// elements[i] = strings.TrimSpace(element)
					metric.Set(reflect.Append(metric, reflect.ValueOf(strings.TrimSpace(element))))
				}

				// metric.Set(reflect.Append(metric, reflect.ValueOf(attr.Value)))
			}
		}
	}
}

func buildAccount(resp zimbraAdmin.GenericResponse) *ZAccount {
	account := &ZAccount{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &account)

	return account
}

func buildAccountQuota(resp zimbraAdmin.QuotaResponse) *ZAccount {
	account := &ZAccount{
		ID:    resp.ID,
		Name:  resp.Name,
		Used:  resp.Used,
		Limit: resp.Limit,
	}

	return account
}

func buildResource(resp zimbraAdmin.GenericResponse) *ZCalendarResource {
	resource := &ZCalendarResource{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &resource)

	return resource
}

func buildCos(resp zimbraAdmin.GenericResponse) *ZCos {
	cos := &ZCos{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &cos)

	return cos
}

func buildDistributionList(resp zimbraAdmin.GenericResponse) *ZDistributionList {
	dl := &ZDistributionList{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &dl)

	return dl
}

func buildDomain(resp zimbraAdmin.GenericResponse) *ZDomain {
	domain := &ZDomain{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &domain)

	return domain
}

func buildLicense(resp zimbraAdmin.GetLicenseResponseContent) *ZLicense {
	license := &ZLicense{}

	for _, attrName := range resp.License {
		setResponseAttrs(attrName.ToAttrsResponse(), &license)
	}

	for _, attrName := range resp.Activation {
		setResponseAttrs(attrName.ToAttrsResponse(), &license)
	}

	for _, attrName := range resp.Info {
		setResponseAttrs(attrName.ToAttrsResponse(), &license)
	}

	return license
}

func buildServer(resp zimbraAdmin.GenericResponse) *ZServer {
	server := &ZServer{
		ID:   resp.ID,
		Name: resp.Name,
	}

	setResponseAttrs(resp.Attrs, &server)

	return server
}
