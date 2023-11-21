<!-- Start SDK Example Usage -->
### Sign in

First you need to send an authentication request to the API by providing your username and password.
In the request body, you should specify the type of token you would like to receive: API key or JSON Web Token.
If your credentials are valid, you will receive a token in the response object: `res.object.token: str`.

```go
package main

import (
	"context"
	bargo "github.com/speakeasy-sdks/bar-go"
	"github.com/speakeasy-sdks/bar-go/models/operations"
	"log"
)

func main() {
	s := bargo.New()

	operationSecurity := operations.LoginSecurity{
		Username: "<USERNAME>",
		Password: "<PASSWORD>",
	}

	ctx := context.Background()
	res, err := s.Authentication.Login(ctx, operations.LoginRequestBody{
		Type: operations.TypeAPIKey,
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```

### Browse available drinks

Once you are authenticated, you can use the token you received for all other authenticated endpoints.
For example, you can filter the list of available drinks by type.

```go
package main

import (
	"context"
	bargo "github.com/speakeasy-sdks/bar-go"
	"github.com/speakeasy-sdks/bar-go/models/components"
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

### Create an order

When you submit an order, you can include a callback URL along your request.
This URL will get called whenever the supplier updates the status of your order.

```go
package main

import (
	"context"
	bargo "github.com/speakeasy-sdks/bar-go"
	"github.com/speakeasy-sdks/bar-go/models/components"
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
			Type:        components.OrderTypeIngredient,
			ProductCode: "AC-A2DF3",
			Quantity:    138554,
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
<!-- End SDK Example Usage -->