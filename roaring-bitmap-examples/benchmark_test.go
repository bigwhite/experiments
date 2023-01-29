package main

import (
	"encoding/json"
	"testing"

	"github.com/RoaringBitmap/roaring"
	"github.com/bytedance/sonic"
)

type Foo struct {
	N    int    `json:"num"`
	Name string `json:"name"`
	Addr string `json:"addr"`
	Age  string `json:"age"`
	RB   MyRB   `json:"myrb"`
}

func BenchmarkSonicJsonEncode(b *testing.B) {
	var f = Foo{
		N: 5,
		RB: MyRB{
			RB: roaring.NewBitmap(),
		},
	}

	for i := 0; i < 3000; i++ {
		f.RB.RB.Add(uint32(i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(&f)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSonicJsonDecode(b *testing.B) {
	var f = Foo{
		N: 5,
		RB: MyRB{
			RB: roaring.NewBitmap(),
		},
	}

	for i := 0; i < 3000; i++ {
		f.RB.RB.Add(uint32(i))
	}

	buf, err := sonic.Marshal(&f)
	if err != nil {
		panic(err)
	}
	var f1 = Foo{
		RB: MyRB{
			RB: roaring.NewBitmap(),
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = sonic.Unmarshal(buf, &f1)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStdJsonEncode(b *testing.B) {
	var f = Foo{
		N: 5,
		RB: MyRB{
			RB: roaring.NewBitmap(),
		},
	}

	for i := 0; i < 3000; i++ {
		f.RB.RB.Add(uint32(i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&f)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStdJsonDecode(b *testing.B) {
	var f = Foo{
		N: 5,
		RB: MyRB{
			RB: roaring.NewBitmap(),
		},
	}

	for i := 0; i < 3000; i++ {
		f.RB.RB.Add(uint32(i))
	}

	buf, err := json.Marshal(&f)
	if err != nil {
		panic(err)
	}
	var f1 = Foo{
		RB: MyRB{
			RB: roaring.NewBitmap(),
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(buf, &f1)
		if err != nil {
			panic(err)
		}
	}
}
