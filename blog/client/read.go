package main

import (
	"context"
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--- readBlog was invoked ---")

	req := &pb.BlogId{Id: id}
	blog, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}
	log.Printf("Blog was read: %v\n", blog)

	return blog
}
