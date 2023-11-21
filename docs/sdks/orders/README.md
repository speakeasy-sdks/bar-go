# Orders
(*Orders*)

## Overview

The orders endpoints.

### Available Operations

* [CreateOrder](#createorder) - Create an order.

## CreateOrder

Create an order for a drink.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bar-go/models/components"
	bargo "github.com/speakeasy-sdks/bar-go"
	"context"
	"log"
)

func main() {
    s := bargo.New(
        bargo.WithSecurity(components.Security{
            APIKey: bargo.String("<YOUR_API_KEY>"),
        }),
    )


    requestBody := []components.OrderInput{
        components.OrderInput{
            Type: components.OrderTypeIngredient,
            ProductCode: "AC-A2DF3",
            Quantity: 138554,
        },
    }

    var callbackURL *string = "string"

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, requestBody, callbackURL)
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `requestBody`                                                    | [][components.OrderInput](../../models/components/orderinput.md) | :heavy_check_mark:                                               | N/A                                                              |
| `callbackURL`                                                    | **string*                                                        | :heavy_minus_sign:                                               | The url to call when the order is updated.                       |


### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |
