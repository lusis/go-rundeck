package rundeck

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// ACLPolicies represents ACL Policies
type ACLPolicies responses.ACLResponse

// ListSystemACLPolicies gets the system ACL Policies
// http://rundeck.org/docs/api/index.html#list-system-acl-policies
func (c *Client) ListSystemACLPolicies() (*ACLPolicies, error) {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return nil, err
	}
	data := &ACLPolicies{}
	res, err := c.httpGet("system/acl/", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &data); jsonErr != nil {
		return nil, jsonErr
	}
	return data, nil
}

// GetSystemACLPolicy returns the named acl policy
// http://rundeck.org/docs/api/index.html#get-an-acl-policy
func (c *Client) GetSystemACLPolicy(policy string) ([]byte, error) {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("system/acl/%s.aclpolicy", policy)
	res, err := c.httpGet(url, accept("application/yaml"), requestExpects(200))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateSystemACLPolicy creates a system acl policy
// http://rundeck.org/docs/api/index.html#create-an-acl-policy
func (c *Client) CreateSystemACLPolicy(name string, contents io.Reader) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	url := fmt.Sprintf("system/acl/%s.aclpolicy", name)
	res, err := c.httpPost(url, withBody(contents),
		accept("application/json"),
		contentType("application/yaml"),
		requestExpects(201),
		requestExpects(400))
	if err != nil {
		return err
	}
	// okay we have a body in the response
	// we should see if it's a validation error response
	jsonError := &responses.FailedACLValidationResponse{}
	jsonErr := json.Unmarshal(res, jsonError)
	if jsonErr != nil {
		// It's not a validation response
		return nil
	}
	var finalErr error
	for _, v := range jsonError.Policies {
		line := fmt.Sprintf("%s: %s", v.Policy, strings.Join(v.Errors, ","))
		finalErr = multierror.Append(finalErr, fmt.Errorf("%s", line))
	}
	if finalErr != nil {
		return &PolicyValidationError{msg: finalErr.Error()}
	}
	return nil
}

// UpdateSystemACLPolicy creates a system acl policy
// http://rundeck.org/docs/api/index.html#update-an-acl-policy
func (c *Client) UpdateSystemACLPolicy(name string, contents io.Reader) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	url := fmt.Sprintf("system/acl/%s.aclpolicy", name)
	res, err := c.httpPut(url, withBody(contents), accept("application/json"), contentType("application/yaml"), requestExpects(201), requestExpects(400))
	if err != nil {
		return err
	}
	jsonError := &responses.FailedACLValidationResponse{}
	jsonErr := json.Unmarshal(res, jsonError)
	if jsonErr != nil {
		// just return the original error
		return nil
	}
	var finalErr error
	for _, v := range jsonError.Policies {
		line := fmt.Sprintf("%s: %s", v.Policy, strings.Join(v.Errors, ","))
		finalErr = multierror.Append(finalErr, fmt.Errorf("%s", line))
	}
	return &PolicyValidationError{msg: finalErr.Error()}
}

// DeleteSystemACLPolicy deletes a system ACL Policy
// http://rundeck.org/docs/api/index.html#delete-an-acl-policy
func (c *Client) DeleteSystemACLPolicy(name string) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	return c.httpDelete("system/acl/"+name+".aclpolicy", requestJSON(), requestExpects(204))
}

// ListProjectACLPolicies gets a project ACL Policies
// http://rundeck.org/docs/api/index.html#list-project-acl-policies
func (c *Client) ListProjectACLPolicies(name string) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectACLPolicy gets a project ACL Policy
// http://rundeck.org/docs/api/index.html#get-a-project-acl-policy
func (c *Client) GetProjectACLPolicy(name string) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DeleteProjectACLPolicy deletes a project ACL Policy
// http://rundeck.org/docs/api/index.html#delete-a-project-acl-policy
func (c *Client) DeleteProjectACLPolicy(name string) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// CreateProjectACLPolicy creates a project ACL Policy
// http://rundeck.org/docs/api/index.html#create-a-project-acl-policy
func (c *Client) CreateProjectACLPolicy(name string) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// UpdateProjectACLPolicy updates a project ACL Policy
// http://rundeck.org/docs/api/index.html#update-a-project-acl-policy
func (c *Client) UpdateProjectACLPolicy(name string) error {
	if err := c.checkRequiredAPIVersion(responses.ACLResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
