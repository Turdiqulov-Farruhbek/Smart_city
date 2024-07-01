package main

import (
	lg "citizen/config"
	"citizen/genproto/citizen"
	"citizen/service"
	"citizen/storage/postgres"
	"log"

	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.DbCon()
	if err != nil {
		panic(err)
	}

	cfg := lg.Load()
	liss, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	citizen.RegisterCitizenServiceServer(server, service.NewCitizenService(db))
	citizen.RegisterCitizenDocumentServiceServer(server, service.NewDocumentsService(db))
	citizen.RegisterCitizenNotificationServiceServer(server, service.NewNotificationService(db))
	citizen.RegisterCitizenRequestServiceServer(server, service.NewServiceRequestService(db))

	log.Println("Server is running on port :2123")
	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}

}
