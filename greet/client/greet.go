package main

import (
	"context"
	"log"

	pb "github.com/IvanRoussev/grpc-course/greet/proto"
)


func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Ivan",
	})
	if err != nil {
		log.Fatalf("Could not Greet: %v\n", err)
	} 

	log.Printf("Greeting: %s\n", res.Result)
}