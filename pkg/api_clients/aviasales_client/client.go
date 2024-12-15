package aviasales_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	token  string
	host   string
	client http.Client
}

func New(host string, token string) *Client {
	return &Client{
		token:  token,
		host:   host,
		client: http.Client{},
	}
}

func (c *Client) GetFlightInfo(origin string, destination string) (*Response, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   "v1/prices/cheap",
	}

	q := url.Values{}
	q.Add("origin", origin)
	q.Add("destination", destination)
	q.Add("token", c.token)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	responseBody := &Response{}
	if err := json.Unmarshal(body, responseBody); err != nil {
		return nil, fmt.Errorf("error when unmarshal: %s", err)
	}

	if responseBody.Success == false {
		return nil, FlightNotFound
	}
	
	return responseBody, nil
}
