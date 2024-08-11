package main

import (
	"context"
	"fmt"
	"net/http"
	"google.golang.org/grpc"
)

// Function to provide analytics via gRPC
func (s *AnalyticsServiceServer) GetMetrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error) {
	// TODO: Query the Data Storage Service via gRPC for the required metrics and return them.
	return &MetricsResponse{}, nil
}

// HTTP handler to provide analytics via REST
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Query the Data Storage Service via HTTP for the required metrics and return them.
	fmt.Fprintf(w, "Metrics data.")
}

func main() {
	// Start the RESTful API for Analytics Service
	go func() {
		http.HandleFunc("/metrics", metricsHandler)
		http.ListenAndServe(":8081", nil)
	}()

	// Start the gRPC server for Analytics Service
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	RegisterAnalyticsServiceServer(grpcServer, &AnalyticsServiceServer{})
	fmt.Println("Analytics Service is running on port 50053...")
	grpcServer.Serve(lis)
}
