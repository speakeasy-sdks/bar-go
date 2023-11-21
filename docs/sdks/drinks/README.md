# Drinks
(*Drinks*)

## Overview

The drinks endpoints.

### Available Operations

* [ListDrinks](#listdrinks) - Get a list of drinks.
* [GetDrink](#getdrink) - Get a drink.

## ListDrinks

Get a list of drinks, if authenticated this will include stock levels and product codes otherwise it will only include public information.

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


    var drinkType *components.DrinkType = components.DrinkTypeSpirit

    ctx := context.Background()
    res, err := s.Drinks.ListDrinks(ctx, drinkType)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `drinkType`                                                                  | [*components.DrinkType](../../models/components/drinktype.md)                | :heavy_minus_sign:                                                           | The type of drink to filter by. If not provided all drinks will be returned. |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |


### Response

**[*operations.ListDrinksResponse](../../models/operations/listdrinksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

## GetDrink

Get a drink by name, if authenticated this will include stock levels and product codes otherwise it will only include public information.

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


    var name string = "string"

    ctx := context.Background()
    res, err := s.Drinks.GetDrink(ctx, name)
    if err != nil {
        log.Fatal(err)
    }

    if res.Drink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `name`                                                | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetDrinkResponse](../../models/operations/getdrinkresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |
