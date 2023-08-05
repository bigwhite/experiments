package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	var allURLs []string

	err := filepath.Walk("/Users/tonybai/blog/gitee.com/gopherdaily", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".txt" && filepath.Ext(path) != ".md" {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		urlRegex := regexp.MustCompile(`https?://[^\s]+`)

		for scanner.Scan() {
			urls := urlRegex.FindAllString(scanner.Text(), -1)
			allURLs = append(allURLs, urls...)
		}

		return scanner.Err()
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, url := range allURLs {
		fmt.Printf("%s\n", url)
	}
	fmt.Println(len(allURLs))
}
