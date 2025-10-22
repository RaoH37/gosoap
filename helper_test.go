package zsoap_test

import (
	"os"
)

var url = os.Getenv("GOSOAP_URL")
var login = os.Getenv("GOSOAP_LOGIN")
var pwd = os.Getenv("GOSOAP_PWD")
var token = os.Getenv("GOSOAP_TOKEN")

var dlName = os.Getenv("GOSOAP_DL_NAME")
var domainName = os.Getenv("GOSOAP_DOMAIN_NAME")
var accountName = os.Getenv("GOSOAP_ACCOUNT_NAME")
