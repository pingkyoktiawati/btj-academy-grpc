package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"
	"strings"

	"google.golang.org/grpc"
	pb "github.com/pingkyoktiawati/grpc-calculator/calculator/calcpb"
)

// gRPC server address
const (
	address = "localhost:50051"
)

// argParser parses command line arguments and returns two integers
func argParser(n1 string, n2 string) (int32, int32) {
	// Convert command line arguments to integers
	N1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot parse arge[1]: %s", err)
	}
	N2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Cannot parse arge[2]: %s", err)
	}
	return int32(N1), int32(N2)
}

// Main function for the client
func main() {
	// Check if two numbers are provided as command line arguments
	if len(os.Args) != 3 {
            log.Fatalf("2 numbers expected: n1 n2")
	}

	// Parse command line arguments
	n1, n2 := argParser(os.Args[1], os.Args[2])

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %s", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewCalculatorClient(conn)

	// Contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addResult, err := client.Add(ctx, &pb.AddRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		log.Fatalf("Adding error: %s", err)
	}
	log.Printf("%d + %d = %d", n1, n2, addResult.R)

	subtractResult, err := client.Subtract(ctx, &pb.SubtractRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		log.Fatalf("Subtracting error: %s", err)
	}
	log.Printf("%d - %d = %d", n1, n2, subtractResult.R)

	multiplyResult, err := client.Multiply(ctx, &pb.MultiplyRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		log.Fatalf("Multiplying error: %s", err)
	}
	log.Printf("%d * %d = %d", n1, n2, multiplyResult.R)

	divideResult, err := client.Divide(ctx, &pb.DivideRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		// Check if the error is due to division by zero
		if strings.Contains(err.Error(), "Cannot divide by zero") {
			log.Printf("Dividing error: Cannot divide by zero")
			return
		}
		log.Fatalf("Dividing error: %s", err)
	}
	log.Printf("%d / %d = %.2f", n1, n2, divideResult.R)
}