package main

import (
	"log"
	"net"

	pb "github.com/rafaelmbcosta/grpc-go-studies/greet/proto"
	"google.golang.org/grpc"
)

const addr string = "0.0.0.0:3000"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on localhost:3000")
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
