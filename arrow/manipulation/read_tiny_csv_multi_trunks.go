package main

import (
	"fmt"
	"io"
	"os"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/csv"
)

func read(data io.ReadCloser) error {
	// read 5 lines at a time to create record batches
	rdr := csv.NewInferringReader(data, csv.WithChunk(5),
		// strings can be null, and these are the values
		// to consider as null
		csv.WithNullReader(true, "", "null", "[]"),
		// assume the first line is a header line which names the columns
		csv.WithHeader(true),
		csv.WithColumnTypes(map[string]arrow.DataType{
			" _vism": arrow.PrimitiveTypes.Float64,
		}),
	)

	for rdr.Next() {
		rec := rdr.Record()
		fmt.Println(rec)
	}

	return nil
}

func main() {
	data, err := os.Open("./testset.tiny.csv")
	if err != nil {
		panic(err)
	}
	read(data)
}
