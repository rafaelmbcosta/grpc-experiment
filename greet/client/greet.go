package main

import (
	"context"
	"log"

	pb "github.com/rafaelmbcosta/grpc-go-studies/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Rafael",
	})

	if err != nil {
		log.Fatalf("Deu ruim no do Greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
