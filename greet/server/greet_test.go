package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/Akos-T/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGreetEveryoneAtOnceWithServer(t *testing.T) {
	// Target param is important, it has to be "bufnet" !!!
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	tests := []struct {
		firstNames []string
		expected   string
	}{
		{
			firstNames: nil,
			expected:   "Hello ",
		},
		{
			firstNames: []string{},
			expected:   "Hello ",
		},
		{
			firstNames: []string{"Jane"},
			expected:   "Hello Jane",
		},
		{
			firstNames: []string{"Jane", "Joe"},
			expected:   "Hello Jane, Joe",
		},
	}

	for _, tt := range tests {
		req := &pb.GreetEveryoneAtOnceRequest{FirstNames: tt.firstNames}
		res, err := client.GreetEveryoneAtOnce(context.Background(), req)
		if err != nil {
			t.Errorf("GreetEveryoneAtOnce got unexpected error: %v", err)
		}
		if res == nil {
			t.Errorf("GreetEveryoneAtOnce, wanted %s, got nil result", tt.expected)
			return // To satisfy go static check for the next if
		}
		if res.Result != tt.expected {
			t.Errorf("GreetEveryoneAtOnce, wanted %s, got %s", tt.expected, res.Result)
		}
	}
}

func TestGreetEveryoneAtOnce(t *testing.T) {
	tests := []struct {
		firstNames []string
		expected   string
	}{
		{
			firstNames: nil,
			expected:   "Hello ",
		},
		{
			firstNames: []string{},
			expected:   "Hello ",
		},
		{
			firstNames: []string{"Jane"},
			expected:   "Hello Jane",
		},
		{
			firstNames: []string{"Jane", "Joe"},
			expected:   "Hello Jane, Joe",
		},
	}

	s := Server{}

	for _, tt := range tests {
		req := &pb.GreetEveryoneAtOnceRequest{FirstNames: tt.firstNames}
		res, err := s.GreetEveryoneAtOnce(context.Background(), req)
		if err != nil {
			t.Errorf("GreetEveryoneAtOnce got unexpected error: %v", err)
		}
		if res == nil {
			t.Errorf("GreetEveryoneAtOnce, wanted %s, got nil result", tt.expected)
			return // To satisfy go static check for the next if
		}
		if res.Result != tt.expected {
			t.Errorf("GreetEveryoneAtOnce, wanted %s, got %s", tt.expected, res.Result)
		}
	}
}
