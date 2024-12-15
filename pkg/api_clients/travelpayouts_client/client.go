package iata_code_definition_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	host   string
	client http.Client
}

func New(host string) *Client {
	return &Client{
		host:   host,
		client: http.Client{},
	}
}

func (c Client) GetIATACode(searchPhrase string) (*Response, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   "widgets_suggest_params",
	}

	q := url.Values{}
	q.Add("q", searchPhrase)

	fmt.Println(q.Encode())

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

	if responseBody.Destination.IATA == "" || responseBody.Origin.IATA == "" {
		return nil, IncorrectResponse
	}

	return responseBody, nil
}
