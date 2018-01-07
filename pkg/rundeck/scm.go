package rundeck

import "fmt"

// ListSCMPlugins list the available plugins for the specified integration
// http://rundeck.org/docs/api/index.html#list-scm-plugins
func (c *Client) ListSCMPlugins() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetSCMPluginInputFields List the input fields for a specific plugin.
// http://rundeck.org/docs/api/index.html#get-scm-plugin-input-fields
func (c *Client) GetSCMPluginInputFields() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// SetupSCMPluginForProject configures and enables a plugin for a project
// http://rundeck.org/docs/api/index.html#setup-scm-plugin-for-a-project
func (c *Client) SetupSCMPluginForProject() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// EnableSCMPluginForProject enables a plugin for a project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
func (c *Client) EnableSCMPluginForProject() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// DisableSCMPluginForProject disables a plugin for a project
// http://rundeck.org/docs/api/index.html#enable-scm-plugin-for-a-project
func (c *Client) DisableSCMPluginForProject() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectSCMStatus Get the SCM plugin status and available actions for the project.
// http://rundeck.org/docs/api/index.html#get-project-scm-status
func (c *Client) GetProjectSCMStatus() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectSCMConfig Get the configuration properties for the current plugin.
// http://rundeck.org/docs/api/index.html#get-project-scm-config
func (c *Client) GetProjectSCMConfig() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetProjectSCMActionInputFields Get the input fields and selectable items for a specific action.
// http://rundeck.org/docs/api/index.html#get-project-scm-action-input-fields
func (c *Client) GetProjectSCMActionInputFields() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// PerformProjectSCMAction Perform the action for the SCM integration plugin, with a set of input parameters, selected Jobs, or Items, or Items to delete.
// http://rundeck.org/docs/api/index.html#perform-project-scm-action
func (c *Client) PerformProjectSCMAction() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetJobSCMStatus gets a job's scm status
// http://rundeck.org/docs/api/index.html#get-job-scm-status
func (c *Client) GetJobSCMStatus() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetJobSCMDiff Retrieve the file diff for the Job, if there are changes for the integration.
// http://rundeck.org/docs/api/index.html#get-job-scm-diff
func (c *Client) GetJobSCMDiff() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// GetJobSCMActionInputFields Get the input fields and selectable items for a specific action for a job.
// http://rundeck.org/docs/api/index.html#get-project-scm-action-input-fields
func (c *Client) GetJobSCMActionInputFields() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}

// PerformJobSCMAction Perform the action for the SCM integration plugin, with a set of input parameters, selected Jobs, or Items, or Items to delete.
// http://rundeck.org/docs/api/index.html#perform-project-scm-action
func (c *Client) PerformJobSCMAction() error {
	if _, err := c.hasRequiredAPIVersion(15, maxRundeckVersionInt); err != nil {
		return err
	}
	return fmt.Errorf("not yet implemented")
}
