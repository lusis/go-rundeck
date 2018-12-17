package rundeck

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	config := ClientConfig{
		BaseURL:    "http://localhost:4440/",
		Token:      "XXXXXXXXXXXXX",
		AuthMethod: "token",
	}
	client, err := NewClient(&config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.IsType(t, &http.Client{}, client.HTTPClient)
}

func TestNewClientSkipVerify(t *testing.T) {
	config := ClientConfig{
		VerifySSL:  false,
		BaseURL:    "http://localhost:4440/",
		Token:      "XXXXXXXXXXXXX",
		AuthMethod: "token",
	}
	client, err := NewClient(&config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	transport := client.HTTPClient.Transport.(*http.Transport)
	assert.True(t, transport.TLSClientConfig.InsecureSkipVerify)
}

func TestNewClientCustomHTTPClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string("Custom transport used")) //nolint: errcheck
	}))

	transport := http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := http.Client{}
	httpClient.Transport = &transport
	conf := ClientConfig{
		BaseURL:    "http://localhost:4440/",
		Token:      "XXXXXXXXXXXXX",
		VerifySSL:  false,
		AuthMethod: "token",
		HTTPClient: &httpClient,
	}
	client, err := NewClient(&conf)
	assert.NoError(t, err)
	assert.ObjectsAreEqual(httpClient, client.HTTPClient)
	res, resErr := client.Get("/foo", requestExpects(200))
	assert.NoError(t, resErr)
	assert.Equal(t, "Custom transport used", string(res))
}

func TestNewTokenAuthClient(t *testing.T) {
	client, err := NewTokenAuthClient("abcdefg", "http://localhost:4440")
	assert.NoError(t, err)
	assert.Equal(t, MaxRundeckVersion, client.Config.APIVersion)
	assert.NotNil(t, client.HTTPClient)
	assert.NotNil(t, client.Config.HTTPClient)
	assert.Equal(t, "token", client.Config.AuthMethod)
	assert.Equal(t, "abcdefg", client.Config.Token)
	assert.True(t, client.Config.VerifySSL)
	assert.Equal(t, "http://localhost:4440", client.Config.BaseURL)
}

func TestNewBasicAuthClient(t *testing.T) {
	client, err := NewBasicAuthClient("abcdefg", "12345", "http://localhost:4440")
	assert.NoError(t, err)
	assert.Equal(t, MaxRundeckVersion, client.Config.APIVersion)
	assert.NotNil(t, client.HTTPClient)
	assert.NotNil(t, client.Config.HTTPClient)
	assert.Equal(t, "basic", client.Config.AuthMethod)
	assert.Equal(t, "abcdefg", client.Config.Username)
	assert.Equal(t, "12345", client.Config.Password)
	assert.True(t, client.Config.VerifySSL)
	assert.Equal(t, "http://localhost:4440", client.Config.BaseURL)
}
func TestNewClientFromEnvToken(t *testing.T) {
	_ = os.Setenv("RUNDECK_TOKEN", "lK2iaQLEkf6rINMAYOXfrFNIpuwHRq67")
	_ = os.Setenv("RUNDECK_URL", "http://localhost:4440")
	defer func() { _ = os.Unsetenv("RUNDECK_TOKEN"); _ = os.Unsetenv("RUNDECK_URL") }()
	client, err := NewClientFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, "token", client.Config.AuthMethod)
}

func TestNewClientFromEnvBasic(t *testing.T) {
	_ = os.Setenv("RUNDECK_USERNAME", "admin")
	_ = os.Setenv("RUNDECK_PASSWORD", "admin")
	_ = os.Setenv("RUNDECK_URL", "http://localhost:4440")
	defer func() {
		_ = os.Unsetenv("RUNDECK_USERNAME")
		_ = os.Unsetenv("RUNDECK_PASSWORD")
		_ = os.Unsetenv("RUNDECK_URL")
	}()

	client, err := NewClientFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, "basic", client.Config.AuthMethod)
}

func TestNewClientFromEnvMissingPassword(t *testing.T) {
	_ = os.Setenv("RUNDECK_USERNAME", "admin")
	_ = os.Setenv("RUNDECK_URL", "http://localhost:4440")
	defer func() {
		_ = os.Unsetenv("RUNDECK_USERNAME")
		_ = os.Unsetenv("RUNDECK_URL")
	}()

	client, err := NewClientFromEnv()
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewClientFromEnvMissingURL(t *testing.T) {
	_ = os.Setenv("RUNDECK_USERNAME", "admin")
	_ = os.Setenv("RUNDECK_PASSWORD", "admin")
	defer func() {
		_ = os.Unsetenv("RUNDECK_USERNAME")
		_ = os.Unsetenv("RUNDECK_PASSWORD")
	}()

	client, err := NewClientFromEnv()
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewClientFromEnvMissingUsername(t *testing.T) {
	_ = os.Setenv("RUNDECK_PASSWORD", "admin")
	_ = os.Setenv("RUNDECK_URL", "http://localhost:4440")
	defer func() {
		_ = os.Unsetenv("RUNDECK_PASSWORD")
		_ = os.Unsetenv("RUNDECK_URL")
	}()

	client, err := NewClientFromEnv()
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewClientFromEnvSetVersion(t *testing.T) {
	_ = os.Setenv("RUNDECK_USERNAME", "admin")
	_ = os.Setenv("RUNDECK_PASSWORD", "admin")
	_ = os.Setenv("RUNDECK_URL", "http://localhost:4440")
	_ = os.Setenv("RUNDECK_VERSION", "18")
	defer func() {
		_ = os.Unsetenv("RUNDECK_PASSWORD")
		_ = os.Unsetenv("RUNDECK_URL")
	}()

	client, err := NewClientFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, "18", client.Config.APIVersion)
}
