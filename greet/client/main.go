package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/IvanRoussev/grpc-course/greet/proto"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))


	if err != nil {
		log.Fatalf("Failed to Connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	doGreetManyTimes(c)
}