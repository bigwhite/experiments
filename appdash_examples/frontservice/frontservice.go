// Package main provides ...
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"sourcegraph.com/sourcegraph/appdash"
	"sourcegraph.com/sourcegraph/appdash/httptrace"
	"sourcegraph.com/sourcegraph/appdash/sqltrace"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var result string
	span := appdash.NewRootSpanID()
	fmt.Println("span is ", span)
	collector := appdash.NewRemoteCollector(":3001")

	httpClient := &http.Client{
		Transport: &httptrace.Transport{
			Recorder: appdash.NewRecorder(span, collector),
			SetName:  true,
		},
	}

	//Service A
	resp, err := httpClient.Get("http://localhost:6601")
	if err != nil {
		log.Println("access serviceA err:", err)
	} else {
		log.Println("access serviceA ok")
		resp.Body.Close()
		result += "access serviceA ok\n"
	}

	//Service B
	resp, err = httpClient.Get("http://localhost:6602")
	if err != nil {
		log.Println("access serviceB err:", err)
		return
	} else {
		log.Println("access serviceB ok")
		resp.Body.Close()
		result += "access serviceB ok\n"
	}

	// SQL event
	traceRec := appdash.NewRecorder(span, collector)
	traceRec.Name("sqlevent example")

	// A random length for the trace.
	length := time.Duration(rand.Intn(1000)) * time.Millisecond
	startTime := time.Now().Add(-time.Duration(rand.Intn(100)) * time.Minute)
	traceRec.Event(&sqltrace.SQLEvent{
		ClientSend: startTime,
		ClientRecv: startTime.Add(length),
		SQL:        "SELECT * FROM table_name;",
		Tag:        fmt.Sprintf("fakeTag%d", rand.Intn(10)),
	})

	result += "sql event ok\n"
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
