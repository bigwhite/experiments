package main

import (
	"fmt"

	"github.com/RoaringBitmap/roaring"
	"github.com/bytedance/sonic"
)

type MyRB struct {
	RB *roaring.Bitmap
}

func (rb *MyRB) MarshalJSON() ([]byte, error) {
	s, err := rb.RB.ToBase64()
	if err != nil {
		return nil, err
	}

	r := fmt.Sprintf(`{"rb":"%s"}`, s)
	return []byte(r), nil
}

func (rb *MyRB) UnmarshalJSON(data []byte) error {
	_, err := rb.RB.FromBase64(string(data[7 : len(data)-2]))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var myrb = MyRB{
		RB: roaring.NewBitmap(),
	}

	for i := 0; i < 31; i++ {
		myrb.RB.Add(uint32(i))
	}
	fmt.Printf("the cardinality of origin bitmap = %d\n", myrb.RB.GetCardinality())

	buf, err := sonic.Marshal(&myrb)
	if err != nil {
		panic(err)
	}

	fmt.Printf("bitmap2json: %s\n", string(buf))

	var myrb1 = MyRB{
		RB: roaring.NewBitmap(),
	}
	err = sonic.Unmarshal(buf, &myrb1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("after json2bitmap, the cardinality of new bitmap = %d\n", myrb1.RB.GetCardinality())
}
