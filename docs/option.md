## Options

### WithServerURL

WithServerURL allows providing an alternative server URL.

```go
bar.WithServerURL("http://api.example.com")
```

## WithTemplatedServerURL

WithTemplatedServerURL allows providing an alternative server URL with templated parameters.

```go
bar.WithTemplatedServerURL("http://{host}:{port}", map[string]string{
    "host": "api.example.com",
    "port": "8080",
})
```

### WithRetries

WithRetries allows customizing the default retry configuration. Only usable with methods that mention they support retries.

```go
bar.WithRetries(utils.RetryConfig{
    Strategy: "backoff",
    Backoff: utils.BackoffStrategy{
        InitialInterval: 500 * time.Millisecond,
        MaxInterval: 60 * time.Second,
        Exponent: 1.5,
        MaxElapsedTime: 5 * time.Minute,
    },
    RetryConnectionErrors: true,
})
```