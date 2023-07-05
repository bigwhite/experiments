package main

import (
	"fmt"
	"io"
	"os"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/csv"
)

func read(data io.ReadCloser) error {
	var total int64
	// read 10000 lines at a time to create record batches
	rdr := csv.NewInferringReader(data, csv.WithChunk(10000),
		// strings can be null, and these are the values
		// to consider as null
		csv.WithNullReader(true, "", "null", "[]"),
		// assume the first line is a header line which names the columns
		csv.WithHeader(true),
		csv.WithColumnTypes(map[string]arrow.DataType{
			" _pressurem": arrow.PrimitiveTypes.Float64,
		}),
		csv.WithIncludeColumns([]string{
			"datetime_utc", // 19961101-11:00
			" _conds",      // Smoke、Haze
			" _fog",        // 0
			" _heatindexm",
			" _pressurem", //
			" _rain",      //
			" _snow",      //
			" _tempm",     //
			" _thunder",   //
			" _tornado",   //
		}),
	)

	for rdr.Next() {
		rec := rdr.Record()
		total += rec.NumRows()
	}

	fmt.Println("total columns =", total)
	return nil
}

func main() {
	data, err := os.Open("./testset.csv")
	if err != nil {
		panic(err)
	}
	read(data)
}
