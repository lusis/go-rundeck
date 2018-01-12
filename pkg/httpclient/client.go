package httpclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	"golang.org/x/net/publicsuffix"
)

// Response represents an http response
type Response struct {
	Body    []byte
	Headers http.Header
	Cookies []*http.Cookie
	Status  int
}

// Request represents an http request
type Request struct {
	httpClient         *http.Client
	cookieJar          *cookiejar.Jar
	url                string
	method             string
	contentType        string
	accept             string
	queryParams        map[string]string
	body               io.Reader
	headers            map[string]string
	allowedStatusCodes []int
	sync.RWMutex
}

// RequestOption is a type for functional options
type RequestOption func(*Request) error

func (cr *Request) setAllowedStatusCode(i int) {
	cr.allowedStatusCodes = append(cr.allowedStatusCodes, i)
}

func (cr *Request) getAllowedStatusCodes() []int {
	return cr.allowedStatusCodes
}

func (cr *Request) setHTTPClient(c *http.Client) {
	cr.httpClient = c
}

// AddHeaders adds custom headers to the request
func AddHeaders(h ...map[string]string) RequestOption {
	return func(r *Request) error {
		for _, pair := range h {
			for k, v := range pair {
				r.headers[k] = v
			}
		}
		return nil
	}
}

// SetClient sets a custom http.Client to use for the request
func SetClient(client *http.Client) RequestOption {
	return func(r *Request) error {
		r.setHTTPClient(client)
		return nil
	}
}

// QueryParams sets the query params for a request
func QueryParams(m map[string]string) RequestOption {
	return func(r *Request) error {
		r.queryParams = m
		return nil
	}
}

// SetCookieJar sets the cookie jar to be used with requests
func SetCookieJar(jar *cookiejar.Jar) RequestOption {
	return func(r *Request) error {
		r.cookieJar = jar
		return nil
	}
}

// JSON sets a request to accept and respond with json
func JSON() RequestOption {
	return func(r *Request) error {
		r.accept = ContentTypeJSON
		r.contentType = ContentTypeJSON
		return nil
	}
}

// ContentType allows setting the content-type for the request
func ContentType(ct string) RequestOption {
	return func(r *Request) error {
		r.contentType = ct
		return nil
	}
}

// Accept allows setting the Accept header individually
func Accept(ct string) RequestOption {
	return func(r *Request) error {
		r.accept = ct
		return nil
	}
}

// RequestXML sets a request to accept and respond with json
func RequestXML() RequestOption {
	return func(r *Request) error {
		r.accept = ContentTypeXML
		r.contentType = ContentTypeXML
		return nil
	}
}

// ExpectStatus sets expected status codes from a response
func ExpectStatus(codes ...int) RequestOption {
	return func(r *Request) error {
		for _, code := range codes {
			r.setAllowedStatusCode(code)
		}
		return nil
	}
}

// WithBody provides the body to be used with the http request
func WithBody(reader io.Reader) RequestOption {
	return func(r *Request) error {
		r.body = reader
		return nil
	}
}

// New creates a ClientRequest
func New(opts ...RequestOption) (*Request, *http.Request, error) {
	return newHTTPRequest(opts...)
}

func delete() RequestOption {
	return func(r *Request) error {
		r.method = http.MethodDelete
		return nil
	}
}

func get() RequestOption {
	return func(r *Request) error {
		r.method = http.MethodGet
		return nil
	}
}

func put() RequestOption {
	return func(r *Request) error {
		r.method = http.MethodPut
		return nil
	}
}

func post() RequestOption {
	return func(r *Request) error {
		r.method = http.MethodPost
		return nil
	}
}

func head() RequestOption {
	return func(r *Request) error {
		r.method = http.MethodHead
		return nil
	}
}
func setURL(u string) RequestOption {
	return func(r *Request) error {
		r.url = u
		return nil
	}
}

// newHTTPRequest returns a new `Request` configured with various options
func newHTTPRequest(opts ...RequestOption) (*Request, *http.Request, error) {
	r := &Request{}
	if r.httpClient == nil {
		r.setHTTPClient(&http.Client{})
	}
	codes := make([]int, 0)
	headers := make(map[string]string)
	r.allowedStatusCodes = codes
	r.headers = headers
	jar, jarErr := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if jarErr != nil {
		return nil, nil, jarErr
	}
	r.cookieJar = jar
	for _, opt := range opts {
		r.Lock()
		if err := opt(r); err != nil {
			r.Unlock()
			return nil, nil, err
		}
		r.Unlock()
	}

	req, err := r.httpRequest()

	return r, req, err
}

func (cr *Request) httpRequest() (*http.Request, error) {

	if cr.accept == "" {
		cr.accept = DefaultAccept
	}

	u, uErr := url.Parse(cr.url)
	if uErr != nil {
		return nil, uErr
	}

	if len(cr.queryParams) > 0 {
		qs := url.Values{}
		for q, p := range cr.queryParams {
			qs.Add(q, p)
		}
		u.RawQuery = qs.Encode()
	}

	req, reqErr := http.NewRequest(cr.method, u.String(), cr.body)

	if reqErr != nil {
		return nil, reqErr
	}

	for k, v := range cr.headers {
		req.Header.Add(k, v)
	}

	if cr.contentType != "" {
		req.Header.Add("Content-Type", cr.contentType)
	}
	req.Header.Add("Accept", cr.accept)
	return req, nil
}

// Get performs an http GET
func Get(url string, opts ...RequestOption) (*Response, error) {
	doOpts := []RequestOption{
		get(),
		setURL(url),
	}
	doOpts = append(doOpts, opts...)
	return doRequest(doOpts...)
}

// Delete performs an http DELETE
func Delete(url string, opts ...RequestOption) (*Response, error) {
	opts = append(opts, delete())
	opts = append(opts, setURL(url))
	return doRequest(opts...)
}

// Post performs an http POST
func Post(url string, opts ...RequestOption) (*Response, error) {
	opts = append(opts, post())
	opts = append(opts, setURL(url))
	return doRequest(opts...)
}

// Put performs an http PUT
func Put(url string, opts ...RequestOption) (*Response, error) {
	doOpts := []RequestOption{
		put(),
		setURL(url),
	}
	doOpts = append(doOpts, opts...)
	//opts = append(opts, put())
	//opts = append(opts, setURL(url))
	return doRequest(doOpts...)
}

// Head performs an http HEAD
func Head(url string, opts ...RequestOption) (*Response, error) {
	opts = append(opts, head())
	opts = append(opts, setURL(url))
	return doRequest(opts...)
}

func doRequest(opts ...RequestOption) (*Response, error) {
	response := &Response{}
	cr, req, reqErr := newHTTPRequest(opts...)
	if reqErr != nil {
		return nil, reqErr
	}
	cr.httpClient.Jar = cr.cookieJar

	resp, respErr := cr.httpClient.Do(req)
	if respErr != nil {
		return nil, respErr
	}
	readBody, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	response.Body = readBody
	response.Headers = resp.Header
	response.Status = resp.StatusCode
	response.Cookies = append(response.Cookies, resp.Cookies()...)
	if len(cr.getAllowedStatusCodes()) != 0 {
		passed := false
		for _, code := range cr.getAllowedStatusCodes() {
			if resp.StatusCode == code {
				passed = true
				break
			}
		}
		if !passed {
			return response, multierror.Append(ErrInvalidStatusCode, fmt.Errorf("status code %d", response.Status))
		}

	}

	return response, nil
}
