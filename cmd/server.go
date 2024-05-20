package main

//go:generate protoc --go_out=../build --go_opt=paths=source_relative --go-grpc_out=../build --go-grpc_opt=paths=source_relative ../proto_specification/hello-world.proto

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    proto "demoneno4ec/go-grpc-template/proto_specification"
)

// !TODO port to config
var (
    port = flag.Int("port", 50051, "The server port")
)

type server struct {
    proto.UnimplementedHelloWorldGoServiceServer
}

func (s *server) Hello(ctx context.Context, request *proto.HelooRequest) (*proto.WorldGoResponse, error) {
    query := request.GetQuery()
    log.Printf("Received: %v", query)
    return &proto.WorldGoResponse{
        Message: "Hello " + query,
    }, nil
}

func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    serv := grpc.NewServer()
    proto.RegisterHelloWorldGoServiceServer(serv, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := serv.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
