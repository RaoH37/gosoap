package zimbraConnector

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/RaoH37/gosoap/zimbraCommon"
	"github.com/pkg/errors"
)

type Connector struct {
	url       string
	tls       bool
	userAgent string
	header    interface{}
	token     string
	Debug     bool
	timeout   time.Duration
}

func (connector *Connector) resetHeaderContext() {
	connector.header = nil
}

func (connector *Connector) SetHeaderContext(token string, serverId string, accountName string, userAgent string) {
	headerContext := HeaderContext{
		Urn:      "urn:zimbra",
		Token:    token,
		ServerID: serverId,
	}

	if accountName != "" {
		by := zimbraCommon.NewByNode(zimbraCommon.NAME, accountName)
		headerContext.Account = &by
	}

	if userAgent != "" {
		headerContext.UserAgent = &NameNode{Name: userAgent}
	}

	connector.header = headerContext
}

func (connector *Connector) Invoke(request interface{}, response interface{}) error {
	envelope := connector.buildEnvelope(request)

	if connector.Debug {
		bb, _ := json.Marshal(envelope)
		fmt.Println(string(bb))
	}

	soapRequestName := getSoapRequestName(request)

	return connector.doRequest(soapRequestName, envelope, response)
}

func (connector *Connector) buildEnvelope(request interface{}) Envelope {
	var envelope Envelope

	if connector.header != nil {
		envelope = Envelope{
			Header: &Header{
				Content: connector.header,
			},
			Body: request,
		}
	} else {
		envelope = Envelope{
			Body: request,
		}
	}

	return envelope
}

func (connector *Connector) doRequest(soapRequestName string, envelope Envelope, response interface{}) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(envelope); err != nil {
		return errors.Wrap(err, "failed to encode envelope")
	}

	ctx, cncl := context.WithTimeout(context.Background(), connector.timeout)
	defer cncl()

	url := connector.url + "/" + soapRequestName

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buffer)
	if err != nil {
		return errors.Wrap(err, "failed to create POST request")
	}

	req.Header.Add("Content-Type", "application/json; charset=\"utf-8\"")
	req.Header.Set("SOAPAction", soapRequestName)
	req.Header.Set("User-Agent", connector.userAgent)
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: connector.tls,
		},
	}

	client := &http.Client{Transport: tr, Timeout: connector.timeout}
	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send SOAP request")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {

		soapFault, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.Wrap(err, "failed to read SOAP fault response body")
		}

		if connector.Debug {
			fmt.Println(string(soapFault))
		}

		var msg string
		fault := Fault{}
		faultEnvelope := Envelope{Body: &fault}

		if err = json.Unmarshal(soapFault, &faultEnvelope); err != nil {
			msg = fmt.Sprintf("HTTP Status Code: %d, SOAP Fault: \n%s", res.StatusCode, string(soapFault))
		} else {
			msg = fault.Error()
		}

		return errors.New(msg)
	}

	rawBody, err := io.ReadAll(res.Body)

	if connector.Debug {
		fmt.Println(string(rawBody))
	}

	if err != nil {
		return errors.Wrap(err, "failed to read SOAP body")
	}

	if len(rawBody) == 0 {
		return nil
	}

	respEnvelope := Envelope{Body: response}

	if err = json.Unmarshal(rawBody, &respEnvelope); err != nil {
		return errors.Wrap(err, "failed to unmarshal response SOAP Envelope")
	}

	return nil
}

func getSoapRequestName(envelope interface{}) string {
	t := reflect.TypeOf(envelope)

	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}
