package Common

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

type ticket struct {
	ErrCode   int    `json:"errcode,omitempty"`
	ErrMsg    string `json:"errmsg,omitempty"`
	Ticket    string `json:"ticket,omitempty"`
	ExpiresIn int    `json:"expires_in,omitempty"`
}

type WeChat struct {
	Timestamp string
	NonceStr  string
	Signature string
}

func getWeChatToken() (*token, error) {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx18d4d0547eee571e&secret=f6810d82b1765ac5fc84003385fbb7b0")

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var tk token

	_ = json.Unmarshal(body, &tk)

	return &tk, nil
}

func getWeChatTicket() (*ticket, error) {
	tk, _ := getWeChatToken()

	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + tk.AccessToken + "&type=jsapi")

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var tc ticket

	_ = json.Unmarshal(body, &tc)

	return &tc, nil
}

func GetWeChat(url string) *WeChat {
	var timestamp, nonceStr, signature string

	tc, _ := getWeChatTicket()

	timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	nonceStr = string(Random(15, 3))

	longStr := "jsapi_ticket=" + tc.Ticket + "&noncestr=" + nonceStr + "&timestamp=" + timestamp + "&url=" + url

	h := sha1.New()

	if _, err := h.Write([]byte(longStr)); err != nil {
		panic(err.Error())
	}

	signature = fmt.Sprintf("%x", h.Sum(nil))

	return &WeChat{
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		Signature: signature,
	}
}
