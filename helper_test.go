package zsoap_test

import (
	"os"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var token = os.Getenv("GOSOAP_TOKEN")
