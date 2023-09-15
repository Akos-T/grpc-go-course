package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Akos-T/grpc-go-course/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{FirstName: "Joe"})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{FirstName: "Joe"}
	stream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Printf("Error while sending LongGreet request: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}

func doGreetEveryone(client pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	waitChan := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Printf("Error while sending request on stream: %v\n", err)
			}
			time.Sleep(1 * time.Second)
		}

		err = stream.CloseSend()
		if err != nil {
			log.Printf("Error while closing stream: %v\n", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %s\n", res.Result)
		}

		close(waitChan)
	}()

	<-waitChan
}
