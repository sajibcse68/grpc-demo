## Run gRPC server and Proxy Server

```go
$ go run server/main.go

2018/07/17 16:52:53 gRPC Server is running at port 50051
2018/07/17 16:52:53 Proxy Server is running at port 8088

```

## Hit from browser

```js
http://localhost:8088/apis/demo/intro/json?name=Sajib%20Khan

{
  message: "Hello Sajib Khan"
}
```

## Run go clilent

```go
$ go run client/main.go

2018/07/17 17:06:20 Greeting: Hello Mr. Alice
```