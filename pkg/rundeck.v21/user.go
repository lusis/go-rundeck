package rundeck

import (
	"bytes"
	"encoding/json"
	"errors"
)

// User represents a user in rundeck
type User struct {
	Login     string `json:"login,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

// GetUsers returns all rundeck users
func (c *Client) GetUsers() ([]*User, error) {
	var users []*User
	res, err := c.httpGet("user/list", requestJSON())
	if err != nil {
		return users, err
	}
	jsonErr := json.Unmarshal(res, &users)
	return users, jsonErr
}

// GetCurrentUserInfo returns information about the current user
func (c *Client) GetCurrentUserInfo() (*User, error) {
	user := &User{}
	res, err := c.httpGet("user/info", requestJSON())
	if err != nil {
		return user, err
	}
	jsonErr := json.Unmarshal(res, &user)
	return user, jsonErr
}

// GetUserInfo returns information about the named user - requires admin privileges
func (c *Client) GetUserInfo(login string) (*User, error) {
	user := &User{}
	res, err := c.httpGet("user/info/"+login, requestJSON())
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(res, &user)
	return user, jsonErr
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
	jsonErr := json.Unmarshal(res, &resUser)
	return resUser, jsonErr
}
