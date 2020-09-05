package main

import (
	"context"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/util"
)

var managementChain = `
{
	"management_chain": {
	  "bob": ["jen", "janet"],
	  "alice": ["janet"]
	}
  }
  
`

func main() {
	ctx := context.Background()

	var data map[string]interface{}

	if err := util.UnmarshalJSON([]byte(managementChain), &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("GOT DATA: ", data)

	store := inmem.NewFromObject(data)

	txn, err := store.NewTransaction(ctx, storage.WriteParams)
	if err != nil {
		log.Fatal(err)
	}

	r := rego.New(
		rego.Query(`data.management.allow`),
		rego.Store(store),
		rego.Transaction(txn),
		rego.Load([]string{"./policy.rego"}, nil),
	)

	ps, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	whoPrompt := promptui.Prompt{
		Label: "Who are you ?",
	}

	me, err := whoPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	queryPrompt := promptui.Prompt{
		Label: "Whoose salary to view ?",
	}

	path, err := queryPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	input := map[string]interface{}{
		"method":  "GET",
		"path":    []string{"salary", path},
		"user_id": me,
	}

	rs, err := ps.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("INPUT: %+v => ALLOW: %v, \n", input, rs[0].Expressions[0])
}
