package main

import (
	"context"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	// Construct a Rego object that can be prepared or evaluated.
	r := rego.New(
		rego.Query("data.opa.examples.allow"),
		rego.Load([]string{"./policy.rego"}, nil),
	)

	// Create a prepared query that can be evaluated.
	query, err := r.PrepareForEval(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	inputs := []map[string]interface{}{
		{
			"name": "alice",
			"subject": map[string]string{
				"role": "manager",
			},
			"object": "employee_info",
			"action": "read",
		},
		{
			"name": "alice",
			"subject": map[string]string{
				"role": "manager",
			},
			"object": "employee_info",
			"action": "write",
		},
		{
			"name": "bob",
			"subject": map[string]string{
				"role": "employee",
			},
			"object": "employee_info",
			"action": "write",
		},
		{
			"name": "cathy",
			"subject": map[string]string{
				"role": "hr",
			},
			"object": "employee_info",
			"action": "read",
		},
		{
			"name": "cathy",
			"subject": map[string]string{
				"role": "hr",
			},
			"object": "employee_info",
			"action": "write",
		},
		{
			"name": "dan",
			"subject": map[string]string{
				"role": "finance",
			},
			"object": "employee_salary",
			"action": "read",
		},
		{
			"name": "bob",
			"subject": map[string]string{
				"role": "employee",
			},
			"object": "employee_salary",
			"action": "write",
		},
	}

	for _, v := range inputs {
		// Execute the prepared query.
		rs, err := query.Eval(context.Background(), rego.EvalInput(v))
		if err != nil {
			log.Fatal(err)
		}

		if len(rs) > 0 {
			fmt.Printf("%s %s can %s %s: %v\n", (v["subject"].(map[string]string))["role"], v["name"],
				v["action"], v["object"], rs[0].Expressions[0].Value)
		} else {
			fmt.Printf("%s %s can %s %s: %v\n", (v["subject"].(map[string]string))["role"], v["name"],
				v["action"], v["object"], false)
		}

	}
}
