package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	appID                = "wx8e0fb2659c2eexxx"
	appSecret            = "22746009b0162fe50cb915851c53fyyy"
	accessTokenFetchUrl  = "https://api.weixin.qq.com/cgi-bin/token"
	customServicePostUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

var openID = "oBQcwuAbKpiSAbbvd_DEZg7q27QI"

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}

// {
//	"touser":"OPENID",
//	"msgtype":"text",
//	"text":
//	{
//		"content":"Hello World"
//	}
// }
type CustomServiceMsg struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	Text    TextMsgContent `json:"text"`
}

type TextMsgContent struct {
	Content string `json:"content"`
}

func fetchAccessToken() (string, float64, error) {
	requestLine := strings.Join([]string{accessTokenFetchUrl,
		"?grant_type=client_credential&appid=",
		appID,
		"&secret=",
		appSecret}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", 0.0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0.0, err
	}

	fmt.Println(string(body))

	//Json Decoding
	if bytes.Contains(body, []byte("access_token")) {
		fmt.Println("return ok")
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			return "", 0.0, err
		}
		return atr.AccessToken, atr.ExpiresIn, nil
	} else {
		fmt.Println("return err")
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			return "", 0.0, err
		}
		return "", 0.0, fmt.Errorf("%s", ater.Errmsg)
	}
}

func pushCustomMsg(accessToken, toUser, msg string) error {
	csMsg := &CustomServiceMsg{
		ToUser:  toUser,
		MsgType: "text",
		Text:    TextMsgContent{Content: msg},
	}

	body, err := json.MarshalIndent(csMsg, " ", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	postReq, err := http.NewRequest("POST",
		strings.Join([]string{customServicePostUrl, "?access_token=", accessToken}, ""),
		bytes.NewReader(body))
	if err != nil {
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

func main() {
	// Fetch access_token
	accessToken, expiresIn, err := fetchAccessToken()
	if err != nil {
		log.Println("Get access_token error:", err)
		return
	}
	fmt.Println(accessToken, expiresIn)

	// Post custom service message
	msg := "你好" + "\U0001f604"
	err = pushCustomMsg(accessToken, openID, msg)
	if err != nil {
		log.Println("Push custom service message err:", err)
		return
	}
}
