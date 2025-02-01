# gozsoap

This is a Zimbra client library written in Go.

## Usage

```go
package main

import (
	"fmt"
	"github.com/RaoH37/gosoap"
)

func main() {
  url := "https://server.domain.tld:7071/service/admin/soap"
  login := "admin@domain.tld"
  pwd := "secret"

  zcs := zsoap.ZAdmin{}
  zcs.Init(url, true)
  zcs.Login(login, pwd)

  accounts, _ := zcs.GetAllAccounts("", "", 1, 1, "", 1, "")
  for _, account := range accounts {
    fmt.Printf("%v\n", account)
  }

  byAccount := zsoap.NewByRequest("name", "foo@domain.tld")
  account := zcs.GetAccount(byAccount, []string{"zimbraMailHost", "zimbraMailTransport"})
  fmt.Printf("%v", account)

  servers := zcs.GetAllServers("")
  fmt.Printf("%v",servers)

  domains := zcs.GetAllDomains()
  fmt.Printf("%v", domains)

  coses := zcs.GetAllCos()
  fmt.Printf("%v", coses)

  qaccounts := zcs.GetQuotaUsage("domain.tld")
  fmt.Printf("%v", qaccounts)
}
```