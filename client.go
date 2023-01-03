package acg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "https://api.acloud.guru/v1/"

type Client struct {
	BaseUrl    string
	HTTPClient *http.Client
	Auth       HeaderStruct
}

type HeaderStruct struct {
	apiKey     string
	consumerId string
}

func NewClient(apiKey, consumerId *string) (*Client, error) {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseUrl:    HostURL,
		Auth: HeaderStruct{
			apiKey:     *apiKey,
			consumerId: *consumerId,
		},
	}, nil
}

func (c *Client) newRequest(requestMethod, requestPath string) (*http.Request, error) {
	req, err := http.NewRequest(requestMethod, fmt.Sprintf("%s/%s", c.BaseUrl, requestPath), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"x-api-key":     []string{c.Auth.apiKey},
		"x-consumer-id": []string{c.Auth.consumerId},
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
