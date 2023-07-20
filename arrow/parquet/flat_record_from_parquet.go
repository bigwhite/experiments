package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/apache/arrow/go/v13/parquet/file"
	"github.com/apache/arrow/go/v13/parquet/pqarrow"
)

func main() {
	f, err := os.Open("flat_record.parquet")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr, err := file.NewParquetReader(f)
	if err != nil {
		panic(err)
	}
	defer rdr.Close()

	arrRdr, err := pqarrow.NewFileReader(rdr,
		pqarrow.ArrowReadProperties{
			BatchSize: 3,
		}, memory.DefaultAllocator)
	if err != nil {
		panic(err)
	}

	s, _ := arrRdr.Schema()
	fmt.Println(*s)

	// put all records into reader stream
	rr, err := arrRdr.GetRecordReader(context.Background(), nil, nil)
	if err != nil {
		panic(err)
	}

	for {
		rec, err := rr.Read()
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(rec)
	}
}
