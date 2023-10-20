package main

import (
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:6061"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	client := pb.NewBlogServiceClient(conn)

	id := createBlog(client)
	readBlog(client, id)           // valid
	readBlog(client, "NotValidID") // invalid
}
