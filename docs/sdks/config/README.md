# Config
(*Config*)

### Available Operations

* [SubscribeToWebhooks](#subscribetowebhooks) - Subscribe to webhooks.

## SubscribeToWebhooks

Subscribe to webhooks.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bar-go/models/components"
	bargo "github.com/speakeasy-sdks/bar-go"
	"context"
	"github.com/speakeasy-sdks/bar-go/models/operations"
	"log"
	"net/http"
)

func main() {
    s := bargo.New(
        bargo.WithSecurity(components.Security{
            APIKey: bargo.String("<YOUR_API_KEY>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Config.SubscribeToWebhooks(ctx, []operations.RequestBody{
        operations.RequestBody{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]operations.RequestBody](../../.md)                 | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.SubscribeToWebhooksResponse](../../models/operations/subscribetowebhooksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |
