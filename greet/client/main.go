package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/rafaelmbcosta/grpc-go-studies/greet/proto"
)

const addr string = "localhost:3000"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect")
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
}
