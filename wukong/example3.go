package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

var (
	searcher = engine.Engine{}
	docId    uint64
)

func main() {
	searcher.Init(types.EngineInitOptions{
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.DocIdsIndex,
		},
		UsePersistentStorage:    true,
		PersistentStorageFolder: "./index",
		PersistentStorageShards: 8,
		SegmenterDictionaries:   "./dict/dictionary.txt",
		StopTokenFile:           "./dict/stop_tokens.txt",
	})
	defer searcher.Close()
	searcher.FlushIndex()
	log.Println("recover index number:", searcher.NumDocumentsIndexed())
	docId = searcher.NumDocumentsIndexed()

	os.MkdirAll("./source", 0777)

	go func() {
		for {
			var paths []string

			//update index dynamically
			time.Sleep(time.Second * 10)
			var path = "./source"
			err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
				if f == nil {
					return err
				}
				if f.IsDir() {
					return nil
				}

				fc, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println("read file:", path, "error:", err)
				}

				docId++
				fmt.Println("indexing file:", path, "... ...")
				searcher.IndexDocument(docId, types.DocumentIndexData{Content: string(fc)}, true)
				fmt.Println("indexed file:", path, " ok")
				paths = append(paths, path)

				return nil
			})
			if err != nil {
				fmt.Printf("filepath.Walk() returned %v\n", err)
				return
			}

			for _, p := range paths {
				err := os.Remove(p)
				if err != nil {
					fmt.Println("remove file:", p, " error:", err)
					continue
				}
				fmt.Println("remove file:", p, " ok!")
			}

			if len(paths) != 0 {
				// 等待索引刷新完毕
				fmt.Println("flush index....")
				searcher.FlushIndex()
				fmt.Println("flush index ok")
			}
		}
	}()

	for {
		var s string
		fmt.Println("Please input your search keywords:")
		fmt.Scanf("%s", &s)
		if s == "exit" {
			break
		}

		fmt.Printf("%#v\n", searcher.Search(types.SearchRequest{Text: s}))
	}
}
