package main

import (
	"log"
	"net"

	pb "github.com/Akos-T/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

const (
	addr = "0.0.0.0:5051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v'n", err)
	}

	log.Printf("Listening on %s'n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
