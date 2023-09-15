package main

import (
	"log"

	pb "github.com/Akos-T/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	doGreetManyTimes(client)
	doLongGreet(client)
	doGreetEveryone(client)
}
