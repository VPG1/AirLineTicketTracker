package aviasales_client

import (
	"AirLineTicketTracker/internal/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	token  string
	host   string
	path   string
	client http.Client
}

func New(host string, path string, token string) *Client {
	return &Client{
		token:  token,
		host:   host,
		path:   path,
		client: http.Client{},
	}
}

func (c *Client) GetFlightInfo(flight entities.Flight) (*entities.Flight, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   c.path,
	}

	q := url.Values{}
	q.Add("origin", flight.OriginIATA)
	q.Add("destination", flight.DestinationIATA)
	q.Add("currency", "usd")
	q.Add("period_type", "year")
	q.Add("page", "1")
	q.Add("limit", "1")
	q.Add("sorting", "price")
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

	MapRespToFlight(responseBody, &flight)
	newFlight := flight
	return &newFlight, nil
}
