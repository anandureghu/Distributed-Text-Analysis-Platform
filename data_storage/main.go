package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// DataStorageServiceServer is the server API for Data Storage service
type DataStorageServiceServer struct {
	UnimplementedDataStorageServiceServer
}

// Function to store processed log data
func (s *DataStorageServiceServer) StoreLogData(ctx context.Context, req *StoreLogRequest) (*StoreLogResponse, error) {
	// TODO: Store the processed data in a database or file system
	return &StoreLogResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	RegisterDataStorageServiceServer(grpcServer, &DataStorageServiceServer{})
	fmt.Println("Data Storage Service is running on port 50052...")
	grpcServer.Serve(lis)
}
