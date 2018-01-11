package rundeck

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"

	"golang.org/x/net/publicsuffix"
)

// ClientConfig represents a client configuration
type ClientConfig struct {
	BaseURL    string
	Token      string
	VerifySSL  bool
	Username   string
	Password   string
	AuthMethod string
	APIVersion string
	HTTPClient *http.Client
}

// Client represents a rundeck client
type Client struct {
	HTTPClient *http.Client
	Config     *ClientConfig
}

func defaultClientConfig() (*ClientConfig, error) {
	c, err := defaultHTTPClient(true)
	if err != nil {
		return nil, err
	}
	return &ClientConfig{
		VerifySSL:  true,
		APIVersion: MaxRundeckVersion,
		HTTPClient: c,
	}, nil
}

func defaultHTTPClient(verifySSL bool) (*http.Client, error) {
	tlsClientConfig := tls.Config{InsecureSkipVerify: !verifySSL}
	transport := &http.Transport{TLSClientConfig: &tlsClientConfig}
	client := &http.Client{Transport: transport}
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil, err
	}
	client.Jar = jar
	return client, nil
}

// NewClient creates a new client from the provided `ClientConfig`
func NewClient(config *ClientConfig) (*Client, error) {
	if config.HTTPClient == nil {
		c, err := defaultHTTPClient(config.VerifySSL)
		if err != nil {
			return nil, err
		}

		config.HTTPClient = c
	}
	rdClient := Client{
		HTTPClient: config.HTTPClient,
		Config:     config,
	}
	return &rdClient, nil
}

func clientConfigFrom(from string) (*ClientConfig, error) {
	config, configErr := defaultClientConfig()
	if configErr != nil {
		return nil, configErr
	}

	switch from {
	case "environment":
		if os.Getenv("RUNDECK_URL") == "" {
			return nil, fmt.Errorf("you must set the environment variable RUNDECK_URL")

		}
		if os.Getenv("RUNDECK_TOKEN") == "" {
			if os.Getenv("RUNDECK_USERNAME") == "" || os.Getenv("RUNDECK_PASSWORD") == "" {
				return nil, fmt.Errorf("you must set either RUNDECK_TOKEN or RUNDECK_USERNAME and RUNDECK_PASSWORD")
			}
			config.AuthMethod = basicAuthType
		} else {
			config.AuthMethod = tokenAuthType
		}
		if os.Getenv("RUNDECK_VERSION") != "" {
			ver := os.Getenv("RUNDECK_VERSION")
			intVer, intverErr := strconv.Atoi(ver)
			if intverErr != nil {
				return nil, intverErr
			}
			if intVer < minJSONSupportedAPIVersion {
				return nil, fmt.Errorf("minimum api version supported is %d", minJSONSupportedAPIVersion)
			}
			config.APIVersion = os.Getenv("RUNDECK_VERSION")
		}
		config.BaseURL = os.Getenv("RUNDECK_URL")
	}
	if config.AuthMethod == tokenAuthType {
		config.Token = os.Getenv("RUNDECK_TOKEN")
	} else {
		config.Username = os.Getenv("RUNDECK_USERNAME")
		config.Password = os.Getenv("RUNDECK_PASSWORD")
	}
	return config, nil
}

// NewClientFromEnv returns a new client from provided env vars
func NewClientFromEnv() (*Client, error) {
	config, configErr := clientConfigFrom("environment")
	if configErr != nil {
		return nil, configErr
	}
	return NewClient(config)
}

// NewBasicAuthClient returns a new client configured for basic auth using default settings
func NewBasicAuthClient(username, password, url string) (*Client, error) {
	config, configErr := defaultClientConfig()
	if configErr != nil {
		return nil, configErr
	}
	config.AuthMethod = basicAuthType
	config.Username = username
	config.Password = password
	config.BaseURL = url
	return NewClient(config)
}

// NewTokenAuthClient returns a new client configured for token auth using default settings
func NewTokenAuthClient(token, url string) (*Client, error) {
	config, configErr := defaultClientConfig()
	if configErr != nil {
		return nil, configErr
	}
	config.AuthMethod = tokenAuthType
	config.Token = token
	config.BaseURL = url
	return NewClient(config)
}
