package zsoap

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const urnAdmin = "urn:zimbraAdmin"

// Envelope envelope
type Envelope struct {
	Header *Header     `json:",omitempty"`
	Body   interface{} `json:"Body,omitempty"`
}

// Header header
type Header struct {
	Content interface{} `json:"context,omitempty"`
}

type HeaderToken struct {
	TOKEN    string `json:"authToken"`
	Urn      string `json:"_jsns,attr"`
	ServerID string `json:"targetServer,omitempty"`
}

// Fault fault
type Fault struct {
	Content FaultContent `json:"Fault,omitempty"`
}
type FaultContent struct {
	Code   interface{} `json:"Code,omitempty"`
	Reason FaultReason `json:"Reason,omitempty"`
	Detail interface{} `json:"Detail,omitempty"`
}
type FaultReason struct {
	Text string `json:"Text,omitempty"`
}

func (f *Fault) Error() string {
	return f.Content.Reason.Text
}

// NewClient return SOAP client
func NewClient(url string, tls bool, header interface{}) *Client {
	return &Client{
		url:    url,
		tls:    tls,
		header: header,
		Debug:  false,
	}
}

// Client SOAP client
type Client struct {
	url       string
	tls       bool
	userAgent string
	header    interface{}
	TOKEN     string
	Debug     bool
	Timeout   time.Duration
}

func dialTimeout(network, addr string) (net.Conn, error) {
	timeout := time.Duration(30 * time.Second)
	return net.DialTimeout(network, addr, timeout)
}

func (s *Client) SetToken(token string) {
	s.TOKEN = token
	s.SetHeader()
}

func (s *Client) SetHeader() {
	s.header = HeaderToken{
		TOKEN: s.TOKEN,
		Urn:   "urn:zimbra",
	}
}

func (s *Client) SetTargetServer(serverID string) {
	s.header = HeaderToken{
		TOKEN:    s.TOKEN,
		Urn:      "urn:zimbra",
		ServerID: serverID,
	}
}

func (s *Client) RemoveTargetServer() {
	s.SetHeader()
}

// Call SOAP client API call
func (s *Client) Call(soapAction string, request interface{}, response interface{}) error {
	var envelope Envelope
	if s.header != nil {
		envelope = Envelope{
			Header: &Header{
				Content: s.header,
			},
			Body: request,
		}
	} else {
		envelope = Envelope{
			Body: request,
		}
	}

	if s.Debug {
		bb, _ := json.Marshal(envelope)
		fmt.Println(string(bb))
	}

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	if err := encoder.Encode(envelope); err != nil {
		return errors.Wrap(err, "failed to encode envelope")
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return errors.Wrap(err, "failed to create POST request")
	}

	req.Header.Add("Content-Type", "application/json; charset=\"utf-8\"")
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Set("User-Agent", s.userAgent)
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr, Timeout: s.Timeout}
	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send SOAP request")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {

		soapFault, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.Wrap(err, "failed to read SOAP fault response body")
		}

		if s.Debug {
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

	rawbody, err := io.ReadAll(res.Body)

	if s.Debug {
		fmt.Println(string(rawbody))
	}

	if err != nil {
		return errors.Wrap(err, "failed to read SOAP body")
	}

	if len(rawbody) == 0 {
		return nil
	}

	respEnvelope := Envelope{Body: response}

	if err = json.Unmarshal(rawbody, &respEnvelope); err != nil {
		return errors.Wrap(err, "failed to unmarshal response SOAP Envelope")
	}

	return nil
}
