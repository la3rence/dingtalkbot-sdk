# DingTalk Robot SDK (Go)

## Usage

```shell
export DINGTALK_TOKEN="change me"
export DINGTALK_SECRET="change me"
```

Test sending message:

```shell
go test ./... -v
```

Send with this SDK in Go:

```go
bot := sdk.NewDingBot(os.Getenv("DINGTALK_TOKEN"), os.Getenv("DINGTALK_SECRET"))
bot.SendSimpleText("hello world")
```

## LICENSE

MIT