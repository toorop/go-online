package gonline

import (
	"net/http"
	"time"
)

const (
	apiEndpointV1 = "https://api.online.net/api/v1/"
)

// Online  Online.net API wrapper
type Online struct {
	client *http.Client
	apiKey string
}

// New return a new online.net API wrapper
// for new the returned err is useless but... who knows
func New(apiKey string) (*Online, error) {
	o := &Online{
		client: &http.Client{},
		apiKey: apiKey,
	}
	o.client.Timeout = 15 * time.Second
	return o, nil
}

// doRequest http.Client.Do with online authentification
func (o Online) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+o.apiKey)
	req.Header.Add("Accept", "application/json")
	return o.client.Do(req)
}

// GET HTTP get
func (o Online) get(ressource string) (*http.Response, error) {
	ressource = apiEndpointV1 + ressource
	req, err := http.NewRequest("GET", ressource, nil)
	if err != nil {
		return nil, err
	}
	return o.doRequest(req)
}
