package httpclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/publicsuffix"
)

func testCustomOption() RequestOption {
	return func(r *Request) error {
		return errors.New("i blew up")
	}
}

type testHTPPBinResponse struct {
	Args    map[string]string `json:"args,omitempty"`
	Headers map[string]string `json:"headers"`
	Form    map[string]string `json:"form"`
	URL     string            `json:"url"`
	Data    string            `json:"data,omitempty"`
}

type testHTTPBinCookieResponse struct {
	Cookies map[string]string `json:"cookies"`
}

func TestNew(t *testing.T) {
	c, r, err := New()
	assert.NoError(t, err)
	assert.IsType(t, &http.Request{}, r)
	assert.Len(t, c.allowedStatusCodes, 0)
	assert.Equal(t, DefaultAccept, c.accept)
	assert.Equal(t, c.httpClient.Timeout, http.DefaultClient.Timeout)
}

func TestNewWithOpt(t *testing.T) {
	c, r, err := New(ExpectStatus(200, 302))
	assert.NoError(t, err)
	assert.IsType(t, &http.Request{}, r)
	assert.Len(t, c.allowedStatusCodes, 2)
}

func TestCustomHTTPClient(t *testing.T) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	c, r, err := New(SetClient(client))
	assert.NoError(t, err)
	assert.IsType(t, &http.Request{}, r)
	assert.Equal(t, 15*time.Second, c.httpClient.Timeout)
}

func TestCookieJarDefault(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "foocookiekey", Value: "foocookievalue"})
	}))
	defer ts.Close()
	url, _ := url.Parse(ts.URL)
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	resp, err := Get(ts.URL)
	assert.NoError(t, err)
	assert.Equal(t, "foocookievalue", resp.Cookies[0].Value)
	assert.Len(t, jar.Cookies(url), 0)
}

func TestCookieJarCustom(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "foocookiekey", Value: "foocookievalue"})
	}))
	defer ts.Close()
	url, _ := url.Parse(ts.URL)
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	resp, err := Get(ts.URL, SetCookieJar(jar))
	assert.NoError(t, err)
	assert.Equal(t, "foocookievalue", resp.Cookies[0].Value)
	assert.Len(t, jar.Cookies(url), 1)
}

func TestErrOpt(t *testing.T) {
	c, r, err := New(testCustomOption())
	assert.Nil(t, c)
	assert.Nil(t, r)
	assert.Error(t, err)
}
func TestAddHeaders(t *testing.T) {
	headers := map[string]string{
		"fooheader": "foovalue",
		"barheader": "barvalue",
	}
	response, err := Get("https://httpbin.org/anything", AddHeaders(headers))
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "foovalue", res.Headers["Fooheader"])
	assert.Equal(t, "barvalue", res.Headers["Barheader"])
}
func TestAccept(t *testing.T) {
	response, err := Get("https://httpbin.org/anything", Accept("application/octet"))
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "application/octet", res.Headers["Accept"])
}

func TestRequestXML(t *testing.T) {
	response, err := Get("https://httpbin.org/anything", RequestXML())
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "application/xml", res.Headers["Accept"])
}
func TestGetAllowedStatusCodesInvalid(t *testing.T) {
	response, err := Get("https://httpbin.org/anything", ExpectStatus(302))
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidStatusCode.Error())
	assert.Equal(t, 200, response.Status)
}

func TestGetAllowedStatusCodesValid(t *testing.T) {
	response, err := Get("https://httpbin.org/anything", ExpectStatus(200, 302))
	assert.NoError(t, err)
	assert.Equal(t, 200, response.Status)
}

func TestGet(t *testing.T) {
	qp := make(map[string]string)
	qp["foo"] = "bar"
	response, err := Get("https://httpbin.org/get")
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "https://httpbin.org/get", res.URL)
}

func TestGetWithOption(t *testing.T) {
	qp := make(map[string]string)
	qp["foo"] = "bar"
	response, err := Get("https://httpbin.org/get", QueryParams(qp))
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "bar", res.Args["foo"])
	assert.Equal(t, "https://httpbin.org/get?foo=bar", res.URL)
}

func TestGetWithMultipleOptions(t *testing.T) {
	qp := make(map[string]string)
	qp["foo"] = "bar"
	response, err := Get("https://httpbin.org/get", QueryParams(qp), JSON())
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "bar", res.Args["foo"])
	assert.Equal(t, "https://httpbin.org/get?foo=bar", res.URL)
	assert.Equal(t, "application/json", res.Headers["Accept"])
}

func TestHead(t *testing.T) {
	response, err := Head("https://httpbin.org/ip")
	assert.NoError(t, err)
	assert.Equal(t, "application/json", response.Headers.Get("Content-Type"))
}

func TestDelete(t *testing.T) {
	response, err := Delete("https://httpbin.org/delete")
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, 200, response.Status)
}

func TestPost(t *testing.T) {
	response, err := Post("https://httpbin.org/post", WithBody(strings.NewReader("this is my body")), ContentType("text/plain"))
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "this is my body", res.Data)
}

func TestPut(t *testing.T) {
	response, err := Put("https://httpbin.org/put", WithBody(strings.NewReader("this is my body")), ContentType("text/plain"))
	assert.NoError(t, err)
	res := &testHTPPBinResponse{}
	jErr := json.Unmarshal(response.Body, &res)
	assert.NoError(t, jErr)
	assert.Equal(t, "this is my body", res.Data)
}
