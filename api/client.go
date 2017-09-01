package api

import (
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	Endpoint string
	Username string
	Password string
}

func (c *Client) GetRequest(url string) (*http.Response, error) {
	return c.Request("GET", url, 200)
}

func (c *Client) Request(verb, url string, code int) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != code {
		return resp, errors.New(fmt.Sprintf("Received %d, expecting %d status code while fetching %s", resp.StatusCode, code, url))
	}
	return resp, err
}
