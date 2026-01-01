package zimbraConnector

import "github.com/RaoH37/gosoap/zimbraCommon"

type Envelope struct {
	Header *Header     `json:",omitempty"`
	Body   interface{} `json:"Body,omitempty"`
}

type Header struct {
	Content interface{} `json:"context,omitempty"`
}

type HeaderToken struct {
	TOKEN    string `json:"authToken"`
	Urn      string `json:"_jsns,attr"`
	ServerID string `json:"targetServer,omitempty"`
}

type HeaderContext struct {
	Urn       string               `json:"_jsns,attr"`
	Token     string               `json:"authToken,omitempty"`
	UserAgent *NameNode            `json:"userAgent,omitempty"`
	Account   *zimbraCommon.ByNode `json:"account,omitempty"`
	ServerID  string               `json:"targetServer,omitempty"`
}

type NameNode struct {
	Name string `json:"name,omitempty"`
}

type Fault struct {
	Content FaultContent `json:"Fault,omitempty"`
}

type FaultContent struct {
	Code   interface{} `json:"Code,omitempty"`
	Reason FaultReason `json:"Reason,omitempty"`
	Detail interface{} `json:"Detail,omitempty"`
}

type FaultReason struct {
	Text string `json:"Text,omitempty"`
}

func (f *Fault) Error() string {
	return f.Content.Reason.Text
}
