package main

import (
	"fmt"
	dhanhq "github.com/tradewithcanvas/godhanhq"
)

const (
	accessToken  = "your_access_token_here"
	dhanClientId = "your_dhan_client_id_here"
)

func main() {
	dhanClient := dhanhq.New(true) // true enables debug logging
	// Set the access token for the Dhan client
	dhanClient.SetAccessToken(accessToken)
	dhanClient.SetDhanClientId(dhanClientId)

	// Get the LTP for a list of securities
	marketInput := dhanhq.MarketDataInput{
		"NSE_EQ":  {11536},
		"NSE_FNO": {49081, 49082},
	}
	ltpResponse, err := dhanClient.GetLTP(marketInput)
	if err != nil {
		panic(err)
	}

	for exchange, securities := range ltpResponse.Data {
		for securityId, quote := range securities {
			println("Exchange:", exchange, "Security ID:", securityId, "Last Price:", quote.LastPrice)
		}
	}

	// Get the OHLC for a list of securities
	ohlcResponse, err := dhanClient.GetOHLC(marketInput)
	if err != nil {
		panic(err)
	}
	for exchange, securities := range ohlcResponse.Data {
		for securityId, quote := range securities {
			println("Exchange:", exchange, "Security ID:", securityId, "Last Price:", quote.LastPrice,
				"Open:", quote.OHLC.Open, "Close:", quote.OHLC.Close,
				"High:", quote.OHLC.High, "Low:", quote.OHLC.Low)
		}
	}

	// Get the Market Depth for a list of securities
	marketDepthResponse, err := dhanClient.GetMarketDepth(marketInput)
	if err != nil {
		panic(err)
	}
	for exchange, securities := range marketDepthResponse.Data {
		for securityId, quote := range securities {
			fmt.Printf("Exchange: %s, Security ID: %s\n", exchange, securityId)
			fmt.Printf("Last Price: %.2f, Volume: %d\n", quote.LastPrice, quote.Volume)

			fmt.Println("Buy Orders:")
			for i, buy := range quote.Depth.Buy {
				fmt.Printf("  %d. Price: %.2f, Quantity: %d, Orders: %d\n",
					i+1, buy.Price, buy.Quantity, buy.Orders)
			}

			fmt.Println("Sell Orders:")
			for i, sell := range quote.Depth.Sell {
				fmt.Printf("  %d. Price: %.2f, Quantity: %d, Orders: %d\n",
					i+1, sell.Price, sell.Quantity, sell.Orders)
			}
		}
	}

}
