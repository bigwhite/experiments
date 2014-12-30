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
	token               = "wechat4go"
	appID               = "wx8e0fb2659c2eexxx"
	appSecret           = "22746009b0162fe50cb915851c53fbbb"
	accessTokenFetchUrl = "https://api.weixin.qq.com/cgi-bin/token"
)

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
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

func main() {
	accessToken, expiresIn, err := fetchAccessToken()
	if err != nil {
		log.Println("Get access_token error:", err)
		return
	}
	fmt.Println(accessToken, expiresIn)
}
