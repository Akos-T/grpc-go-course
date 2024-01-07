package main

import (
	"context"
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

	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary RPCs. To configure an interceptor
		// for streaming RPCs, see: https://godoc.org/google.golang.org/grpc#StreamInterceptor
		grpc.UnaryInterceptor(exampleInterceptor),
	}
	tls := true // Change it to false if you want to run without TLS
	if tls {
		// Enable TLS for all incoming connections.
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

func exampleInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Println(" --- This is the UNARY interceptor before running the actual operation! We can do authentication/authorization here for example and stop if necessary.")

	// Continue execution of handler
	res, err := handler(ctx, req)

	log.Println(" --- This is the UNARY interceptor after running the actual operation! We can manipulate the response here before sending it back.")

	return res, err
}
