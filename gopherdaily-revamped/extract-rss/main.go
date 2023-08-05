package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	rss  = regexp.MustCompile(`<link[^>]*type="application/rss\+xml"[^>]*href="([^"]+)"`)
	atom = regexp.MustCompile(`<link[^>]*type="application/atom\+xml"[^>]*href="([^"]+)"`)
)

func main() {
	var sites = []string{
		"http://research.swtch.com",
		"https://tonybai.com",
		"https://benhoyt.com/writings",
	}

	for _, url := range sites {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		matches := rss.FindAllStringSubmatch(string(body), -1)
		if len(matches) == 0 {
			matches = atom.FindAllStringSubmatch(string(body), -1)
			if len(matches) == 0 {
				continue
			}
		}

		fmt.Printf("\"%s\" -> rss: \"%s\"\n", url, matches[0][1])
	}
}
