package main

import (
	"log"
	"net"

	pb "github.com/Akos-T/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}
	tls := true // Change it to false if you want to run without TLS
	if tls {
		certFile := "../ssl/server.crt"
		keyFile := "../ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
