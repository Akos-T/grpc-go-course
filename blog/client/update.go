package main

import (
	"context"
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--- updateBlog was invoked ---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Clement",
		Title:    "A new Title",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.TODO(), newBlog)
	if err != nil {
		log.Fatalf("Error happend while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
