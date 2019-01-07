package conekta

import (
	"log"
	"net/http"
	"time"
)

type requestor struct {
	httpClient *http.Client
}

func newRequestor() *requestor {

	// Create a new requestor with a 10 second request timeout limit.
	requestor := &requestor{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}

	return requestor
}

// SendRequest receives an interface, marshals it as json data and then sends it to the specified URI using the provided HTTP method (GET, POST, PUT, DELETE, etc).
func (requestor *requestor) sendRequest(method string, data interface{}, uri string, key string) (*http.Response, error) {

	// Create a new request.
	request, err := http.NewRequest(method, uri, nil)

	// If the request creation returned an error, return a nil conekta response and the error
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	// Prepare request headers
	setHeaders(request, key)

	// Send the request
	response, err := requestor.httpClient.Do(request)

	// If the request returned an error, return a nil conekta response and the error
	if err != nil {
		log.Fatal("HTTPClient Error: ", err)
		return nil, err
	}

	// Defer the body closure to ensure it after we finish reading.
	defer response.Body.Close()

	// Return the response.
	return response, nil
}

func setHeaders(request *http.Request, key string) {

	request.SetBasicAuth(key, "")
	request.Header.Add("Accept", "application/vnd.conekta-v2.0.0+json")
}
