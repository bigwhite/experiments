package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
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
	"strconv"
	"strings"
	"time"
)

const (
	token          = "wechat4go"
	appID          = "wx5b5c2614d269ddb2"
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
	//CreateTime   time.Duration
	CreateTime string
	MsgType    CDATAText
	Content    CDATAText
}

type EncryptRequestBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

type EncryptResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      CDATAText
	MsgSignature CDATAText
	TimeStamp    string
	Nonce        CDATAText
}

type EncryptResponseBody1 struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string
	MsgSignature string
	TimeStamp    string
	Nonce        string
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
	textResponseBody.CreateTime = strconv.Itoa(int(time.Duration(time.Now().Unix())))
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

func makeEncryptResponseBody(fromUserName, toUserName, content, nonce, timestamp string) ([]byte, error) {
	encryptBody := &EncryptResponseBody{}

	encryptXmlData, _ := makeEncryptXmlData(fromUserName, toUserName, timestamp, content)
	encryptBody.Encrypt = value2CDATA(encryptXmlData)
	encryptBody.MsgSignature = value2CDATA(makeMsgSignature(timestamp, nonce, encryptXmlData))
	encryptBody.TimeStamp = timestamp
	encryptBody.Nonce = value2CDATA(nonce)

	return xml.MarshalIndent(encryptBody, " ", "  ")
}

func makeEncryptXmlData(fromUserName, toUserName, timestamp, content string) (string, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(fromUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = timestamp
	body, err := xml.MarshalIndent(textResponseBody, " ", "  ")
	if err != nil {
		return "", errors.New("xml marshal error")
	}

	buf := bytes.NewBuffer(make([]byte, 4))
	fmt.Println("====len of body = ", len(body))
	binary.Write(buf, binary.BigEndian, len(body))
	bodyLength := buf.Bytes()

	randomBytes := []byte("abcdefghijklmnop")

	xmlData := bytes.Join([][]byte{randomBytes, bodyLength, body, []byte(appID)}, nil)
	return aesEncode(xmlData, realAESKey)
}

// PadLength calculates padding length, from github.com/vgorin/cryptogo
func PadLength(slice_length, blocksize int) (padlen int) {
	padlen = blocksize - slice_length%blocksize
	if padlen == 0 {
		padlen = blocksize
	}
	return padlen
}

//from github.com/vgorin/cryptogo
func PKCS7Pad(message []byte, blocksize int) (padded []byte) {
	// block size must be bigger or equal 2
	if blocksize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	// block size up to 255 requires 1 byte padding
	if blocksize < 1<<8 {
		// calculate padding length
		padlen := PadLength(len(message), blocksize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

		// apply padding
		padded = append(message, padding...)
		return padded
	}
	// block size bigger or equal 256 is not currently supported
	panic("unsupported block size")
}

func aesEncode(data []byte, aesKey []byte) (string, error) {
	if len(data)%aes.BlockSize != 0 {
		data = PKCS7Pad(data, aes.BlockSize)
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	encryptData := make([]byte, aes.BlockSize+len(data))
	iv := encryptData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(encryptData[aes.BlockSize:], data)

	return base64.StdEncoding.EncodeToString(encryptData[aes.BlockSize:]), nil
}

func aesDecode(encryptData string, aesKey []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("crypto/cipher: ciphertext too short")
	}

	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not correct")
	}

	iv := data[:aes.BlockSize]
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

func parseEncryptResponse(responseEncryptTextBody []byte) {
	textResponseBody := &EncryptResponseBody1{}
	xml.Unmarshal(responseEncryptTextBody, textResponseBody)

	if !validateMsg(textResponseBody.TimeStamp, textResponseBody.Nonce, textResponseBody.Encrypt, textResponseBody.MsgSignature) {
		fmt.Println("msg signature is invalid")
		return
	}

	data, err := aesDecode(textResponseBody.Encrypt, realAESKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func procRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signature := strings.Join(r.Form["signature"], "")
	encryptType := strings.Join(r.Form["encrypt_type"], "")
	msgSignature := strings.Join(r.Form["msg_signature"], "")

	fmt.Println("timestamp = ", timestamp)
	fmt.Println("nonce= ", nonce)
	fmt.Println("signature= ", signature)
	fmt.Println("msgSignature=", msgSignature)

	if !validateUrl(timestamp, nonce, signature) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}

	if r.Method == "POST" {
		if encryptType == "aes" {
			log.Println("Wechat Service: in safe mode")
			encryptRequestBody := parseEncryptRequestBody(r)
			fmt.Println("\n")
			fmt.Println(encryptRequestBody.Encrypt)
			if !validateMsg(timestamp, nonce, encryptRequestBody.Encrypt, msgSignature) {
				log.Println("Wechat Service: msg_signature is invalid")
				return
			}
			log.Println("Wechat Service: msg_signature validation is ok!")

			originData, err := aesDecode(encryptRequestBody.Encrypt, realAESKey)
			if err != nil {
				fmt.Println(err)
				return
			}

			textRequestBody, _ := parseOriginalData(originData)
			fmt.Println(textRequestBody)
			fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
				textRequestBody.Content,
				textRequestBody.FromUserName)

			responseEncryptTextBody, _ := makeEncryptResponseBody(textRequestBody.ToUserName,
				textRequestBody.FromUserName,
				"Hello, "+textRequestBody.FromUserName,
				nonce,
				timestamp)
			w.Header().Set("Content-Type", "text/xml")
			fmt.Println("\n", string(responseEncryptTextBody))
			fmt.Fprintf(w, string(responseEncryptTextBody))

			parseEncryptResponse(responseEncryptTextBody)
		} else if encryptType == "raw" {
			log.Println("Wechat Service: in raw mode")
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
