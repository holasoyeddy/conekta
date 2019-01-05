package conekta

// Client structure contains the base information required for the library to operate. Its methods will allow communication to the Conekta API server via HTTP requests.
type Client struct {
	BaseURI string
	Locale  string
	Key     string
	Version string
}

// NewClient sets up a new client with the API key provided. The other values are set by default, but can be changed using the appropriate setters.
func NewClient(key string) *Client {
	return &Client{BaseURI: "https://api.conekta.io/", Locale: "es", Key: key, Version: "2.0.0"}
}
