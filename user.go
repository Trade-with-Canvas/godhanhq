package dhanhq

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type GenerateConsentResponse struct {
	ConsentId     string `json:"consentId"`
	ConsentStatus string `json:"consentStatus"`
}

type ConsumeConsentResponse struct {
	DhanClientId         string `json:"dhanClientId"`
	DhanClientName       string `json:"dhanClientName"`
	DhanClientUcc        string `json:"dhanClientUcc"`
	GivenPowerOfAttorney bool   `json:"givenPowerOfAttorney"`
	AccessToken          string `json:"accessToken"`
	ExpiryTime           string `json:"expiryTime"`
}

func (c *Client) GenerateConsent(partnerSecret string) (GenerateConsentResponse, error) {
	// This contains the logic to generate the consent from the partner_id and
	// partner_secret using the DhanHQ API
	// Add the partner_secret and partner_id to the headers
	consentHeaders := http.Header{
		"partner_secret": {partnerSecret},
		"partner_id":     {c.partnerId},
	}

	resp, err := c.httpClient.Do(http.MethodGet, c.authURI+URIPartnerGenerateConsent, consentHeaders, nil)
	if err != nil {
		return GenerateConsentResponse{}, err
	}

	var consentResponse GenerateConsentResponse
	if err := json.Unmarshal(resp.Body, &consentResponse); err != nil {
		return GenerateConsentResponse{}, err
	}

	return consentResponse, nil
}

func (c *Client) ConsumeConsent(tokenId string, partnerSecret string) (ConsumeConsentResponse, error) {
	// This contains the logic to consume the consent and thus return
	// valid accessToken and clientId into type ConsumeConsentResponse for the client

	consumeHeaders := http.Header{
		"partner_secret": {partnerSecret},
		"partner_id":     {c.partnerId},
	}

	// Create the params for the request
	var consumeParams url.Values
	consumeParams = make(url.Values)

	// Add the tokenId to the params
	consumeParams["tokenId"] = []string{tokenId}

	resp, err := c.httpClient.Do(http.MethodPost, c.authURI+URIPartnerConsumeConsent, consumeHeaders, consumeParams)
	if err != nil {
		return ConsumeConsentResponse{}, err
	}
	var consumeResponse ConsumeConsentResponse
	if err := json.Unmarshal(resp.Body, &consumeResponse); err != nil {
		return ConsumeConsentResponse{}, err
	}

	return consumeResponse, nil
}

func (c *Client) GenerateConsentLoginURL(consentId string) string {
	// This returns the consent login URL for the user to login and get the tokenId
	// which is then used to consume the consent
	return c.authURI + URIPartnerConsentLogin + "?consentId=" + consentId
}
