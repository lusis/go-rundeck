package rundeck

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

type ClientConfig struct {
	BaseURL   string
	Token     string
	VerifySSL bool
}

type RundeckClient struct {
	Client *http.Client
	Config *ClientConfig
}

func NewClient(config *ClientConfig) (c RundeckClient) {
	verifySSL := func() bool {
		if config.VerifySSL != true {
			return false
		} else {
			return true
		}
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: verifySSL()},
	}
	client := &http.Client{Transport: tr}
	return RundeckClient{Client: client, Config: config}
}

func clientConfigFrom(from string) (c *ClientConfig) {
	switch from {
	case "environment":
		if os.Getenv("RUNDECK_TOKEN") == "" || os.Getenv("RUNDECK_URL") == "" {
			fmt.Printf("You must set the environment variables  RUNDECK_URL and RUNDECK_TOKEN\n")
			os.Exit(1)
		}
	}
	config := ClientConfig{
		BaseURL: os.Getenv("RUNDECK_URL"),
		Token:   os.Getenv("RUNDECK_TOKEN"),
	}
	return &config
}

func NewClientFromEnv() (c RundeckClient) {
	config := clientConfigFrom("environment")
	client := NewClient(config)
	return client
}
