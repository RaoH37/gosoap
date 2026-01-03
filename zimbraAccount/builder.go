package zimbraAccount

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

const Urn = "urn:zimbraAccount"

func NewAuthRequestByPassword(by zimbraCommon.ByNode, password string) (*AuthRequest, AuthResponse) {
	return &AuthRequest{
		Content: AuthRequestContent{
			Account:  by,
			Password: password,
			Urn:      Urn,
		},
	}, AuthResponse{}
}

func NewAuthRequestByPreauth(by zimbraCommon.ByNode, preauth Preauth) (*AuthRequest, AuthResponse) {
	return &AuthRequest{
		Content: AuthRequestContent{
			Account: by,
			Preauth: preauth,
			Urn:     Urn,
		},
	}, AuthResponse{}
}

func NewPreauth(by zimbraCommon.ByNode, expires int, domainKey string) Preauth {
	ts := time.Now().UnixMilli()
	data := fmt.Sprintf("%s|%s|%d|%d", by.Value, by.By, expires, ts)
	h := hmac.New(sha1.New, []byte(domainKey))
	h.Write([]byte(data))

	return Preauth{
		Timestamp: ts,
		Expires:   expires,
		Value:     hex.EncodeToString(h.Sum(nil)),
	}
}

func NewGetInfoRequest(rights zimbraCommon.StringList, sections InfoSectionList) (*GetInfoRequest, GetInfoResponse) {
	content := GetInfoRequestContent{
		Urn: Urn,
	}

	if len(rights) > 0 {
		content.Rights = rights
	}

	if len(sections) > 0 {
		content.Sections = sections
	}

	return &GetInfoRequest{
		Content: content,
	}, GetInfoResponse{}
}
