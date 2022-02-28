package main

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetSubscription() (*Subscription, error) {
	req, err := c.newRequest("GET", "subscription")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	getSubscription := Subscription{}
	err = json.Unmarshal(body, &getSubscription)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &getSubscription, nil
}
