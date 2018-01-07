package rundeck

import (
	"fmt"
	"strconv"
)

func (c *Client) hasRequiredAPIVersion(min, max int) (bool, error) {
	reqVersion, err := strconv.Atoi(c.Config.APIVersion)
	if err != nil {
		return false, err
	}
	if reqVersion >= min && reqVersion <= max {
		return true, nil
	}
	return false, fmt.Errorf("Requested API version (%d) does not meet the requirements for this api call (min: %d, max: %d)",
		reqVersion, min, max)
}
