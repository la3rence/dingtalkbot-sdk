package dingtalkbot_sdk

import (
	"os"
	"testing"
)

func TestSign(t *testing.T) {
	hash := sign(1000000000000, "DINGTALK-SECRET")
	if hash != "ZPQgSHjVIaMVfVZwrHP3Wjy3zDDNEs6LPcnDChf+2M0=" {
		t.Errorf("wrong sign method")
	}
}

func TestDingBot_SendSimpleText(t *testing.T) {
	dingBot := NewDingBot(os.Getenv("DINGTALK_TOKEN"), os.Getenv("DINGTALK_SECRET"))
	err := dingBot.SendSimpleText("`go test`: PASS")
	if err != nil {
		t.Errorf(err.Error())
	}
}
