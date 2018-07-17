package proxy

import (
  "flag"
  "log"
  "net/http"

  "github.com/golang/glog"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "golang.org/x/net/context"
  "google.golang.org/grpc"

  gw "github.com/grpc-demo/proto"
)

var (
  echoEndPoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of YourService")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterDemoServiceHandlerFromEndpoint(ctx, mux, *echoEndPoint, opts)

  if err != nil {
    return err
  }

  return http.ListenAndServe(":8088", mux)
}

func Call() {
  flag.Parse()
  defer glog.Flush()

  log.Println("Proxy Server is running at port 8088")

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}