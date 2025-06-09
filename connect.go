/*
License or smth idk
*/

package dhanhq

import (
	"net/http"
)

// Information about the library and some useful constants
const (
	name    string = "dhanhq"
	version string = "0.1.0"
	baseURI string = "https://api.dhan.co/v2"
	authURI string = "https://auth.dhan.co"
)

// Client represents the interface for DhanHQ API client
type Client struct {
	dhanClientId string
	accessToken  string
	baseURI      string
	authURI      string
	partnerId    string

	// HTTP client for making requests
	httpClient HTTPClient
}

// API endpoints for DhanHQ
const (
	// Partner endpoints for auth

	URIPartnerGenerateConsent = "/partner/generate-consent"
	URIPartnerConsentLogin    = "/partner/consent-login"
	URIPartnerConsumeConsent  = "/partner/consume-consent"

	// Data endpoints

	URIMarketfeedLTP   = "/marketfeed/ltp"
	URIMarketfeedOHLC  = "/marketfeed/ohlc"
	URIMarketfeedQuote = "/marketfeed/quote"

	// Historical data endpoints

	URIChartsHistorical = "/charts/historical"
	URIChartsIntraday   = "/charts/intraday"

	// Option chains endpoints

	URIOptionchain           = "/optionchain"
	URIOptionchainExpiryList = "/optionchain/expirylist"

	// Portfolio endpoints

	URIHoldings        = "/holdings"
	URIPositions       = "/positions"
	URIPositionConvert = "/positions/convert"

	// Fund related endpoints

	URIMarginCalculator = "/margincalculator"
	URIFundLimit        = "/fundlimit"

	// Profile endpoints

	URIProfile = "/profile"
)

type ErrorMessage struct {
	ErrorType    string `json:"errorType"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// New creates a new DhanHQ API client with the provided parameters.
func New(partnerId string, debug bool) *Client {
	client := &Client{
		baseURI:   baseURI,
		authURI:   authURI,
		partnerId: partnerId,
	}
	// Initialize the HTTP client
	client.httpClient = NewHTTPClient(
		&http.Client{},
		debug, // Pass the debug flag to the HTTP client
	)

	return client
}

func (c *Client) GetBaseURI() string {
	return c.baseURI
}
func (c *Client) GetAuthURI() string {
	return c.authURI
}
func (c *Client) GetDhanClientId() string {
	return c.dhanClientId
}
func (c *Client) GetAccessToken() string {
	return c.accessToken
}
func (c *Client) GetPartnerId() string {
	return c.partnerId
}

// Setters for Client fields

func (c *Client) SetBaseURI(baseURI string) {
	c.baseURI = baseURI
}
func (c *Client) SetAuthURI(authURI string) {
	c.authURI = authURI
}
func (c *Client) SetDhanClientId(dhanClientId string) {
	c.dhanClientId = dhanClientId
}
func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}
func (c *Client) SetPartnerId(partnerId string) {
	c.partnerId = partnerId
}
func (c *Client) SetHTTPClient(h *http.Client, debug bool) {
	// Implement a new HTTPClient interface that wraps the standard http.Client
	c.httpClient = NewHTTPClient(
		h,
		debug, // Pass the debug flag to the HTTP client
	)
}
func (c *Client) GetHTTPClient() HTTPClient {
	if c.httpClient == nil {
		c.httpClient = NewHTTPClient(
			&http.Client{},
			true, // Default to no debug logging
		)
	}
	return c.httpClient
}
