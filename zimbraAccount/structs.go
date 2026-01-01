package zimbraAccount

import "github.com/RaoH37/gosoap/zimbraCommon"

type AuthRequest struct {
	Content AuthRequestContent `json:"AuthRequest"`
}

type AuthRequestContent struct {
	Account  zimbraCommon.ByNode `json:"account,attr"`
	Password string              `json:"password,attr,omitempty"`
	Preauth  *Preauth            `json:"preauth,omitempty"`
	Urn      string              `json:"_jsns,attr"`
}

type Preauth struct {
	Timestamp int64  `json:"timestamp,attr"`
	Expires   int    `json:"expires,attr"`
	Value     string `json:"_content,attr"`
}

type AuthResponse struct {
	Content AuthResponseContent `json:"AuthResponse"`
}

type AuthResponseContent struct {
	TOKEN    []AuthResponseToken `json:"authToken"`
	Lifetime int                 `json:"lifetime"`
}

type AuthResponseToken struct {
	Content string `json:"_content"`
}
