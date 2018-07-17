package main
import (
  "log"
  "os"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/grpc-demo/proto"
)

const (
  address     = "localhost:50051"
  defaultName = "Mr. Alice"
)

func main() {
  // Set up a connection to the server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewDemoServiceClient(conn)

  // Contact the server and print out its response.
  name := defaultName
  if len(os.Args) > 1 {
    name = os.Args[1]
  }

  r, err := c.Intro(context.Background(), &pb.IntroRequest{Name: name})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Greeting: %s", r.Message)
}
