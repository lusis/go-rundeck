package rundeck

import (
	"fmt"
	"strconv"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses"
)

func (c *Client) checkRequiredAPIVersion(r responses.VersionedResponse) error {
	reqVersion, err := strconv.Atoi(c.Config.APIVersion)
	if err != nil {
		return err
	}
	min := responses.GetMinVersionFor(r)
	max := responses.GetMaxVersionFor(r)
	if reqVersion >= min && reqVersion <= max {
		return nil
	}
	return fmt.Errorf("Requested API version (%d) does not meet the requirements for this api call (min: %d, max: %d)",
		reqVersion, min, max)
}
