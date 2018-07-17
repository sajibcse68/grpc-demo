package main

import (
  "log"
  "net"
  
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/grpc-demo/proto"
  pr "github.com/grpc-demo/proxy"
  "google.golang.org/grpc/reflection"
)

const (
  port = ":50051"
)

type server struct{}

func (s *server) Intro(ctx context.Context, in *pb.IntroRequest) (*pb.IntroResponse, error) {
  log.Println("Client request: " + in.Name)
  return &pb.IntroResponse{Message: "Hello " + in.Name}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterDemoServiceServer(s, &server{})
  // Register reflection service on gRPC server.
  reflection.Register(s)

  log.Println("gRPC Server is running at port 50051")

  go pr.Call() // run reverse proxy server in another routine
  
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to server: %v", err)
  }
}