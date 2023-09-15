package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Akos-T/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, input *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with: %v\n", input)
	return &pb.GreetResponse{
		Result: "Hello " + input.FirstName,
	}, nil
}

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", req)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", req.FirstName, i+1)

		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.GreetResponse{Result: res})
			}

			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", msg)
		res += fmt.Sprintf("Hello %s!\n", msg.FirstName)
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
