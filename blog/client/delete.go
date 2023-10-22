package main

import (
	"context"
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- deleteBlog was invoked ---")

	_, err := c.DeleteBlog(context.TODO(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v", err)
	}

	log.Println("Blog was deleted!")
}
