package rundeck

import (
	"fmt"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

// UploadKey stores keys on the rundeck server
// http://rundeck.org/docs/api/index.html#upload-keys
func (c *Client) UploadKey() error {
	if err := c.checkRequiredAPIVersion(responses.GenericVersionedResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// ListKeys lists key resources
// http://rundeck.org/docs/api/index.html#list-keys
func (c *Client) ListKeys() error {
	if err := c.checkRequiredAPIVersion(responses.ListKeysResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetKeyMetaData returns the metadata about a stored key
// http://rundeck.org/docs/api/index.html#get-key-metadata
func (c *Client) GetKeyMetaData() error {
	if err := c.checkRequiredAPIVersion(responses.KeyMetaResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetKeyContents provides the public key content
// http://rundeck.org/docs/api/index.html#get-key-contents
func (c *Client) GetKeyContents() error {
	if err := c.checkRequiredAPIVersion(responses.GenericVersionedResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DeleteKey deletes a key
// http://rundeck.org/docs/api/index.html#delete-keys
func (c *Client) DeleteKey() error {
	if err := c.checkRequiredAPIVersion(responses.GenericVersionedResponse{}); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
