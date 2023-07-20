package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/apache/arrow/go/v13/parquet"
	"github.com/apache/arrow/go/v13/parquet/pqarrow"
)

func main() {
	f, err := os.Open("flat_table_compressed.parquet")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tbl, err := pqarrow.ReadTable(context.Background(), f, parquet.NewReaderProperties(memory.DefaultAllocator),
		pqarrow.ArrowReadProperties{}, memory.DefaultAllocator)
	if err != nil {
		panic(err)
	}

	dumpTable(tbl)
}

func dumpTable(tbl arrow.Table) {
	s := tbl.Schema()
	fmt.Println(s)
	fmt.Println("------")

	fmt.Println("the count of table columns=", tbl.NumCols())
	fmt.Println("the count of table rows=", tbl.NumRows())
	fmt.Println("------")

	for i := 0; i < int(tbl.NumCols()); i++ {
		col := tbl.Column(i)
		fmt.Printf("arrays in column(%s):\n", col.Name())
		chunk := col.Data()
		for _, arr := range chunk.Chunks() {
			fmt.Println(arr)
		}
		fmt.Println("------")
	}
}
