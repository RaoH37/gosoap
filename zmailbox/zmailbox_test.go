package zmailbox_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/RaoH37/gosoap/zimbraAccount"
	"github.com/RaoH37/gosoap/zimbraCommon"
	"github.com/RaoH37/gosoap/zimbraMail"
	"github.com/RaoH37/gosoap/zmailbox"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var domainKey = os.Getenv("GOSOAP_DOMAIN_KEY")
var token = os.Getenv("GOSOAP_TOKEN")

var letters = os.Getenv("GOSOAP_LETTERS")

func RandStringRunes(n int) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}

	newStr := string(b)

	fmt.Println(newStr)

	return newStr
}

func NewZMailbox() (zmailbox.ZMailbox, error) {
	zcs := zmailbox.NewZMailbox(url, true, "", login, pwd, "", testing.Verbose(), time.Second*5, "zsoap", time.Second*30)

	if len(token) > 0 {
		zcs.SetToken(token)
	} else {
		resp, err := zcs.AuthRequest()

		if testing.Verbose() {
			fmt.Printf("%v\n", resp)
		}

		if err != nil {
			return zcs, err
		}

		zcs.SetToken(resp.Content.TOKEN[0].Content)
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

	resp, err := zcs.GetFolderRequest(zimbraMail.FolderViewMessage)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
		for _, folder := range resp.Content.Folders {
			fmt.Printf("%v\n", folder)
			if len(folder.Folders) > 0 {
				for _, subfolder := range folder.Folders {
					fmt.Printf("%v\n", subfolder)
				}
			}
		}
	}
}

func TestCreateFolderRequest(t *testing.T) {
	zcs, err := NewZMailbox()
	if err != nil {
		t.Fatalf("%v", err)
	}

	name := "test_" + RandStringRunes(6)

	resp, err := zcs.CreateFolderRequest(name, zimbraMail.FolderViewMessage, "1", func(req *zimbraMail.CreateFolderRequest) {
		req.Content.Folder.Color = 1
	})

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}

func TestGetInfoRequest(t *testing.T) {
	zcs, err := NewZMailbox()
	if err != nil {
		t.Fatalf("%v", err)
	}

	sections := zimbraAccount.InfoSectionList{
		zimbraAccount.InfoSectionAttrs,
		zimbraAccount.InfoSectionPrefs,
	}

	resp, err := zcs.GetInfoRequest(zimbraCommon.StringList{}, sections)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if testing.Verbose() {
		fmt.Printf("%v\n", resp)
	}
}
