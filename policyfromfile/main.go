package main

import (
	"context"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	ctx := context.Background()

	r := rego.New(
		rego.Query("data.sample.allow"),
		rego.Load([]string{"./policy.rego"}, nil),
	)

	ps, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	input := map[string]interface{}{
		"identity": "admin",
		"method":   "GET",
	}
	rs, err := ps.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("INPUT: %+v -> ALLOW: %v \n", input, rs[0].Expressions)

	input = map[string]interface{}{
		"identity": "bob",
		"method":   "GET",
	}

	rs, err = ps.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("INPUT: %+v -> ALLOW: %v \n", input, rs[0].Expressions)

}
