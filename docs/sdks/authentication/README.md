# Authentication
(*Authentication*)

## Overview

The authentication endpoints.

### Available Operations

* [Login](#login) - Authenticate with the API by providing a username and password.

## Login

Authenticate with the API by providing a username and password.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bar"
	"context"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [LoginRequestBody](../../loginrequestbody.md)         | :heavy_check_mark:                                    | The request object to use for the request.            |
| `security`                                            | [LoginSecurity](../../loginsecurity.md)               | :heavy_check_mark:                                    | The security requirements to use for the request.     |


### Response

**[*LoginResponse](../../loginresponse.md), error**
| Error Object     | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| bar.APIError     | 5XX              | application/json |
| bar.SDKError     | 4xx-5xx          | */*              |
