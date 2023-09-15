package main

import (
	"log"
	"net"

	pb "github.com/Akos-T/grpc-go-course/calculator/proto/calculator/v1"
	"google.golang.org/grpc"
)

const address = "0.0.0.0:6061"

func main() {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	log.Printf("gRPC server is listening on %s", address)

	if err = s.Serve(l); err != nil {
		log.Fatalf("Couldn't start the server: %v\n", err)
	}
}
