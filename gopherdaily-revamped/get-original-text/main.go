package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	s, err := getOriginText("http://research.swtch.com/coro")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

func getOriginText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	converter := md.NewConverter("", true, nil).Remove("header",
		"footer", "aside", "table", "nav") //"table" is used to store code

	markdown, err := converter.ConvertString(string(body))
	if err != nil {
		return "", err
	}
	return markdown, nil
}
