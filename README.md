# DingTalk Robot SDK (Go)

## Usage

```shell
export DINGTALK_TOKEN="change me"
export DINGTALK_SECRET="change me"
```

Test sending messages:

```shell
go test ./... -v
```

Send messages with this SDK in Go:

```shell
go get -u github.com/Lonor/dingtalkbot-sdk
```

```go
package main

import (
	sdk "github.com/Lonor/dingtalkbot-sdk"
	"os"
)

func main() {
	bot := sdk.NewDingBot(os.Getenv("DINGTALK_TOKEN"), os.Getenv("DINGTALK_SECRET"))
	_ = bot.SendSimpleText("hello world")
}
```

## LICENSE

MIT