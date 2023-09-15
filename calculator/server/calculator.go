package main

import (
	"context"
	"io"
	"log"
	"math"

	pb "github.com/Akos-T/grpc-go-course/calculator/proto/calculator/v1"
)

type Server struct {
	pb.CalculatorServiceServer
}

func (s *Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Add invoked with the following input: %+v\n", req)
	sum := req.A + req.B
	return &pb.AddResponse{Result: sum}, nil
}

func (s *Server) Primes(req *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes invoked with: %+v\n", req)

	// The pseudo-algorithm was provided by the course
	// k = 2
	// N = 210
	// while N > 1:
	//     if N % k == 0:   // if k evenly divides into N
	//         print k      // this is a factor
	//         N = N / k    // divide N by k so that we have the rest of the number left.
	//     else:
	//         k = k + 1

	n := req.N
	m := int64(2)
	for n > 1 {
		if n%m == 0 {
			err := stream.Send(&pb.PrimesResponse{M: m})
			if err != nil {
				log.Printf("Error while sending message: %v\n", err)
			}

			n = n / m
		} else {
			m = m + 1
		}
	}

	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg invoked")

	var sum, counter int64 = 0, 0
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.AvgResponse{Average: float64(sum) / float64(counter)})
			}

			log.Fatalf("Error reading from stream: %v\n", err)
		}

		log.Printf("Received: %d\n", req.Number)
		sum += req.Number
		counter++
	}
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max was invoked")
	max := int64(math.MinInt64)

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			log.Fatalf("Error while reading from client stream: %v\n", err)
		}

		log.Printf("Received: %d", msg.Number)
		if max < msg.Number {
			max = msg.Number
			err = stream.Send(&pb.MaxResponse{Max: max})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
