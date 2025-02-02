package zsoap_test

import (
	"fmt"
	zsoap "github.com/RaoH37/gosoap"
	"os"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var token = os.Getenv("GOSOAP_TOKEN")

func NewZcsClient() (zsoap.ZAdmin, error) {
	zcs := zsoap.ZAdmin{}
	zcs.Init(url, true)

	if len(token) > 0 {
		zcs.Client.SetToken(token)
	} else {
		err := zcs.Login(login, pwd)
		if err != nil {
			return zcs, err
		}
	}

	fmt.Println(zcs.Client.TOKEN)

	return zcs, nil
}
