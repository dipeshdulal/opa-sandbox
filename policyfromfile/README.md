#### Policy From File

Read a policy from file and then run the query in the input.
Here, `admin with GET` method can only access the resource otherwise they cannot access it.

```go
package sample

default allow = false

allow {
    input.identity = "admin"		
    input.method = "GET"
}
```