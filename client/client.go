package client

import "net/http"

var (
	DefaultAPIURL = "https://pokeapi.co"
)

type Client struct {
	apiURL     string
	httpClient *http.Client
}

type Option func(c *Client)

func NewClient(opts ...Option) *Client {
	client := &Client{
		apiURL:     DefaultAPIURL,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithAPIURL(url string) Option {
	return func(c *Client) {
		c.apiURL = url
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}
