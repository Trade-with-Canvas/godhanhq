package dhanhq

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// HTTPResponse represents the response from an HTTP request
type HTTPResponse struct {
	Body     []byte
	Response *http.Response
}

// httpClient is a client for making HTTP requests to the DhanHQ API
// TODO: Add custom loggers and timeouts to the httpClient
type httpClient struct {
	client *http.Client
	debug  bool // debug is used to enable/disable debug logging
}

// rURL stands for the relative URL for the API endpoints
// as seen in the constants declared in connect.go

// HTTPClient is an interface that defines methods for making HTTP requests
type HTTPClient interface {
	// Do is for sending form-data in POST/PUT and for query params in GET requests
	Do(method, rURL string, headers http.Header, params url.Values) (HTTPResponse, error)

	// DoRaw handles all the raw HTTP requests
	DoRaw(method, rURL string, reqBody []byte, headers http.Header) (HTTPResponse, error)

	// DoJSON is for sending JSON bodies in POST/PUT requests
	DoJSON(method, rURL string, queryParams url.Values, jsonBody interface{}, headers http.Header, respObj interface{}) (HTTPResponse, error)

	// GetClient returns the HTTP client
	GetClient() *httpClient
}

func NewHTTPClient(h *http.Client, debug bool) HTTPClient {
	return &httpClient{
		client: h,
		debug:  debug,
	}
}

// Do sends an HTTP request with the specified method, URL, headers, and parameters
// parameters are form-data in POST/PUT and query params in GET methods
func (c *httpClient) Do(method, rURL string, headers http.Header, params url.Values) (HTTPResponse, error) {
	var body []byte
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		if params != nil {
			// If params are provided, encode them as form data
			body = []byte(params.Encode())
		}
		if headers == nil {
			headers = http.Header{}
		}
		if headers.Get("Content-Type") == "" {
			headers.Set("Content-Type", "application/x-www-form-urlencoded")
		}

	} else {
		// For GET or other methods, params are treated as query parameters
		if len(params) > 0 {
			parsedURL, err := url.Parse(rURL)
			if err != nil {
				return HTTPResponse{}, err
			}
			parsedURL.RawQuery = params.Encode()
			rURL = parsedURL.String()
		}
	}

	// Call DoRaw with the updated stuff
	return c.DoRaw(method, rURL, body, headers)
}

// DoRaw sends an HTTP request with a raw body, typically in JSON format
func (c *httpClient) DoRaw(method, rURL string, reqBody []byte, headers http.Header) (HTTPResponse, error) {
	var resp HTTPResponse

	var bodyReader io.Reader
	if len(reqBody) > 0 {
		bodyReader = bytes.NewReader(reqBody)
	}

	if c.debug {
		log.Println("Request URL:", rURL)
	}
	req, err := http.NewRequest(method, rURL, bodyReader)

	if err != nil {
		return resp, err
	}
	// Set headers if provided
	if headers != nil {
		if c.debug {
			log.Println("Request Headers:")
			for key, values := range headers {
				for _, value := range values {
					log.Printf("%s: %s\n", key, value)
				}
			}
		}
		req.Header = headers
	}

	// Set the User-Agent header at last
	req.Header.Set("User-Agent", "DhanHQ Go SDK")

	// Log the request body if debug is enabled
	if c.debug {
		log.Println("Request Body:", string(reqBody))
	}

	// Log the headers if debug is enabled
	if c.debug {
		log.Println("Request Headers:", req.Header)
	}

	httpResponse, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(httpResponse.Body)

	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return resp, err
	}
	resp.Response = httpResponse
	resp.Body = data
	if c.debug {
		log.Println("Response Status:", httpResponse.Status)
		log.Println("Response Body:", string(data))
	}
	// Check if the response status code indicates an error
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		log.Println("Error Response Status:", httpResponse.Status)
		var errResp ErrorResponse
		if err := json.Unmarshal(data, &errResp); err != nil {
			return resp, err
		}
		if c.debug {
			log.Println("Error Response:", errResp)
		}
		return resp, err
	}
	return resp, nil
}

// DoJSON sends an HTTP request with a JSON body, typically in POST/PUT requests
func (c *httpClient) DoJSON(method, rURL string, queryParams url.Values, jsonBody interface{}, headers http.Header, respObj interface{}) (HTTPResponse, error) {
	var body []byte
	var err error
	if jsonBody != nil {
		// Convert jsonBody to []byte
		body, err = json.Marshal(jsonBody)
		if err != nil {
			return HTTPResponse{}, err
		}
		if headers == nil {
			headers = http.Header{}
		}
		headers.Set("Content-Type", "application/json")
	}
	// Add query parameters to the URL if provided
	if len(queryParams) > 0 {
		parsedURL, err := url.Parse(rURL)
		if err != nil {
			return HTTPResponse{}, err
		}
		parsedURL.RawQuery = queryParams.Encode()
		rURL = parsedURL.String()
	}

	// Call DoRaw with the updated stuff
	return c.DoRaw(method, rURL, body, headers)
}

// GetClient returns the HTTP client instance
func (c *httpClient) GetClient() *httpClient {
	return c
}
