package opensea

import "net/http"

// Config is the global configuration for OpenSea client.
type Config struct {
	DefaultURL string `json:"openSeaDefaultURL" default:"https://api.opensea.io/api/v1/"`
	APIKey     string `json:"api-key"`
}

// Client implementation of the OpenSea API.
type Client struct {
	http   http.Client
	config Config
}

// New is a constructor for OpenSea API client.
func New(config Config) *Client {
	return &Client{config: config}
}
