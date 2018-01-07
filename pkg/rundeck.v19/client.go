package rundeck

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	napping "gopkg.in/jmcvetta/napping.v2"
)

// MaxRundeckVersion is the maximum version of the api this library supports
// can be overridden
const MaxRundeckVersion = "19"

// ClientConfig represents a client configuration
type ClientConfig struct {
	BaseURL    string
	Token      string
	VerifySSL  bool
	Username   string
	Password   string
	AuthMethod string
	APIVersion string
	Transport  *http.Transport
	HTTPClient *http.Client
}

// Client represents a rundeck client
type Client struct {
	Client     *napping.Session
	HTTPClient *http.Client
	Config     *ClientConfig
	Transport  *http.Transport
}

// NewClient creates a new client from the provided `ClientConfig`
func NewClient(config *ClientConfig) (c Client) {
	verifySSL := func() bool {
		if !config.VerifySSL { // nolint: gosimple
			return false
		}
		return true
	}
	if config.Transport == nil {
		config.Transport = new(http.Transport)
	}
	config.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: verifySSL()}
	if config.HTTPClient == nil {
		config.HTTPClient = new(http.Client)
	}
	config.HTTPClient.Transport = config.Transport
	s := napping.Session{
		Client: config.HTTPClient,
	}
	return Client{Client: &s, Config: config}
}

func clientConfigFrom(from string) (c *ClientConfig) {
	config := ClientConfig{}

	switch from {
	case "environment":
		if os.Getenv("RUNDECK_TOKEN") == "" {
			if os.Getenv("RUNDECK_USER") == "" && os.Getenv("RUNDECK_PASSWORD") == "" {
				fmt.Printf("You must set either RUNDECK_TOKEN or RUNDECK_USERNAME and RUNDECK_PASSWORD\n")
				os.Exit(1)
			} else {
				config.AuthMethod = "basic"
			}
		} else {
			config.AuthMethod = "token"
		}
		if os.Getenv("RUNDECK_VERSION") == "" {
			config.APIVersion = MaxRundeckVersion
		} else {
			config.APIVersion = os.Getenv("RUNDECK_VERSION")
		}

		if os.Getenv("RUNDECK_URL") == "" {
			fmt.Printf("You must set the environment variable RUNDECK_URL\n")
			os.Exit(1)
		} else {
			config.BaseURL = os.Getenv("RUNDECK_URL")
		}
	}
	if config.AuthMethod == "token" {
		config.Token = os.Getenv("RUNDECK_TOKEN")
	} else {
		config.Username = os.Getenv("RUNDECK_USERNAME")
		config.Password = os.Getenv("RUNDECK_PASSWORD")
	}
	return &config
}

// NewClientFromEnv returns a new client from provided env vars
func NewClientFromEnv() (c Client) {
	config := clientConfigFrom("environment")
	client := NewClient(config)
	return client
}
