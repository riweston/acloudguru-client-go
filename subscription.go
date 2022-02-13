package acloudguru

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetSubscription() (*Subscription, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/subscription", c.HostURL), nil)
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
