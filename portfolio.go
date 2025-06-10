package dhanhq

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Position struct {
	DhanClientId          string  `json:"dhanClientId"`
	TradingSymbol         string  `json:"tradingSymbol"`
	SecurityId            string  `json:"securityId"`
	PositionType          string  `json:"positionType"`
	ExchangeSegment       string  `json:"exchangeSegment"`
	ProductType           string  `json:"productType"`
	BuyAvg                float64 `json:"buyAvg"`
	CostPrice             float64 `json:"costPrice"`
	BuyQty                int32   `json:"buyQty"`
	SellAvg               float64 `json:"sellAvg"`
	SellQty               int32   `json:"sellQty"`
	NetQty                int32   `json:"netQty"`
	RealizedProfit        float64 `json:"realizedProfit"`
	UnrealizedProfit      float64 `json:"unrealizedProfit"`
	RbiReferenceRate      float64 `json:"rbiReferenceRate"`
	Multiplier            int32   `json:"multiplier"`
	CarryForwardBuyQty    int32   `json:"carryForwardBuyQty"`
	CarryForwardSellQty   int32   `json:"carryForwardSellQty"`
	CarryForwardBuyValue  float64 `json:"carryForwardBuyValue"`
	CarryForwardSellValue float64 `json:"carryForwardSellValue"`
	DayBuyQty             int32   `json:"dayBuyQty"`
	DaySellQty            int32   `json:"daySellQty"`
	DayBuyValue           float64 `json:"dayBuyValue"`
	DaySellValue          float64 `json:"daySellValue"`
	DrvExpiryDate         string  `json:"drvExpiryDate"`
	DrvOptionType         string  `json:"drvOptionType"`
	DrvStrikePrice        float64 `json:"drvStrikePrice"`
	CrossCurrency         bool    `json:"crossCurrency"`
}

type Positions struct {
	Positions []Position
}

type Holding struct {
	Exchange        string  `json:"exchange"`
	TradingSymbol   string  `json:"tradingSymbol"`
	SecurityId      string  `json:"securityId"`
	ISIN            string  `json:"isin"`
	TotalQty        int32   `json:"totalQty"`
	DpQty           int32   `json:"dpQty"`
	T1Qty           int32   `json:"t1Qty"`
	MtfT1Qty        int32   `json:"mtf_tq_qty"`
	MtfQty          int32   `json:"mtf_qty"`
	AvailableQty    int32   `json:"availableQty"`
	CollateralQty   int32   `json:"collateralQty"`
	AvgCostPrice    float64 `json:"avgCostPrice"`
	LastTradedPrice float64 `json:"lastTradedPrice"`
}

type Holdings struct {
	Holdings []Holding `json:"holdings"`
}

type FundLimit struct {
	DhanClientId        string  `json:"dhanClientId"`
	AvailabelBalance    float64 `json:"availabelBalance"` // typo in api response as seen at https://api.dhan.co/v2/#/operations/fundlimit
	SodLimit            float64 `json:"sodLimit"`
	CollateralAmount    float64 `json:"collateralAmount"`
	ReceivableAmount    float64 `json:"receivableAmount"`
	UtilizedAmount      float64 `json:"utilizedAmount"`
	BlockedPayoutAmount float64 `json:"blockedPayoutAmount"`
	WithdrawableBalance float64 `json:"withdrawableBalance"`
}

type ConvertPositionRequest struct {
	DhanClientId    string `json:"dhanClientId"`
	FromProductType string `json:"fromProductType"`
	ExchangeSegment string `json:"exchangeSegment"`
	PositionType    string `json:"positionType"`
	SecurityId      string `json:"securityId"`
	TradingSymbol   string `json:"tradingSymbol"`
	ConvertQty      int32  `json:"convertQty"`
	ToProductType   string `json:"toProductType"`
}

// GetPositions retrieves the positions for a given client
func (c *Client) GetPositions() (Positions, error) {
	headers := http.Header{
		"access-token": {c.accessToken},
	}

	resp, err := c.httpClient.Do(http.MethodGet, c.baseURI+URIPositions, headers, nil)
	if err != nil {
		return Positions{}, err
	}

	var positionsSlice []Position
	if err = json.Unmarshal(resp.Body, &positionsSlice); err != nil {
		return Positions{}, err
	}

	positions := Positions{
		Positions: positionsSlice,
	}
	return positions, nil
}

// GetHoldings retrieves the holdings for a given client
func (c *Client) GetHoldings() (Holdings, error) {
	headers := http.Header{
		"access-token": {c.accessToken},
	}

	resp, err := c.httpClient.Do(http.MethodGet, c.baseURI+URIHoldings, headers, nil)
	if err != nil {
		return Holdings{}, err
	}

	// Check if the response contains an error message
	var errorResp ErrorResponse
	if err = json.Unmarshal(resp.Body, &errorResp); err == nil {
		if errorResp.ErrorMessage != "" {
			// If the error code is "DH-1111", it indicates no holdings found
			if errorResp.ErrorCode == "DH-1111" {
				return Holdings{
					Holdings: []Holding{},
				}, nil
			}
			return Holdings{}, fmt.Errorf("API error: %s (Code: %s, Type: %s)",
				errorResp.ErrorMessage, errorResp.ErrorCode, errorResp.ErrorType)
		}
	}

	var holdingsSlice []Holding
	if err = json.Unmarshal(resp.Body, &holdingsSlice); err != nil {
		return Holdings{}, err
	}

	holdings := Holdings{
		Holdings: holdingsSlice,
	}
	return holdings, nil
}

func (c *Client) GetFundLimit() (FundLimit, error) {
	headers := http.Header{
		"access-token": {c.accessToken},
	}
	resp, err := c.httpClient.Do(http.MethodGet, c.baseURI+URIFundLimit, headers, nil)
	if err != nil {
		return FundLimit{}, err
	}

	var fundLimit FundLimit
	if err = json.Unmarshal(resp.Body, &fundLimit); err != nil {
		return FundLimit{}, err
	}

	return fundLimit, nil
}

func (c *Client) ConvertPosition(req ConvertPositionRequest) error {
	headers := http.Header{
		"access-token": {c.accessToken},
	}

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.httpClient.DoJSON(http.MethodPost, c.baseURI+URIPositionConvert, nil, body, headers, nil)

	if err != nil {
		return fmt.Errorf("failed to convert position: %w", err)
	}

	var errorResp ErrorResponse
	if err = json.Unmarshal(resp.Body, &errorResp); err == nil && errorResp.ErrorMessage != "" {
		return fmt.Errorf("API error: %s (Code: %s, Type: %s)",
			errorResp.ErrorMessage, errorResp.ErrorCode, errorResp.ErrorType)
	}

	return nil
}
