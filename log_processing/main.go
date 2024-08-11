package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// LogProcessingServiceServer is the server API for Log Processing service
type LogProcessingServiceServer struct {
	UnimplementedLogProcessingServiceServer
}

// Function to process log data and extract metrics
func (s *LogProcessingServiceServer) ProcessLogData(ctx context.Context, req *ProcessLogRequest) (*ProcessLogResponse, error) {
	// TODO: Parse the log data, extract metrics, and send them to the Data Storage Service via gRPC.
	return &ProcessLogResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	RegisterLogProcessingServiceServer(grpcServer, &LogProcessingServiceServer{})
	fmt.Println("Log Processing Service is running on port 50051...")
	grpcServer.Serve(lis)
}
