package main

import (
	"log"
	"os"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

const (
	text1 = `在苏黎世的FIFA颁奖典礼上，巴萨球星、阿根廷国家队队长梅西赢得了生涯第5个金球奖，继续创造足坛的新纪录`
	text2 = `12月6日，网上出现照片显示国产第五代战斗机歼-20的尾翼已经涂上五位数部队编号`
	text3 = `你们很感兴趣的 .NET Core 1.1 来了哦`
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
		SegmenterDictionaries:   "./dict/dictionary.txt",
		StopTokenFile:           "./dict/stop_tokens.txt",
	})
	defer searcher.Close()

	os.MkdirAll("./index", 0777)

	docId++
	searcher.IndexDocument(docId, types.DocumentIndexData{Content: text1}, false)
	docId++
	searcher.IndexDocument(docId, types.DocumentIndexData{Content: text2}, false)
	docId++
	searcher.IndexDocument(docId, types.DocumentIndexData{Content: text3}, false)

	searcher.FlushIndex()
	log.Println("Created index number:", searcher.NumDocumentsIndexed())
}
