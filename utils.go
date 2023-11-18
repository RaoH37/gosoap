package zsoap

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
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

func setResponseAttrs(attrs []AttrResponse, object interface{}) {
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
				metric.Set(reflect.Append(metric, reflect.ValueOf(attr.Value)))
			}
		}
	}
}
