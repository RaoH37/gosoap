package zimbraCommon

import (
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Token struct {
	raw      string
	keyID    string
	hmac     string
	encoded  string
	metadata map[string]string
}

func NewToken(str string) *Token {
	parts := strings.Split(str, "_")
	if len(parts) < 3 {
		return &Token{raw: str} // Retourne un token potentiellement invalide si format incorrect
	}

	t := &Token{
		raw:     str,
		keyID:   parts[0],
		hmac:    parts[1],
		encoded: parts[2],
	}

	t.parseMetadata()
	return t
}

func (t *Token) String() string {
	return t.raw
}

func (t *Token) parseMetadata() {
	t.metadata = make(map[string]string)

	decodedBytes, err := hex.DecodeString(t.encoded)
	if err != nil {
		return
	}
	decodedStr := string(decodedBytes)

	parts := strings.Split(decodedStr, ";")

	re := regexp.MustCompile(`=\d+:`)

	for _, part := range parts {
		kv := re.Split(part, 2)
		if len(kv) == 2 {
			t.metadata[kv[0]] = kv[1]
		}
	}
}

func (t *Token) ZimbraID() string {
	return t.metadata["id"]
}

func (t *Token) IsAdmin() bool {
	return t.metadata["admin"] == "1"
}

func (t *Token) ServerVersion() string {
	return t.metadata["version"]
}

func (t *Token) ExpireAt() time.Time {
	expStr, ok := t.metadata["exp"]
	if !ok {
		return time.Time{}
	}

	ms, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.UnixMilli(ms)
}

func (t *Token) IsExpired() bool {
	if len(t.metadata) == 0 {
		return true
	}
	return t.ExpireAt().Before(time.Now())
}
