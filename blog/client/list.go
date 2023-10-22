package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("--- listBlog was invoked ---")

	stream, err := c.ListBlog(context.TODO(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Something happened: %v", err)
		}

		log.Println(res)
	}
}
