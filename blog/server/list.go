package main

import (
	"context"
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListBlog was invoked")

	cur, err := collection.Find(context.TODO(), primitive.D{{}})
	if err != nil {
		return status.Errorf(codes.Internal, "Unknown internal error: %v", err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		data := &BlogItem{}
		err = cur.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, "Error while decoding data from mongodb: %v", err)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, "Unknown internal error: %v", err)
	}

	return nil
}
