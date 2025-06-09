package main

import (
	"fmt"
	dhanhq "github.com/tradewithcanvas/godhanhq"
)

const (
	accessToken = "access_token_here"
)

func main() {
	dhanClient := dhanhq.New(true) // true enables debug logging
	// Set the access token for the Dhan client
	dhanClient.SetAccessToken(accessToken)

	// Get the Positions for a client
	positions, err := dhanClient.GetPositions()
	if err != nil {
		panic(err)
	}
	fmt.Println(positions)

	// Get the Holdings for a client
	holdings, err := dhanClient.GetHoldings()
	if err != nil {
		panic(err)
	}
	fmt.Println(holdings)
}
