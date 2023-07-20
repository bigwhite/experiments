package main

import (
	"os"
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/apache/arrow/go/v13/parquet"
	"github.com/apache/arrow/go/v13/parquet/compress"
	"github.com/apache/arrow/go/v13/parquet/pqarrow"
)

func main() {
	var records []arrow.Record
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "archer", Type: arrow.BinaryTypes.String},
			{Name: "location", Type: arrow.BinaryTypes.String},
			{Name: "year", Type: arrow.PrimitiveTypes.Int16},
		},
		nil,
	)

	rb := array.NewRecordBuilder(memory.DefaultAllocator, schema)
	defer rb.Release()

	for i := 0; i < 3; i++ {
		postfix := strconv.Itoa(i)
		rb.Field(0).(*array.StringBuilder).AppendValues([]string{"tony" + postfix, "amy" + postfix, "jim" + postfix}, nil)
		rb.Field(1).(*array.StringBuilder).AppendValues([]string{"beijing" + postfix, "shanghai" + postfix, "chengdu" + postfix}, nil)
		rb.Field(2).(*array.Int16Builder).AppendValues([]int16{1992 + int16(i), 1993 + int16(i), 1994 + int16(i)}, nil)
		rec := rb.NewRecord()
		records = append(records, rec)
	}

	// write to parquet
	f, err := os.Create("flat_record_compressed.parquet")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	props := parquet.NewWriterProperties(parquet.WithCompression(compress.Codecs.Zstd),
		parquet.WithCompressionFor("year", compress.Codecs.Brotli))
	writer, err := pqarrow.NewFileWriter(schema, f, props,
		pqarrow.DefaultWriterProps())
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	for _, rec := range records {
		if err := writer.Write(rec); err != nil {
			panic(err)
		}
		rec.Release()
	}
}
