package main

import (
	"encoding/hex"

	"github.com/bigwhite/xmltest/xml"

	"fmt"
)

type DescCDATA struct {
	Desc string `xml:",cdata"`
}

type Person struct {
	Name string    `xml:"name"`
	Age  int       `xml:"age"`
	Desc DescCDATA `xml:"desc"`
}

var profileFmt = `<person>
<name>"tony bai"</name>
<age>33</age>
<desc><![CDATA[%s]]></desc>
</person>`

func main() {
	c := fmt.Sprintf(profileFmt, "hello\r\nxml")
	var p Person
	err := xml.Unmarshal([]byte(c), &p)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Println("unmarshal ok")

	fmt.Println(hex.Dump([]byte("hello\r\nxml")))
	fmt.Println(hex.Dump([]byte(p.Desc.Desc)))
}
