package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {

	var feeds = []string{
		"https://research.swtch.com/feed.atom",
		"https://tonybai.com/feed/",
		"https://benhoyt.com/writings/rss.xml",
	}

	fp := gofeed.NewParser()
	for _, feed := range feeds {
		feedInfo, err := fp.ParseURL(feed)
		if err != nil {
			fmt.Printf("parse feed [%s] error: %s\n", feed, err.Error())
			continue
		}
		fmt.Printf("The info of feed url: %s\n", feed)
		for _, item := range feedInfo.Items {
			fmt.Printf("\t title: %s\n", item.Title)
			fmt.Printf("\t link: %s\n", item.Link)
			fmt.Printf("\t published: %s\n", item.Published)
		}
		fmt.Println("")
	}
}
