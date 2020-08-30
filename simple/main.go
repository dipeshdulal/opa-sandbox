package main

import (
	"context"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	ctx := context.Background()

	r := rego.New(rego.Query("x = 1"))
	rs, err := r.Eval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULT: ", rs)
	fmt.Println("bindings:", rs[0].Bindings)
}
