package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	ctx := context.Background()

	raw := `{"users": [{"id": "bob"}, {"id": "alice"}]}`
	d := json.NewDecoder(bytes.NewBufferString(raw))

	d.UseNumber()

	var input interface{}

	if err := d.Decode(&input); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Input is ", input)

	r := rego.New(
		rego.Query("input.users[idx].id = user_id"),
		rego.Input(input),
	)

	rs, err := r.Eval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v \n", rs)
	fmt.Println("len:", len(rs))
	fmt.Println("bindings.idx:", rs[1].Bindings["idx"])
	fmt.Println("bindings.user_id:", rs[1].Bindings["user_id"])

}
