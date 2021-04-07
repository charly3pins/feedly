package feedly

import (
	"net/http"
)

// ClientConfig
type ClientConfig struct {
	BaseURL string
	Version string
	Token   string
}

// Client encapsulates the logic for connect to Feedly's API
type Client struct {
	// Config encapsulates the configuration need it for the Client
	Config ClientConfig
	// Client is the HTTP Client for make the calls
	Client http.Client
}
