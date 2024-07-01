package service

import (
	"context"

	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
	"gitlab.com/acumen931026/energy_management/storage"
)

type MeterReadingService struct {
	storage storage.StorageI
	pb.UnimplementedMeterReadingServiceServer
}

func NewMeterReadingrStorage(storage storage.StorageI) *MeterReadingService {
	return &MeterReadingService{
		storage: storage,
	}
}

func (s *MeterReadingService) CreateMeterReading(c context.Context, req *pb.MeterReading) (*pb.Void, error) {
	return s.storage.Meter().CreateMeterReading(req)
}
func (s *MeterReadingService) UpdateMeterReading(c context.Context, req *pb.MeterReadingUpdate) (*pb.Void, error) {
	return s.storage.Meter().UpdateMeterReading(req)
}
func (s *MeterReadingService) DeleteMeterReading(c context.Context, id *pb.ById) (*pb.Void, error) {
	return s.storage.Meter().DeleteMeterReading(id)
}

func (s *MeterReadingService) GetMeterReading(c context.Context, id *pb.ById) (*pb.GetMeterReadingWithEnergy, error) {
	return s.storage.Meter().GetMeterReading(id)
}
func (s *MeterReadingService) GetHourlyEnergyForBuilding(c context.Context, id *pb.ByHour) (*pb.EnergyReportBuilding, error) {
	return s.storage.Meter().GetHourlyEnergyForBuilding(id)
}
func (s *MeterReadingService) GetDailyEnergyForBuilding(c context.Context, id *pb.ById) (*pb.EnergyReportBuilding, error) {
	return s.storage.Meter().GetDailyEnergyForBuilding(id)
}

