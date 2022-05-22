package main

import "tdat/semantic"

type Processor struct {
	name  string // for ruleid
	model *semantic.Model
}

func (p *Processor) Exec(in []map[string]interface{}) (map[string]interface{}, error) {
	return p.model.Exec(in)
}
