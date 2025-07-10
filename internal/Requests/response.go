package Requests

type jsonMap struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []jsonMapResult `json:"results"`
}

type jsonMapResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var JsonMapData jsonMap
