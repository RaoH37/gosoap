package zimbraAccount

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

const urnAccount = "urn:zimbraAccount"

func NewAuthRequestByPassword(by zimbraCommon.ByNode, password string) (*AuthRequest, AuthResponse) {
	return &AuthRequest{
		Content: AuthRequestContent{
			Account:  by,
			Password: password,
			Urn:      urnAccount,
		},
	}, AuthResponse{}
}

func NewAuthRequestByPreauth(by zimbraCommon.ByNode, preauth Preauth) (*AuthRequest, AuthResponse) {
	return &AuthRequest{
		Content: AuthRequestContent{
			Account: by,
			Preauth: &preauth,
			Urn:     urnAccount,
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
