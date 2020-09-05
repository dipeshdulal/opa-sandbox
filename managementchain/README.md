### Management Chain Example with `rego` and `openpolicyagent`


#### Requirement

- Current user can access theirs information as well.
- Jen and Janet can access bob's information.
- Janet can access alice's information.

#### Input

```json
{
  "management_chain": {
    "bob": ["jen", "janet"],
    "alice": ["janet"]
  }
}
```

#### Policy

```ruby
package management

default allow = false

# Allow when ourself want to access
# our data
allow {
    # allow when method is GET
    input.method = "GET"
    # allow when path matches ["salary", "bob"] array
    input.path = ["salary", id]
    # allow when input user_id is path's id
    input.user_id = id
}

allow {
    input.method = "GET"

    input.path = ["salary", id]

    # managers variable result is data.management_chain[input_id from path]
    managers = data.management_chain[id]

    # allow if user_id is one of the managers in array
    # _ part is saying rego to go through array one by one
    input.user_id = managers[_]
}
```

#### Loading REGO in library

We need to create in-memory store and transaction to load input data into rego library as below;

```go
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
```

#### Evaluate and view output

```go
rs, err := ps.Eval(ctx, rego.EvalInput(input))
if err != nil {
    log.Fatal(err)
}

fmt.Printf("INPUT: %+v => ALLOW: %v, \n", input, rs[0].Expressions[0])
```
