package zimbraConnector

import "time"

func BuildConnector(
	url string,
	tls bool,
	userAgent string,
	header interface{},
	Debug bool,
	timeout time.Duration) *Connector {
	return &Connector{
		url:       url,
		tls:       tls,
		userAgent: userAgent,
		header:    header,
		Debug:     Debug,
		timeout:   timeout,
	}
}
