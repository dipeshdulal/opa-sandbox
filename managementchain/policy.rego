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
    input.user_id = managers[_]
}