# GoUrl
this is a simple url task awaiter written in go

this project is itself a server which will await tasks in threads for you
Task is just url to fetch, method to use and body to send
```json
{
    "url":"string",
    "method":"string",
    "body":"string"
}
```

## Usage
first start server by running `go run main.go`
or build and it `go build main.go`

then you can send Tasks to `http://localhost:8080`
and every task will be awaited and retranslated back to you
