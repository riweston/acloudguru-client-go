package main

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetUsers(i int) (*[]User, error) {
	path := fmt.Sprintf("users?page=%d&page-size=250", i)
	req, err := c.newRequest("GET", path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var getUsers []User
	err = json.Unmarshal(body, &getUsers)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &getUsers, nil
}

func (c *Client) GetUsersAll() (*[]User, error) {
	var getUsers []User
	i := 0
	for {
		i++
		tempUsers, err := c.GetUsers(i)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if len(*tempUsers) == 0 {
			break
		}

		getUsers = append(getUsers, *tempUsers...)
		if len(*tempUsers) < 250 {
			break
		}
	}

	return &getUsers, nil
}

func (c *Client) SetUserActivated(user *User, activate bool) (*Response, error) {
	var path string
	if activate {
		path = fmt.Sprintf("users/%s?action=activate", user.UserId)
	} else {
		path = fmt.Sprintf("users/%s?action=deactivate", user.UserId)
	}
	req, err := c.newRequest("POST", path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	postActivateUser := Response{}
	err = json.Unmarshal(body, &postActivateUser)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &postActivateUser, nil
}

func (c Client) GetUserFromEmail(email string) (*[]User, error) {
	path := fmt.Sprintf("users?email=%s", email)
	req, err := c.newRequest("GET", path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var getUserEmail []User
	err = json.Unmarshal(body, &getUserEmail)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &getUserEmail, nil
}
