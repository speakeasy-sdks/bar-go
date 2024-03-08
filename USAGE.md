<!-- Start SDK Example Usage [usage] -->
### Sign in

First you need to send an authentication request to the API by providing your username and password.
In the request body, you should specify the type of token you would like to receive: API key or JSON Web Token.
If your credentials are valid, you will receive a token in the response object: `res.object.token: str`.

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/bar"
	"log"
)

func main() {
	s := bar.New()

	operationSecurity := bar.LoginSecurity{
		Username: "<USERNAME>",
		Password: "<PASSWORD>",
	}

	ctx := context.Background()
	res, err := s.Authentication.Login(ctx, bar.LoginRequestBody{
		Type: bar.TypeAPIKey,
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
	"github.com/speakeasy-sdks/bar"
	"log"
)

func main() {
	s := bar.New(
		bar.WithSecurity(bar.Security{
			APIKey: bar.String("<YOUR_API_KEY>"),
		}),
	)

	var drinkType *DrinkType = bar.DrinkTypeSpirit.ToPointer()

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
	"github.com/speakeasy-sdks/bar"
	"log"
)

func main() {
	s := bar.New(
		bar.WithSecurity(bar.Security{
			APIKey: bar.String("<YOUR_API_KEY>"),
		}),
	)

	requestBody := []bar.OrderInput{
		bar.OrderInput{
			Type:        bar.OrderTypeIngredient,
			ProductCode: "AC-A2DF3",
			Quantity:    138554,
		},
	}

	var callbackURL *string = bar.String("<value>")

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

### Subscribe to webhooks to receive stock updates

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/bar"
	"log"
)

func main() {
	s := bar.New(
		bar.WithSecurity(bar.Security{
			APIKey: bar.String("<YOUR_API_KEY>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Config.SubscribeToWebhooks(ctx, []bar.RequestBody{
		bar.RequestBody{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->