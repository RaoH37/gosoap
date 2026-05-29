package zimbraCommon

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const ID = "id"
const NAME = "name"

type ByNode struct {
	By    string `json:"by"`
	Value string `json:"_content"`
}

func NewByNode(by string, value string) ByNode {
	return ByNode{
		By:    by,
		Value: value,
	}
}

func NewByIdOrNameNode(id string, name string) ByNode {
	if id != "" {
		return NewByNode(ID, id)
	}

	return NewByNode(NAME, name)
}

type ContentString struct {
	Content string `json:"_content"`
}

type IdNameNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AttrNode struct {
	Name  string `json:"n"`
	Value string `json:"_content"`
}

type AttrsNode []AttrNode

func BuildAttrsNode(attrs map[string]string) AttrsNode {
	a := make(AttrsNode, len(attrs))

	i := 0
	for name, value := range attrs {
		a[i] = AttrNode{
			Name:  name,
			Value: value,
		}
		i++
	}

	return a
}

type NameNode struct {
	Name  string `json:"name"`
	Value string `json:"_content,omitempty"`
}

type AttrNamesNode struct {
	Attributs []NameNode `json:"omitempty"`
}

func (anr *AttrNamesNode) ToAttrsNode() []AttrNode {
	collection := make([]AttrNode, len(anr.Attributs))

	for i, attr := range anr.Attributs {
		collection[i] = AttrNode(attr)
	}

	return collection
}

type GrantNode struct {
	Perm        string `json:"perm"`
	GranteeType string `json:"gt"`
	ZID         string `json:"zid,omitempty"`
	Expiry      int64  `json:"expiry,omitempty"`
	GranteeName string `json:"d,omitempty"`
	Password    string `json:"pw,omitempty"`
}

type UrnRequestContent struct {
	Urn string `json:"_jsns"`
}

type ZBool bool

func (zb ZBool) MarshalJSON() ([]byte, error) {
	if zb {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

func (zb *ZBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "1" || s == `"1"` {
		*zb = true
	} else {
		*zb = false
	}
	return nil
}

type FlexString string

func (fs *FlexString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*fs = FlexString(s)
		return nil
	}

	var n float64
	if err := json.Unmarshal(data, &n); err == nil {
		*fs = FlexString(strconv.FormatFloat(n, 'f', -1, 64))
		return nil
	}

	if string(data) == "null" {
		*fs = ""
		return nil
	}

	return fmt.Errorf("cannot convert %s to string", string(data))
}

type StringList []string

func (l StringList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = v
	}
	return strings.Join(s, ",")
}

func (l StringList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
