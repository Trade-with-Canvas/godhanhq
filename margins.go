package dhanhq

import (
	"log"
	"net/http"
)

type Margin struct {
	DhanClientId    string  `json:"dhanClientId"`
	ExchangeSegment string  `json:"exchangeSegment"`
	TransactionType string  `json:"transactionType"`
	Quantity        int32   `json:"quantity"`
	ProductType     string  `json:"productType"`
	SecurityId      string  `json:"securityId"`
	Price           float64 `json:"price"`
	TriggerPrice    float64 `json:"triggerPrice"`
}

type MarginResponse struct {
	TotalMargin         float64 `json:"totalMargin"`
	SpanMargin          float64 `json:"spanMargin"`
	ExposureMargin      float64 `json:"exposureMargin"`
	AvailableBalance    float64 `json:"availableBalance"`
	VariableMargin      float64 `json:"variableMargin"`
	InsufficientBalance float64 `json:"insufficientBalance"` // As seen at https://api.dhan.co/v2/#/operations/margincalculator
	Brokerage           float64 `json:"brokerage"`
	Leverage            float64 `json:"leverage"`
}

func (c *Client) CalculateMargins(margin Margin) (MarginResponse, error) {
	headers := http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"access-token": {c.accessToken},
	}
	var respObj MarginResponse
	httpResp, err := c.httpClient.DoJSON(http.MethodPost, c.baseURI+URIMarginCalculator, nil, margin, headers, &respObj)
	if err != nil {
		return MarginResponse{}, err
	}
	if httpResp.Response.StatusCode != http.StatusOK {
		log.Printf("Error: %s - %s", httpResp.Response.Status, string(httpResp.Body))
		return MarginResponse{}, err
	}
	return respObj, nil
}
