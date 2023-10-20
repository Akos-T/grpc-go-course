package main

import (
	"context"
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, id *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with: %v\n", id)

	oid, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse ID: %v", err)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog with the ID provided: ID=%s", id.Id)
	}

	return documentToBlog(data), nil
}
