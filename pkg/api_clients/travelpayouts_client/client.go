package iata_code_definition_api

import (
	"AirLineTicketTracker/internal/entities"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func (c Client) GetIATACodes(searchPhrase string) (*entities.Flight, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   "widgets_suggest_params",
	}

	q := url.Values{}
	q.Add("q", searchPhrase)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Println("NewRequest: ", err)
		return nil, err
	}

	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println("Client.Do: ", err)
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

	return getFlightFromResp(responseBody), nil
}
