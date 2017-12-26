package rundeck

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func newTestRundeckClient(content []byte, contentType string, statusCode int) (*Client, *httptest.Server, error) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", contentType)
		fmt.Fprintf(w, string(content))
	}))

	transport := http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := http.Client{}
	httpClient.Transport = &transport
	conf := &ClientConfig{
		BaseURL:    "http://localhost:4440/",
		Token:      "XXXXXXXXXXXXX",
		VerifySSL:  false,
		AuthMethod: "token",
		HTTPClient: &httpClient,
	}
	client, err := NewClient(conf)
	if err != nil {
		return nil, nil, err
	}
	return client, server, nil
}
