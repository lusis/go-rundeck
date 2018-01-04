package rundeck

import (
	"bytes"
	"encoding/json"
	"errors"

	multierror "github.com/hashicorp/go-multierror"
	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// User represents a user in rundeck
type User responses.UserInfoResponse

// GetUsers returns all rundeck users
func (c *Client) GetUsers() ([]*User, error) {
	var users []*User
	res, err := c.httpGet("user/list", requestJSON(), requestExpects(200))
	if err != nil {
		return users, err
	}
	if jsonErr := json.Unmarshal(res, &users); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return users, nil
}

// GetCurrentUserInfo returns information about the current user
func (c *Client) GetCurrentUserInfo() (*User, error) {
	user := &User{}
	res, err := c.httpGet("user/info", requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &user); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return user, nil
}

// GetUserInfo returns information about the named user - requires admin privileges
func (c *Client) GetUserInfo(login string) (*User, error) {
	user := &User{}
	res, err := c.httpGet("user/info/"+login, requestJSON(), requestExpects(200))
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, &user); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return user, nil
}

// UpdateUserInfo updates a user
func (c *Client) UpdateUserInfo(u *User) (*User, error) {
	currentUser, currentUserErr := c.GetCurrentUserInfo()
	if currentUserErr != nil {
		return nil, currentUserErr
	}
	if u.Login == "nil" {
		return nil, errors.New("must provide login and at least one field to update")
	}
	updatePath := "user/info"
	if u.Login != currentUser.Login {
		// we're not updating ourself so we need to append the login to the path
		updatePath = updatePath + "/" + u.Login
	}
	newUser := &User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
	postData, postDataErr := json.Marshal(newUser)
	if postDataErr != nil {
		return nil, postDataErr
	}
	res, resErr := c.httpPost(updatePath, withBody(bytes.NewReader(postData)), requestJSON())
	if resErr != nil {
		return nil, resErr
	}
	resUser := &User{}
	if jsonErr := json.Unmarshal(res, &resUser); jsonErr != nil {
		return nil, &UnmarshalError{msg: multierror.Append(errDecoding, jsonErr).Error()}
	}
	return resUser, nil
}
