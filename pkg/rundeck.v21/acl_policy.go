package rundeck

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// ACLPolicies represents ACL Policies
type ACLPolicies responses.ACLResponse

// GetACLPolicies gets the system ACL Policies
func (c *Client) GetACLPolicies() (*ACLPolicies, error) {
	data := &ACLPolicies{}
	res, err := c.httpGet("system/acl/", requestJSON())
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(res, &data)
	return data, jsonErr
}

// GetACLPolicy returns the named acl policy
func (c *Client) GetACLPolicy(policy string) (string, error) {
	url := fmt.Sprintf("system/acl/%s.aclpolicy", policy)
	res, err := c.httpGet(url, contentType("application/json"), accept("text/plain"))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// CreateACLPolicy creates a system acl policy
func (c *Client) CreateACLPolicy(name string, contents io.Reader) error {
	url := fmt.Sprintf("system/acl/%s.aclpolicy", name)
	res, err := c.httpPost(url, withBody(contents), accept("application/json"), contentType("application/yaml"), requestExpects(201))
	if err != nil {
		jsonError := &responses.ErrorResponse{}
		jsonErr := json.Unmarshal(res, jsonError)
		if jsonErr != nil {
			return jsonErr
		}
		return errors.New(jsonError.Message)
	}
	return nil
}

// UpdateACLPolicy creates a system acl policy
func (c *Client) UpdateACLPolicy(name string, contents io.Reader) error {
	url := fmt.Sprintf("system/acl/%s.aclpolicy", name)
	res, err := c.httpPut(url, withBody(contents), accept("application/json"), contentType("application/yaml"), requestExpects(200))
	if err != nil {
		jsonError := &responses.ErrorResponse{}
		jsonErr := json.Unmarshal(res, jsonError)
		if jsonErr != nil {
			return jsonErr
		}
		return errors.New(jsonError.Message)
	}
	return nil
}
