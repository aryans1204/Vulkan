package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Client struct {
	authToken string
	HTTPClient *http.Client
	url string
	orgID string
}

func createClient(authToken string, orgID string) Client {
	fmt.Print("Instantiating client...")
	return Client{
		authToken: authToken,
		HTTPClient: &http.Client{},
		orgID: orgID,
		url: "https://api.openai.com/v1"
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	
}
