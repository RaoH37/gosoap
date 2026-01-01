package zimbraCommon

import (
	"encoding/json"
	"fmt"
	"strconv"
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

type UrnRequestContent struct {
	Urn string `json:"_jsns,attr"`
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

	return fmt.Errorf("impossible de convertir %s en string", string(data))
}
