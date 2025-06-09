package main

import (
	"fmt"
	dhanhq "github.com/Trade-with-Canvas/godhanhq"
)

const (
	accessToken  = "your_access_token_here"
	dhanClientId = "your_dhan_client_id_here"
)

func main() {

	dhanClient := dhanhq.New(false) // true enables debug logging

	// Set the access token and clientId for the Dhan client
	dhanClient.SetAccessToken(accessToken)
	dhanClient.SetDhanClientId(dhanClientId)

	// Get the margins for a given security and transaction type
	input := dhanhq.Margin{
		DhanClientId:    dhanClient.GetDhanClientId(),
		ExchangeSegment: dhanhq.ExchangeSegmentEquityNSE,
		TransactionType: dhanhq.TransactionTypeBuy,
		Quantity:        1,
		ProductType:     dhanhq.ProductTypeIntraday,
		SecurityId:      "NSE:RELIANCE",
		Price:           1447.00,
		TriggerPrice:    1600.00,
	}

	margins, err := dhanClient.CalculateMargins(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(margins)

	fundLimit, err := dhanClient.GetFundLimit()
	if err != nil {
		panic(err)
	}
	fmt.Println(fundLimit)

}
