package main

import (
	"fmt"
	"gateway/api"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up gRPC connections
	authservice, err := grpc.NewClient("localhost:2021", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		slog.Error("Failed to connect to authservice service %v", err)
	}
	defer authservice.Close()

	citizenService, err := grpc.NewClient("localhost:2023", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
        slog.Error("Failed to connect to citizenservice service %v", err)
    }


	router := api.NewGin(citizenService)
	fmt.Println("API Gateway running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		slog.Error("Failed to connect to gin engine: %v", err)
	}

}
