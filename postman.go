package fingerplus

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type method int

// Method type
const (
	MethodGet method = iota
	MethodPost
	MethodPostRaw
	MethodPostJSON
	MethodPostXML
)

// defaultHTTPClient is reusable global object of http client
var defaultHTTPClient = &http.Client{
	Timeout: time.Second * 30,
}

// Postman will help you to send your request
type Postman struct {
	Client *http.Client

	Method method
	Header http.Header

	URL    string
	Params url.Values
	Body   interface{}
}

// NewPostman returns new Postman instance
func NewPostman() *Postman {
	return &Postman{}
}

// Send the request without context
func (p *Postman) Send() (*http.Response, []byte, error) {
	return p.send(nil)
}

// SendWithContext sends the request with context
func (p *Postman) SendWithContext(ctx context.Context) (*http.Response, []byte, error) {
	return p.send(ctx)
}

// send the request
func (p *Postman) send(ctx context.Context) (*http.Response, []byte, error) {
	// parse url string
	url, err := url.Parse(p.URL)
	if err != nil {
		return nil, nil, err
	}

	// get *http.Request by method type
	request, err := p.getRequest(url)
	if err != nil {
		return nil, nil, err
	}
	if request == nil {
		return nil, nil, errors.New("failed to create new request")
	}

	// use context, if exist
	if ctx != nil {
		request.WithContext(ctx)
	}

	// set request headers
	for key, value := range p.Header {
		request.Header.Set(key, strings.Join(value, ","))
	}

	// close the request after finish
	request.Close = true

	// send the request
	response, err := p.getClient().Do(request)
	if err != nil {
		return response, nil, err
	}
	if response != nil {
		defer response.Body.Close()
	}

	// read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response, nil, err
	}

	return response, body, err
}

func (p *Postman) getRequest(url *url.URL) (*http.Request, error) {
	switch p.Method {
	case MethodGet:
		return p.get(url)
	case MethodPost, MethodPostRaw, MethodPostJSON, MethodPostXML:
		return p.post(p.Method, url)
	}

	return nil, errors.New("invalid method type")
}

func (p *Postman) getClient() *http.Client {
	if p.Client != nil {
		return p.Client
	}

	return defaultHTTPClient
}

func (p *Postman) get(url *url.URL) (*http.Request, error) {
	if p.Params != nil {
		url.RawQuery = p.Params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (p *Postman) post(m method, url *url.URL) (*http.Request, error) {
	var body io.Reader
	var err error

	switch m {
	case MethodPost:
		body = strings.NewReader(p.Params.Encode())
	case MethodPostRaw:
		body, err = p.getBodyRaw()
	case MethodPostJSON:
		p.Header.Add("Content-Type", "application/json")
		body, err = p.getBodyJSON()
	case MethodPostXML:
		body, err = p.getBodyXML()
	}

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (p *Postman) getBodyRaw() (io.Reader, error) {
	switch p.Body.(type) {
	case []byte:
		return bytes.NewBuffer(p.Body.([]byte)), nil
	case string:
		return bytes.NewBuffer([]byte(p.Body.(string))), nil
	}

	return nil, errors.New("invalid raw data")
}

func (p *Postman) getBodyJSON() (io.Reader, error) {
	body, err := json.Marshal(p.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}

func (p *Postman) getBodyXML() (io.Reader, error) {
	body, err := xml.Marshal(p.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}
