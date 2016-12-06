package main

import (
	"fmt"
	"log"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

var (
	searcher = engine.Engine{}
)

func main() {
	searcher.Init(types.EngineInitOptions{
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.DocIdsIndex,
		},
		UsePersistentStorage:    true,
		PersistentStorageFolder: "./index",
		SegmenterDictionaries:   "./dict/dictionary.txt",
		StopTokenFile:           "./dict/stop_tokens.txt",
	})
	defer searcher.Close()

	searcher.FlushIndex()
	log.Println("recover index number:", searcher.NumDocumentsIndexed())

	fmt.Printf("%#v\n", searcher.Search(types.SearchRequest{Text: "巴萨 梅西"}))
}
