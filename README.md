# thailandpost-sdk-go

Unofficial Go SDK for Thailand Post's Track And Track API

## Installation
1. Use command below to install SDK
```
$ go get -u github.com/xerodotc/thailandpost-sdk-go
```
2. Import into your code
```go
import "github.com/xerodotc/thailandpost-sdk-go"
```

## Usage / Quickstart

### API

Initialize API with:
```go
api := thailandpost.TrackingAPIInit(thailandpost.LangEN, "YOUR_TOKEN")
```

or if you want to use custom HTTP client:
```go
client := &http.Client{Timeout: 5 * time.Second}
api := thailandpost.TrackingAPIInitWithClient(thailandpost.LangEN,
    "YOUR_TOKEN", client)
```

use `api.GetItems("AAXXXXXXXXXTH")` to retrieve item's tracking status
```go
status, err := api.GetItems("AAXXXXXXXXXTH")
```

### Webhook

Initialize Webhook API with:
```go
api := thailandpost.TrackingWebhookInit(thailandpost.LangEN,
    "YOUR_TOKEN", "YOUR_BEARER_TOKEN")
```

use `api.HookTrack("AAXXXXXXXXXTH")` to make webhook send data
for specified item:
```go
status, err := api.HookTrack("AAXXXXXXXXXTH")
```

for parsing incoming webhook request, use `api.ParseHookData(req)`
(an example is using [Gin Web Framework](https://github.com/gin-gonic/gin)):
```go
func handler(c *gin.Context) {
    data, err := api.ParseHookData(c.Request)
}
```

## License / Copyright
This project is licensed under MIT license.
See [LICENSE](LICENSE) file for details.

The author of this project is not related to Thailand Post in anyway.

Thailand Post is a trademark of Thailand Post Co., Ltd.
