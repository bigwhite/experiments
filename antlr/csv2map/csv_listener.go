package main

import (
	"csvparser/parser"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CSVMapListener struct {
	*parser.BaseCSVListener
	headers []string
	cm      []map[string]string
	fields  []string // a slice of fields in current row
}

func (cl *CSVMapListener) lastHeader(header string) bool {
	return header == cl.headers[len(cl.headers)-1]
}

func (cl *CSVMapListener) String() string {
	var s strings.Builder
	s.WriteString("[")

	for i, m := range cl.cm {
		s.WriteString("{")
		for _, h := range cl.headers {
			s.WriteString(fmt.Sprintf("%s=%v", h, m[h]))
			if !cl.lastHeader(h) {
				s.WriteString(", ")
			}
		}
		s.WriteString("}")
		if i != len(cl.cm)-1 {
			s.WriteString(",\n")
			continue
		}
	}
	s.WriteString("]")
	return s.String()
}

/* for debug
func (this *CSVMapListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
*/

func (cl *CSVMapListener) ExitHdr(c *parser.HdrContext) {
	cl.headers = cl.fields
}

func (cl *CSVMapListener) ExitField(c *parser.FieldContext) {
	cl.fields = append(cl.fields, c.GetText())
}

func (cl *CSVMapListener) EnterRow(c *parser.RowContext) {
	cl.fields = []string{} // create a new field slice
}

func (cl *CSVMapListener) ExitRow(c *parser.RowContext) {
	// get the rule index of parent context
	if i, ok := c.GetParent().(antlr.RuleContext); ok {
		if i.GetRuleIndex() == parser.CSVParserRULE_hdr {
			// ignore this row
			return
		}
	}

	// it is a data row
	m := map[string]string{}

	for i, h := range cl.headers {
		m[h] = cl.fields[i]
	}
	cl.cm = append(cl.cm, m)
}
