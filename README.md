# DhanHQ Go SDK
[![Go Reference](https://pkg.go.dev/badge/github.com/tradewithcanvas/godhanhq.svg)](https://pkg.go.dev/github.com/tradewithcanvas/godhanhq)
[![Go Report Card](https://goreportcard.com/badge/github.com/tradewithcanvas/godhanhq)](https://goreportcard.com/report/github.com/tradewithcanvas/godhanhq)
### Overview

`godhanhq` is a Go SDK for interacting with the DhanHQ API. It provides a simple and efficient way to access various endpoints of the DhanHQ platform in a simple and easy manner.

### Documentation

- [GoDoc Documentation](https://godoc.org/github.com/tradewithcanvas/godhanhq)
- [Official DhanHQ v2 API Documentation](https://dhanhq.co/docs/v2/)

### Installation

```bash
go get github.com/tradewithcanvas/godhanhq
```

### Usage

```go
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
```

### Examples:

You can check the [examples](https://github.com/tradewithcanvas/godhanhq/tree/main/examples) folder for examples of usage.

#### Some examples:

[Consent Login](https://github.com/tradewithcanvas/godhanhq/tree/main/examples/consent)

[Market](https://github.com/tradewithcanvas/godhanhq/tree/main/examples/market)

[Margins](https://github.com/tradewithcanvas/godhanhq/tree/main/examples/margins)

[Portfolio](https://github.com/tradewithcanvas/godhanhq/tree/main/examples/portfolio)
