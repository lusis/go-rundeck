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

// NewClient creates a new client from the provided `ClientConfig`
func NewClient(config *ClientConfig) (*Client, error) {
	if config.HTTPClient == nil {
		tlsClientConfig := tls.Config{InsecureSkipVerify: !config.VerifySSL}
		transport := &http.Transport{TLSClientConfig: &tlsClientConfig}
		client := &http.Client{Transport: transport}
		config.HTTPClient = client
	}
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil, err
	}
	config.HTTPClient.Jar = jar
	client := Client{
		HTTPClient: config.HTTPClient,
		Config:     config,
	}
	return &client, nil
}

func clientConfigFrom(from string) (*ClientConfig, error) {
	config := ClientConfig{}

	switch from {
	case "environment":
		if os.Getenv("RUNDECK_TOKEN") == "" {
			if os.Getenv("RUNDECK_USERNAME") == "" || os.Getenv("RUNDECK_PASSWORD") == "" {
				return nil, fmt.Errorf("you must set either RUNDECK_TOKEN or RUNDECK_USERNAME and RUNDECK_PASSWORD")
			}
			config.AuthMethod = basicAuthType
		} else {
			config.AuthMethod = "token"
		}
		if os.Getenv("RUNDECK_VERSION") == "" {
			config.APIVersion = MaxRundeckVersion
		} else {
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

		if os.Getenv("RUNDECK_URL") == "" {
			return nil, fmt.Errorf("you must set the environment variable RUNDECK_URL")

		}
		config.BaseURL = os.Getenv("RUNDECK_URL")
	}
	if config.AuthMethod == "token" {
		config.Token = os.Getenv("RUNDECK_TOKEN")
	} else {
		config.Username = os.Getenv("RUNDECK_USERNAME")
		config.Password = os.Getenv("RUNDECK_PASSWORD")
	}
	return &config, nil
}

// NewClientFromEnv returns a new client from provided env vars
func NewClientFromEnv() (*Client, error) {
	config, configErr := clientConfigFrom("environment")
	if configErr != nil {
		return nil, configErr
	}
	return NewClient(config)
}
