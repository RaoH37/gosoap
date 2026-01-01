package zmailbox_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/RaoH37/gosoap/zmailbox"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var domainKey = os.Getenv("GOSOAP_DOMAIN_KEY")
var token = os.Getenv("GOSOAP_TOKEN")

func NewZMailbox() (zmailbox.ZMailbox, error) {
	zcs := zmailbox.NewZMailbox(url, true, "", login, "", domainKey, testing.Verbose(), time.Second*5, "zsoap", time.Second*30)

	if len(token) > 0 {
		zcs.AuthToken = token
	} else {
		resp, err := zcs.AuthRequest()

		if testing.Verbose() {
			fmt.Printf("%v\n", resp)
		}

		if err != nil {
			return zcs, err
		}

		zcs.AuthToken = resp.Content.TOKEN[0].Content
	}

	return zcs, nil
}

func TestAuth(t *testing.T) {
	_, err := NewZMailbox()
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestGetFolderRequest(t *testing.T) {
	zcs, err := NewZMailbox()
	if err != nil {
		t.Fatalf("%v", err)
	}

	resp, err := zcs.GetFolderRequest()

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}
