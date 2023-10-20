package main

import (
	"context"
	"log"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, blog *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with: %v\n", blog)

	oid, err := primitive.ObjectIDFromHex(blog.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse ID: %v", err)
	}

	data := &BlogItem{
		AuthorID: blog.AuthorId,
		Title:    blog.Title,
		Content:  blog.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update: %v", err)
	}
	if res.MatchedCount < 1 {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog with ID=%s, err=%v", blog.Id, err)
	}

	return &emptypb.Empty{}, nil
}
