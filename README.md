# DhanHQ Go SDK

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

[Consent Login](https://github.com/tradewithcanvas/godhanhq/tree/main/examples)

[Margins](https://github.com/tradewithcanvas/godhanhq/tree/main/examples)

[Portfolio](https://github.com/tradewithcanvas/godhanhq/tree/main/examples)
