package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
	pb "github.com/pingkyoktiawati/grpc-calculator/proto"
)

// Define the server struct for handling gRPC requests
type notifServiceServer struct {
	pb.UnimplementedNotifServiceServer
}

// Define a struct to represent the structure of the quotes API response
type QuotesWord struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

// Define a struct to represent the structure of the foodish image API response
type FoodishImage struct {
	ImageURL string `json:"image"`
}

// Constants for API base URLs
const (
	quotesAPIBaseURL = "https://zenquotes.io/api/random";
	foodishAPIBaseURL = "https://foodish-api.com/api/"
) 

// Function to generate a random quote using the quotes API
func generateQuote() (string, error) {
	// Create an HTTP GET request
	resp, err := http.Get(quotesAPIBaseURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	var quotesWord []QuotesWord
	err = json.NewDecoder(resp.Body).Decode(&quotesWord)
	if err != nil {
		return "", err
	}

	if len(quotesWord) == 0 {
		return "", nil
	}

	quote := quotesWord[0]
	return fmt.Sprintf("%s (%s)", quote.Quote, quote.Author), nil
}

// Function to generate a random foodish image URL using the foodish API
func generateFoodishImage() (string, error) {
	// Create an HTTP GET request
	resp, err := http.Get(foodishAPIBaseURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	var foodishImage FoodishImage
	err = json.NewDecoder(resp.Body).Decode(&foodishImage)
	if err != nil {
		return "", err
	}

	return foodishImage.ImageURL, nil
}

// gRPC method to handle GetQuote requests
func (s *notifServiceServer) GetQuote(req *pb.QuoteRequest, stream pb.NotifService_GetQuoteServer) error {
	log.Printf("Received request for %d quotes", req.NumQuote)

	// Generate and send the requested number of quotes to the client
	for i := 0; i < int(req.NumQuote); i++ {
		quote, err := generateQuote()
		if err != nil {
			log.Printf("Error generating quote: %v", err)
			return err
		}

		response := &pb.QuoteResponse{
			GeneratedQuote: quote,
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to stream: %v", err)
			return err
		}

		// Simulating delay for each quote
		time.Sleep(1 * time.Second)
	}

	return nil
}

// gRPC method to handle GenerateImages requests
func (s *notifServiceServer) GenerateImages(req *pb.ImagesRequest, stream pb.NotifService_GenerateImagesServer) error {
	log.Printf("Received request for %d images", req.NumImage)

	// Generate and send the requested number of foodish image URLs to the client
	for i := 0; i < int(req.NumImage); i++ {
		imageURL, err := generateFoodishImage()
		if err != nil {
			log.Printf("Error generating foodish image: %v", err)
			return err
		}

		response := &pb.ImagesResponse{
			GeneratedImageUrls: imageURL,
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to stream: %v", err)
			return err
		}

		// Simulating delay for each image URL
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterNotifServiceServer(server, &notifServiceServer{})

	// Start gRPC server
	log.Printf("Starting gRPC server on port 50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 50051: %v", err)
	}
}