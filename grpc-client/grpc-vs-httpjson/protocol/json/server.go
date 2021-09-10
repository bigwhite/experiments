package main

import (
	"log"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type JsonMessage struct {
	ClientId string `json:"clientid"`
	Topic    string `json:"topic"`
	Payload  []byte `json:"payload"`
}

func HandleMessage(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	var m JsonMessage
	err := sonic.Unmarshal(body, &m)
	if err != nil {
		ctx.Response.Header.SetStatusCode(http.StatusBadRequest)
		log.Println(err)
		return
	}
	//log.Println(m)
}

func main() {
	r := router.New()
	r.POST("/", HandleMessage)

	server := &fasthttp.Server{
		Handler: r.Handler,
	}

	addr := "127.0.0.1:10001"
	err := server.ListenAndServe(addr)
	if err != nil {
		panic(err)
	}
}
