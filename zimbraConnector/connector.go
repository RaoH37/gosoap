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
	"sync"
	"time"

	"github.com/RaoH37/gosoap/zimbraCommon"
)

const Urn = "urn:zimbra"

// Cache pour la réflexion pour ne pas recalculer les noms à chaque fois
var nameCache sync.Map

type Connector struct {
	url       string
	userAgent string
	header    *HeaderContext
	client    *http.Client
	Debug     bool
	timeout   time.Duration
}

func NewConnector(url string, insecure bool, userAgent string, debug bool, timeout time.Duration) *Connector {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecure,
		},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
	}

	return &Connector{
		url:       url,
		userAgent: userAgent,
		timeout:   timeout,
		Debug:     debug,
		client:    &http.Client{Transport: tr, Timeout: timeout},
	}
}

func (connector *Connector) SetHeaderContext(token string, serverId string, accountContext *zimbraCommon.ByNode) {
	header := HeaderContext{
		Urn:      Urn,
		Token:    token,
		ServerID: serverId,
		Account:  accountContext,
	}

	if connector.userAgent != "" {
		header.UserAgent = &zimbraCommon.NameNode{Name: connector.userAgent}
	}

	connector.header = &header
}

func (connector *Connector) Invoke(request interface{}, response interface{}) error {
	envelope := connector.buildEnvelope(request)
	soapRequestName := getSoapRequestName(request)

	if connector.Debug {
		bb, _ := json.MarshalIndent(envelope, "", "  ")
		fmt.Printf(">>> REQUEST %s:\n%s\n", soapRequestName, string(bb))
	}

	return connector.doRequest(soapRequestName, envelope, response)
}

func (connector *Connector) buildEnvelope(request interface{}) Envelope {
	env := Envelope{Body: request}
	if connector.header != nil {
		env.Header = &Header{Content: connector.header}
	}
	return env
}

func (connector *Connector) doRequest(soapRequestName string, envelope Envelope, response interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(envelope); err != nil {
		return fmt.Errorf("failed to encode envelope: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connector.timeout)
	defer cancel()

	targetURL := fmt.Sprintf("%s/%s", connector.url, soapRequestName)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, &buf)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("SOAPAction", soapRequestName)
	if connector.userAgent != "" {
		req.Header.Set("User-Agent", connector.userAgent)
	}

	res, err := connector.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return connector.handleFault(res)
	}

	var bodyReader io.Reader = res.Body
	if connector.Debug {
		raw, _ := io.ReadAll(res.Body)
		fmt.Printf("<<< RESPONSE %s:\n%s\n", soapRequestName, string(raw))
		bodyReader = bytes.NewReader(raw)
	}

	respEnvelope := Envelope{Body: response}
	if err := json.NewDecoder(bodyReader).Decode(&respEnvelope); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

func (connector *Connector) handleFault(res *http.Response) error {
	raw, _ := io.ReadAll(res.Body)
	if connector.Debug {
		fmt.Printf("!!! FAULT: %s\n", string(raw))
	}

	var faultEnv Envelope
	fault := &Fault{}
	faultEnv.Body = fault

	if err := json.Unmarshal(raw, &faultEnv); err != nil {
		return fmt.Errorf("HTTP %d: %s", res.StatusCode, string(raw))
	}
	return fault
}

func getSoapRequestName(v interface{}) string {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if name, ok := nameCache.Load(t); ok {
		return name.(string)
	}

	name := t.Name()
	if name == "" {
		name = "Request"
	}

	nameCache.Store(t, name)

	return name
}
