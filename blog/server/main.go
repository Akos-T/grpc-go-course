package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Akos-T/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const address = "0.0.0.0:6061"

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Testing the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	log.Printf("gRPC server is listening on %s", address)

	if err = s.Serve(l); err != nil {
		log.Fatalf("Couldn't start the server: %v\n", err)
	}
}
