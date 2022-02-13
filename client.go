package acloudguru

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HeaderStruct struct {
	xApiKey     string `json:"x-api-key"`
	xConsumerId string `json:"x-consumer-id"`
}

const HostURL string = "https://api.acloud.guru/v1/"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       HeaderStruct
}

func NewClient(host, xApiKey, xConsumerId *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if xApiKey == nil || xConsumerId == nil {
		return &c, nil
	}

	c.Auth = HeaderStruct{
		xApiKey:     *xApiKey,
		xConsumerId: *xConsumerId,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header = http.Header{
		"Host":          []string{c.HostURL},
		"Content-Type":  []string{"application/json"},
		"x-api-key":     []string{c.Auth.xApiKey},
		"x-consumer-id": []string{c.Auth.xConsumerId},
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
