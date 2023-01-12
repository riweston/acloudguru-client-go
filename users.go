package acg

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) GetUsersByPage(i int) (user *[]User, error error) {
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

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func (c *Client) GetUsersAll() (user *[]User, error error) {
	var getUsers []User
	i := 0
	for {
		i++
		tempUsers, err := c.GetUsersByPage(i)
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

func (c *Client) SetUserActivated(user *User, activate bool) (response *Response, error error) {
	var path string
	if activate {
		path = fmt.Sprintf("users/%s?action=activate", user.UserId)
		path = url.PathEscape(path)
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

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response.ErrorResponse.ErrorMessage != "" {
		return nil, fmt.Errorf("%s", response.ErrorResponse.ErrorMessage)
	}

	return response, nil
}

func (c *Client) GetUserFromEmail(email string) (user *User, error error) {
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

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}
