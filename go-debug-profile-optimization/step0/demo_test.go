package main

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleHi_Recorder(t *testing.T) {
	rw := httptest.NewRecorder()
	handleHi(rw, req(t, "GET / HTTP/1.0\r\n\r\n"))
	if !strings.Contains(rw.Body.String(), "visitor number") {
		t.Errorf("Unexpected output: %s", rw.Body)
	}
}

func req(t *testing.T, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestHandleHi_TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleHi))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	if g, w := res.Header.Get("Content-Type"), "text/html; charset=utf-8"; g != w {
		t.Errorf("Content-Type = %q; want %q", g, w)
	}
	slurp, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("Got: %s", slurp)
}
