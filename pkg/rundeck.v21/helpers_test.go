package rundeck

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func newTestRundeckClient(content io.Reader, contentType string, statusCode int) (*Client, *httptest.Server, error) {
	data, dataErr := ioutil.ReadAll(content)
	if dataErr != nil {
		return nil, nil, dataErr
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", contentType)
		fmt.Fprintf(w, string(data))
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	conf := &ClientConfig{
		BaseURL:    "http://127.0.0.1:8080/",
		Token:      "XXXXXXXXXXXXX",
		VerifySSL:  false,
		AuthMethod: "token",
		Transport:  transport,
	}
	client := NewClient(conf)
	return client, server, nil
}
