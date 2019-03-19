package conekta

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
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

// Response ...
type Response struct {
}

// Global endpoint map.
var endpoints = map[string]string{
	"Order": "/orders",
}

// NewClient sets up a new Client structure. The API Key must be provided as a
// paramater. The BaseURI, Locale and Version fields are set by default, but
// can be modified directly.
//
// The client also has an HTTP client to make the corresponding API calls. It
// has a default timeout of 10 seconds, but can be changed directly as with any
// other HTTP client from the 'net/http' library.
// (https://golang.org/pkg/net/http/#Client)
func NewClient(key string) (*Client, error) {

	client := &Client{
		BaseURI: "https://api.conekta.io",
		Locale:  "es",
		Key:     key,
		Version: "2.0.0",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		Debug: true,
	}

	return client, nil
}

// Create ...
func (client *Client) Create(data interface{}) (*http.Response, error) {

	// Reflect data type.
	dataType := reflect.TypeOf(data).Name()

	// Get the corresponding endpoint from the map.
	endpoint := endpoints[dataType]

	// Marshal our data to use as the request body.
	body, err := json.Marshal(data)

	// Return error if any.
	if err != nil {
		return nil, err
	}

	// // Send the request.
	request, err := client.createHTTPRequest("POST", client.BaseURI+endpoint, body)

	client.debug(request)

	// Return error if any.
	if err != nil {
		return nil, err
	}

	// // Retrieve response.
	response, err := client.HTTPClient.Do(request)

	// Return error if any.
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 202 {
		return nil, fmt.Errorf("%#v", response)
	}

	// Print the response. 
	client.debug(response)

	return response, nil
}

// Update ...
func Update(data interface{}) {

	// Reflect data type.
	// Get the corresponding endpoint from the map.
	// Send the request.
	// Retrieve response.
	// Validate if an error occurred.
	// Return error, if any.
	// Return a response.

}

// Delete ...
func Delete(data interface{}) {

	// Reflect data type.
	// Get the corresponding endpoint from the map.
	// Send the request.
	// Retrieve response.
	// Validate if an error occurred.
	// Return error, if any.
	// Return a response.
}

func (client *Client) debug(data interface{}) {
	if client.Debug {

		log.Printf("%#v", data)
	}
}

func (client *Client) createHTTPRequest(method, url string, body []byte) (*http.Request, error) {

	request, err := http.NewRequest(method, url, nil)

	// Convert key to Base64 encoded string
	key := base64.StdEncoding.EncodeToString([]byte(client.Key))

	// Create the user agent.
	userAgent := `{"binding_version":"` + client.Version + `","lang":"go","lang_version":"go1.11.4","publisher":"holasoyeddy","uname":"uname"}`

	request.Header.Add("Authorization", " Basic "+key+":")
	request.Header.Add("Accept-Language", client.Locale)
	request.Header.Add("Accept", "application/vnd.conekta-v"+client.Version+"json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Conekta/v1 DotNetBindings10/Conekta::"+client.Version)
	request.Header.Add("X-Conekta-Client-User-Agent", userAgent)

	if err != nil {
		return nil, err
	}

	return request, nil
}
