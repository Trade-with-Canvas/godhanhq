package dhanhq

import (
	"encoding/json"
	"net/http"
)

// MarketDataInput represents the input for market data requests
// which is a JSON object with keys as exchange segments and values
// as arrays of integers representing security IDs
type MarketDataInput map[string][]int

type LTPResponse struct {
	Status string                         `json:"status"`
	Data   map[string]map[string]LTPQuote `json:"data"`
}

type LTPQuote struct {
	LastPrice float64 `json:"last_price"`
}

type OHLCResponse struct {
	Status string                          `json:"status"`
	Data   map[string]map[string]OHLCQuote `json:"data"`
}
type OHLCQuote struct {
	LastPrice float64 `json:"last_price"`
	OHLC      OHLC    `json:"ohlc"`
}

type OHLC struct {
	Open  float64 `json:"open"`
	Close float64 `json:"close"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
}

type MarketDepthResponse struct {
	Status string                                 `json:"status"`
	Data   map[string]map[string]MarketDepthQuote `json:"data"`
}

type MarketDepthItem struct {
	Quantity int32   `json:"quantity"`
	Orders   int32   `json:"orders"`
	Price    float64 `json:"price"`
}

type MarketDepth struct {
	Buy  []MarketDepthItem `json:"buy"`
	Sell []MarketDepthItem `json:"sell"`
}

type MarketDepthQuote struct {
	AveragePrice      float64     `json:"average_price"`
	BuyQuantity       int32       `json:"buy_quantity"`
	SellQuantity      int32       `json:"sell_quantity"`
	Depth             MarketDepth `json:"depth"`
	LastPrice         float64     `json:"last_price"`
	LastQuantity      int32       `json:"last_quantity"`
	LastTradeTime     string      `json:"last_trade_time"`
	LowerCircuitLimit float64     `json:"lower_circuit_limit"`
	UpperCircuitLimit float64     `json:"upper_circuit_limit"`
	NetChange         float64     `json:"net_change"`
	Volume            int32       `json:"volume"`
	OI                int32       `json:"oi"`
	OIDayHigh         int32       `json:"oi_day_high"`
	OIDayLow          int32       `json:"oi_day_low"`
	OHLC              OHLC        `json:"ohlc"`
}

func (c *Client) GetLTP(input MarketDataInput) (LTPResponse, error) {
	headers := http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"access-token": {c.accessToken},
		"client-id":    {c.GetDhanClientId()},
	}
	resp, err := c.httpClient.DoJSON(http.MethodPost, c.baseURI+URIMarketfeedLTP, nil, input, headers, &LTPResponse{})
	if err != nil {
		return LTPResponse{}, err
	}
	if resp.Response.StatusCode != http.StatusOK {
		return LTPResponse{}, nil
	}
	var ltpResponse LTPResponse
	if err := json.Unmarshal(resp.Body, &ltpResponse); err != nil {
		return LTPResponse{}, err
	}

	return ltpResponse, nil
}

func (c *Client) GetOHLC(input MarketDataInput) (OHLCResponse, error) {
	headers := http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"access-token": {c.accessToken},
		"client-id":    {c.GetDhanClientId()},
	}
	resp, err := c.httpClient.DoJSON(http.MethodPost, c.baseURI+URIMarketfeedOHLC, nil, input, headers, &OHLCResponse{})
	if err != nil {
		return OHLCResponse{}, err
	}
	if resp.Response.StatusCode != http.StatusOK {
		return OHLCResponse{}, nil
	}
	var ohlcResponse OHLCResponse
	if err := json.Unmarshal(resp.Body, &ohlcResponse); err != nil {
		return OHLCResponse{}, err
	}

	return ohlcResponse, nil
}

func (c *Client) GetMarketDepth(input MarketDataInput) (MarketDepthResponse, error) {
	headers := http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"access-token": {c.accessToken},
		"client-id":    {c.GetDhanClientId()},
	}
	resp, err := c.httpClient.DoJSON(http.MethodPost, c.baseURI+URIMarketfeedQuote, nil, input, headers, &MarketDepthResponse{})
	if err != nil {
		return MarketDepthResponse{}, err
	}
	if resp.Response.StatusCode != http.StatusOK {
		return MarketDepthResponse{}, nil
	}
	var marketDepthResponse MarketDepthResponse
	if err := json.Unmarshal(resp.Body, &marketDepthResponse); err != nil {
		return MarketDepthResponse{}, err
	}

	return marketDepthResponse, nil
}
