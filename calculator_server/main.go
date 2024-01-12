package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "github.com/pingkyoktiawati/grpc-calculator/calculator/calcpb"
)

// Server configuration
var (
	port = flag.Int("port", 50051, "The server port")
)

// Server is used to implement the calc.CalculatorServer
type server struct{
	pb.UnimplementedCalculatorServer
}

// Implementation of the Add gRPC method for addition
func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received request: %d + %d\n", in.N1, in.N2)
	return &pb.AddResponse{R: in.N1 + in.N2}, nil
}

// Implementation of the Subtract gRPC method for subtraction
func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	log.Printf("Received request: %d - %d\n", in.N1, in.N2)
	return &pb.SubtractResponse{R: in.N1 - in.N2}, nil
}

// Implementation of the Multiply gRPC method for multiplication
func (s *server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	log.Printf("Received request: %d * %d\n", in.N1, in.N2)
	return &pb.MultiplyResponse{R: in.N1 * in.N2}, nil
}

// Implementation of the Divide gRPC method for division
func (s *server) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideResponse, error) {
	log.Printf("Received request: %d / %d\n", in.N1, in.N2)
    // Check for division by zero
	if in.N2 == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Cannot divide by zero")
    }
	result := float32(in.N1) / float32(in.N2)
    return &pb.DivideResponse{R: result}, nil
}

// Main function where the server is configured and started
func main() {
	// Create a listener on the specified protocol and port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the Calculator server implementation
	pb.RegisterCalculatorServer(s, &server{})

	// Print a log message indicating the server is listening at the specified address
	log.Printf("server listening at %v", lis.Addr())
	
	// Start serving the gRPC requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}