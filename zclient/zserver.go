package zclient

import "time"

type ZServer struct {
	ID                    string
	Name                  string
	ZimbraCreateTimestamp string
	ZimbraServiceEnabled  []string
	ZimbraSmtpHostname    []string
	ZimbraSmtpPort        int
}

func (server *ZServer) CreatedAt() int {
	t, err := time.Parse("20060102150405Z", server.ZimbraCreateTimestamp)

	if err != nil {
		return -1
	}

	return int(t.Unix())
}
