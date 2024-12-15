package iata_code_definition_api

type Response struct {
	Origin struct {
		IATA string `json:"iata"`
		Name string `json:"name"`
	} `json:"origin"`
	Destination struct {
		IATA string `json:"iata"`
		Name string `json:"name"`
	} `json:"destination"`
}
