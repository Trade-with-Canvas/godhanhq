/*
This is an example for generating and consuming consent
using this SDK for DhanHQ API, probably your first place to start with in the SDK aswell
Goodluck!
- h0i5
*/

package main

import (
	"fmt"
	dhanhq "github.com/tradewithcanvas/godhanhq"
)

const (
	// Partner ID and secret for DhanHQ API
	// Replace these with your actual partner ID and secret
	partnerId     = "your_partner_id"
	partnerSecret = "your_partner_secret"
)

func main() {

	// Create a new DhanHQ client with the partner ID
	dhanClient := dhanhq.New(true)

	// Set the partner secret
	dhanClient.SetPartnerId(partnerId)

	// Generate Consent
	consentResponse, err := dhanClient.GenerateConsent(partnerSecret)
	if err != nil {
		fmt.Println("Error generating consent:", err)
		return
	}

	fmt.Printf("Consent ID: %s, Consent Status: %s\n", consentResponse.ConsentId, consentResponse.ConsentStatus)

	// Consume Consent with the token ID returned from Consent Login (to be done on the web ui) at
	// https://auth.dhan.co/consent-login?consentId=[consentId]
	// GenerateConsentLoginURL is a helper function to generate the URL for consent login

	url := dhanClient.GenerateConsentLoginURL(consentResponse.ConsentId)
	fmt.Println("Consent login URL:", url)

	// Ask for the tokenID
	var tokenId string
	fmt.Print("Enter the token ID obtained from the consent login: ")
	fmt.Scanf("%s", &tokenId)

	if tokenId == "" {
		fmt.Println("Token ID cannot be empty")
	}

	consumeResponse, err := dhanClient.ConsumeConsent(tokenId, partnerSecret)
	if err != nil {
		fmt.Println("Error consuming consent:", err)
		return
	}

	// Print the details of the consumed consent
	fmt.Println("Consent consumed successfully:")
	fmt.Println("Client ID:", consumeResponse.DhanClientId)
	fmt.Println("Client Name:", consumeResponse.DhanClientName)
	fmt.Println("Client UCC:", consumeResponse.DhanClientUcc)
	fmt.Println("Given Power of Attorney:", consumeResponse.GivenPowerOfAttorney)
	fmt.Println("Access Token:", consumeResponse.AccessToken)
}
