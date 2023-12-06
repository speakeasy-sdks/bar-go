# github.com/speakeasy-sdks/bar-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/bar-go.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bar-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/bar
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

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

	var drinkType *DrinkType = bar.DrinkTypeSpirit

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

### Subscribe to webhooks to receive stock updates

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/bar"
	"log"
	"net/http"
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

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Authentication](docs/sdks/authentication/README.md)

* [Login](docs/sdks/authentication/README.md#login) - Authenticate with the API by providing a username and password.

### [Drinks](docs/sdks/drinks/README.md)

* [ListDrinks](docs/sdks/drinks/README.md#listdrinks) - Get a list of drinks.
* [GetDrink](docs/sdks/drinks/README.md#getdrink) - Get a drink.

### [Ingredients](docs/sdks/ingredients/README.md)

* [ListIngredients](docs/sdks/ingredients/README.md#listingredients) - Get a list of ingredients.

### [Orders](docs/sdks/orders/README.md)

* [CreateOrder](docs/sdks/orders/README.md#createorder) - Create an order.

### [Config](docs/sdks/config/README.md)

* [SubscribeToWebhooks](docs/sdks/config/README.md#subscribetowebhooks) - Subscribe to webhooks.
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries.  If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API.  However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a retryConfig object to the call:
```go
package main

import (
	""
	"context"
	"github.com/speakeasy-sdks/bar"
	"github.com/speakeasy-sdks/bar/internal/utils"
	"log"
	"net/http"
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
	}, bar.WithRetries(
		utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can provide a retryConfig at SDK initialization:
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/bar"
	"github.com/speakeasy-sdks/bar/internal/utils"
	"log"
	"net/http"
)

func main() {
	s := bar.New(
		bar.WithRetryConfig(
			utils.RetryConfig{
				Strategy: "backoff",
				Backoff: &utils.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object     | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| bar.BadRequest   | 400              | application/json |
| bar.APIError     | 5XX              | application/json |
| bar.SDKError     | 400-600          | */*              |

### Example

```go
package main

import (
	"context"
	"errors"
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

		var e *BadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *bar.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `prod` | `https://speakeasy.bar` | None |
| `staging` | `https://staging.speakeasy.bar` | None |
| `customer` | `https://{organization}.{environment}.speakeasy.bar` | `organization` (default is `api`), `environment` (default is `prod`) |

#### Example

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/bar"
	"log"
)

func main() {
	s := bar.New(
		bar.WithServer("customer"),
		bar.WithSecurity(bar.Security{
			APIKey: bar.String("<YOUR_API_KEY>"),
		}),
	)

	ingredients := []string{
		"string",
	}

	ctx := context.Background()
	res, err := s.Ingredients.ListIngredients(ctx, ingredients)
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithOrganization string`
 * `WithEnvironment bar.ServerEnvironment`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/bar"
	"log"
)

func main() {
	s := bar.New(
		bar.WithServerURL("https://speakeasy.bar"),
		bar.WithSecurity(bar.Security{
			APIKey: bar.String("<YOUR_API_KEY>"),
		}),
	)

	ingredients := []string{
		"string",
	}

	ctx := context.Background()
	res, err := s.Ingredients.ListIngredients(ctx, ingredients)
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
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

	var drinkType *DrinkType = bar.DrinkTypeSpirit

	ctx := context.Background()
	res, err := s.Drinks.ListDrinks(ctx, bar.WithServerURL("https://speakeasy.bar"), drinkType)
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `APIKey`     | apiKey       | API key      |
| `BearerAuth` | http         | HTTP Bearer  |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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

	ingredients := []string{
		"string",
	}

	ctx := context.Background()
	res, err := s.Ingredients.ListIngredients(ctx, ingredients)
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
