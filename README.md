# gozsoap

This is a Zimbra client library written in Go.

Go 1.10 is required !

## Usage

```go
package main

import "fmt"
import "local/gozsoap"

func main() {
  url := "https://server.domain.tld:7071/service/admin/soap"
  login := "admin@domain.tld"
  pwd := "secret"

  zcs := zsoap.ZAdmin{}
  zcs.Init(url, true)
  zcs.Login(login, pwd)

  byDomain := zsoap.NewByRequest("name", "domain.tld")

  accounts := zcs.GetAllAccounts(nil, &byDomain)
  fmt.Printf("%v", accounts)

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