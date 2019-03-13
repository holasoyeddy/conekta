package conekta

import (
	"net/http"
)

// Client is the structure used to communicate with the Conekta Web API.
type Client struct {
	BaseURI    string
	Locale     string
	Key        string
	Version    string
	HTTPClient *http.Client
	Debug      bool
}

// NewClient sets up a new Client structure. The API Key must be provided as a
// paramater. The BaseURI, Locale and Version fields are set by default, but
// can be modified directly.
//
// The client also has an HTTP client to make the corresponding API calls. It
// has a default timeout of 10 seconds, but can be changed directly as with any
// other HTTP client from the 'net/http' library.
// (https://golang.org/pkg/net/http/#Client)
func NewClient(key string) *Client {

	client := &Client{
		BaseURI:    "https://api.conekta.io/",
		Locale:     "es",
		Key:        key,
		Version:    "2.0.0",
		HTTPClient: newHTTPClient(),
		Debug:      false,
	}

	return client
}
