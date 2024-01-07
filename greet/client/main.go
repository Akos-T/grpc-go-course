package main

import (
	"log"
	"time"

	pb "github.com/Akos-T/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const addr = "localhost:5051"

func main() {
	opts := []grpc.DialOption{}
	tls := true // Change it to false if you want to run without TLS
	if tls {
		certFile := "../ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	doGreetManyTimes(client)
	doLongGreet(client)
	doGreetEveryone(client)
	doGreetWithDeadline(client, time.Second*5)
	doGreetWithDeadline(client, time.Second*1)
	doGreetEveryoneAtOnce(client)
}
