package dingtalkbot_sdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type DingBot struct {
	token, secret string
}

type DingResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func NewDingBot(token, secret string) *DingBot {
	return &DingBot{
		token:  token,
		secret: secret,
	}
}

// sign method reference:
// https://developers.dingtalk.com/document/app/custom-robot-access
func sign(timestamp int64, secret string) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	hmc256 := hmac.New(sha256.New, []byte(secret))
	hmc256.Write([]byte(stringToSign))
	data := hmc256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

func (bot *DingBot) SendMessage(msg interface{}) error {
	sendBody := bytes.NewBuffer(nil)
	err := json.NewEncoder(sendBody).Encode(msg)
	if err != nil {
		return fmt.Errorf(`encode message %v json error: %v`, msg, err.Error())
	}

	values := url.Values{}
	values.Set(`access_token`, bot.token)
	timestamp := time.Now().Unix() * 1e3
	values.Set(`timestamp`, fmt.Sprint(timestamp))
	values.Set(`sign`, sign(timestamp, bot.secret))
	request, err := http.NewRequest(http.MethodPost, `https://oapi.dingtalk.com/robot/send`, sendBody)
	if err != nil {
		return fmt.Errorf(`build http request error: %v`, err.Error())
	}

	request.Header.Set(`Content-Type`, `application/json`)
	request.URL.RawQuery = values.Encode()
	httpClient := &http.Client{Timeout: 6 * time.Second}
	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf(`client send http request error: %v`, err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return fmt.Errorf(`fail to send message with response %s`, responseData)
	}
	if err != nil {
		return fmt.Errorf(`read response from sending message error: %v`, err.Error())
	}

	var jsonResponse DingResponse
	err = json.Unmarshal(responseData, &jsonResponse)
	if err != nil {
		return fmt.Errorf(`fail to unmarshal json from dingtalk response: %v`, err.Error())
	}

	if jsonResponse.ErrCode != 0 {
		return fmt.Errorf(`wrong configuration caused by %s`, jsonResponse.ErrMsg)
	}
	_ = response.Body.Close()
	return nil
}

// SendSimpleText to send a plain text
func (bot *DingBot) SendSimpleText(text string) error {
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": text,
		},
		"at": map[string]interface{}{
			"atMobiles": nil,
			"isAtAll":   false,
		},
	}
	return bot.SendMessage(message)
}
