package main

import (
	"log"
	"net"
	"gitlab.com/acumen931026/energy_management/config"
	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
	"gitlab.com/acumen931026/energy_management/service"
	"gitlab.com/acumen931026/energy_management/storage/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	cfg := config.Load()
	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Println("error while connecting to postgres: ", err)
	}

	liss, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBuildingServiceServer(s, service.NewBuildingsStorage(db))
	pb.RegisterEnenrgyMeterServiceServer(s, service.NewEnergyMeterStorage(db))
	pb.RegisterMeterReadingServiceServer(s, service.NewMeterReadingrStorage(db))

	reflection.Register(s)
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
