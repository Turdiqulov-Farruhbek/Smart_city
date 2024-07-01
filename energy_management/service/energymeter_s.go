package service

import (
	"context"

	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
	"gitlab.com/acumen931026/energy_management/storage"
)

type EnergyMeterService struct {
	storage storage.StorageI
	pb.UnimplementedEnenrgyMeterServiceServer
}

func NewEnergyMeterStorage(storage storage.StorageI) *EnergyMeterService {
	return &EnergyMeterService{
		storage: storage,
	}
}

func (s *EnergyMeterService) CreateEnergyMeter(c context.Context, req *pb.EnergyMeterCreate) (*pb.EnergyMeter, error) {
	return s.storage.EnergyMeter().CreateEnergyMeter(req)
}
func (s *EnergyMeterService) UpdateEnergyMeter(c context.Context, req *pb.EnergyMeterUpdate) (*pb.EnergyMeter, error) {
	return s.storage.EnergyMeter().UpdateEnergyMeter(req)
}
func (s *EnergyMeterService) DeleteEnergyMeter(c context.Context, id *pb.ById) (*pb.Void, error) {
	return s.storage.EnergyMeter().DeleteEnergyMeter(id)
}

func (s *EnergyMeterService) GetEnergyMeter(c context.Context, id *pb.ById) (*pb.GetEnergyWithBuildings, error) {
	return s.storage.EnergyMeter().GetEnergyMeter(id)
}
