// #################################
// go get -u google.golang.org/grpc
// go install google.golang.org/protobuf/cmd/protoc-gen-go
// #################################
package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = "9000"

func main() {
	fmt.Println("---------------------------------")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to Listen:%v\n", err)
	}
	grpcServer := grpc.NewServer()
	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Server:%s\n", err)
	}
	fmt.Println("---------------------------------")
}
