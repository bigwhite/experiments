package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	token = "wechat4go"
	appID = "wx5b5c2614d269ddb2"
	//appsecret      = "7d1b214e5dd5b66e4daf5e71ff1a253b"
	encodingAESKey = "kZvGYbDKbtPbhv4LBWOcdsp5VktA3xe9epVhINevtGg"
)

var realAESKey []byte

func encodingAESKey2RealAESKey(wechatAESKey string) []byte {
	data, _ := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	return data
}

func init() {
	realAESKey = encodingAESKey2RealAESKey(encodingAESKey)
	fmt.Println(len(realAESKey))
}

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

type EncryptRequestBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

/*
type CDATAText struct {
	Text []byte `xml:",innerxml"`
}
*/

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func makeMsgSignature(timestamp, nonce, msg_encrypt string) string {
	sl := []string{token, timestamp, nonce, msg_encrypt}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateUrl(timestamp, nonce, signatureIn string) bool {
	signatureGen := makeSignature(timestamp, nonce)
	if signatureGen != signatureIn {
		return false
	}
	return true
}

func validateMsg(timestamp, nonce, msgEncrypt, msgSignatureIn string) bool {
	msgSignatureGen := makeMsgSignature(timestamp, nonce, msgEncrypt)
	if msgSignatureGen != msgSignatureIn {
		return false
	}
	return true
}

func parseEncryptRequestBody(r *http.Request) *EncryptRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &EncryptRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody

}

func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

func value2CDATA(v string) CDATAText {
	//return CDATAText{[]byte("<![CDATA[" + v + "]]>")}
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(fromUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

func aesDecode(encryptData string, aesKey []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(realAESKey)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("crypto/cipher: ciphertext too short")
	}

	//iv := data[:aes.BlockSize]
	iv := make([]byte, aes.BlockSize)
	blockMode := cipher.NewCBCDecrypter(block, iv)

	originData := make([]byte, len(data))
	blockMode.CryptBlocks(originData, data)
	return originData, nil
}

func validateAppId(id []byte) bool {
	if string(id) == appID {
		return true
	}
	return false
}

func parseOriginalData(originData []byte) (*TextRequestBody, error) {
	fmt.Println(string(originData))

	buf := bytes.NewBuffer(originData[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)
	fmt.Println(string(originData[20 : 20+length]))

	appIDstart := 20 + length
	id := originData[appIDstart : int(appIDstart)+len(appID)]
	if !validateAppId(id) {
		log.Println("Wechat Service: appid is invalid!")
		return nil, errors.New("Appid is invalid")
	}
	log.Println("Wechat Service: appid validation is ok!")

	textRequestBody := &TextRequestBody{}
	xml.Unmarshal(originData[20:20+length], textRequestBody)
	return textRequestBody, nil
}

func procRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signature := strings.Join(r.Form["signature"], "")
	encryptType := strings.Join(r.Form["encrypt_type"], "")
	msgSignature := strings.Join(r.Form["msg_signature"], "")

	if !validateUrl(timestamp, nonce, signature) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}

	if r.Method == "POST" {
		if encryptType == "aes" {
			log.Println("Wechat Service: in safe mode")
			encryptRequestBody := parseEncryptRequestBody(r)
			if !validateMsg(timestamp, nonce, encryptRequestBody.Encrypt, msgSignature) {
				log.Println("Wechat Service: msg_signature is invalid")
				return
			}
			log.Println("Wechat Service: msg_signature validation is ok!")

			originData, _ := aesDecode(encryptRequestBody.Encrypt, realAESKey)
			textRequestBody, _ := parseOriginalData(originData)
			fmt.Println(textRequestBody)

		} else if encryptType == "raw" {
			log.Println("Wechat Service: in raw mode")
			textRequestBody := parseTextRequestBody(r)
			if textRequestBody != nil {
				fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
					textRequestBody.Content,
					textRequestBody.FromUserName)
				responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
					textRequestBody.FromUserName,
					"Hello, "+textRequestBody.FromUserName)
				if err != nil {
					log.Println("Wechat Service: makeTextResponseBody error: ", err)
					return
				}
				w.Header().Set("Content-Type", "text/xml")
				fmt.Println(string(responseTextBody))
				fmt.Fprintf(w, string(responseTextBody))
			}
		}
	}
}

func main() {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", procRequest)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)
	}
	log.Println("Wechat Service: Stop!")
}
