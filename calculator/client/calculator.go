package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Akos-T/grpc-go-course/calculator/proto/calculator/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doCalculation(client pb.CalculatorServiceClient) {
	var a, b, expected int64 = 3, 5, 8
	log.Println("doCalculation was invoked")
	log.Printf("%d + %d => We expect %d\n", a, b, expected)

	res, err := client.Add(context.Background(), &pb.AddRequest{A: a, B: b})
	if err != nil {
		log.Fatalf("Failed to add 2 numbers: %v\n", err)
	}

	log.Printf("%d + %d = %d \u2705", a, b, res.Result)
}

func doPrimes(client pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	n := int64(120)
	stream, err := client.Primes(context.Background(), &pb.PrimesRequest{N: n})
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	log.Printf("Prime number decomposition of %d:\n", n)
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error while reading the stram: %v\n", err)
		}

		log.Println(msg.M)
	}
}

func doAvg(client pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error calling Avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Printf("Error while sending Avg request: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Average)
}

func doMax(client pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while getting stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: -2},
		{Number: 8},
	}

	go func() {
		for _, req := range reqs {
			log.Printf("Sending: %v\n", req)
			err = stream.Send(req)
			if err != nil {
				log.Printf("Error while sending request on stream: %v\n", err)
			}
			time.Sleep(2 * time.Second)
		}

		err = stream.CloseSend()
		if err != nil {
			log.Printf("Error while closing stream: %v\n", err)
		}
	}()

	waitChan := make(chan bool)
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				log.Printf("Error while receiving message: %v\n", err)
				break
			}

			log.Printf("Max: %d\n", msg.Max)
		}

		close(waitChan)
	}()

	<-waitChan
}

func doSqrt(client pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		}

		log.Fatalf("A non gRPC error: %v\n", err)
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
