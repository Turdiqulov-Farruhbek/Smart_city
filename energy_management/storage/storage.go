package storage

import (
	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
)

type StorageI interface {
	Building() BuildingI
	EnergyMeter() EnergyMeterI
	Meter() MeterI
}

type BuildingI interface {
	CreateBuilding(req *pb.BuildingCreate) (*pb.Building, error)
	UpdateBuilding(req *pb.BuildingUpdate) (*pb.Building, error)
	GetBuilding(req *pb.ById) (*pb.Building, error)
	DeleteBuilding(req *pb.ById) (*pb.Void, error)
}

type EnergyMeterI interface {
	CreateEnergyMeter(req *pb.EnergyMeterCreate) (*pb.EnergyMeter, error)
	UpdateEnergyMeter(req *pb.EnergyMeterUpdate) (*pb.EnergyMeter, error)
	GetEnergyMeter(req *pb.ById) (*pb.GetEnergyWithBuildings, error)
	DeleteEnergyMeter(req *pb.ById) (*pb.Void, error)
}

type MeterI interface {
	CreateMeterReading(req *pb.MeterReading) (*pb.Void, error)
	UpdateMeterReading(req *pb.MeterReadingUpdate) (*pb.Void, error)
	GetMeterReading(req *pb.ById) (*pb.GetMeterReadingWithEnergy, error)
	DeleteMeterReading(req *pb.ById) (*pb.Void, error)
	GetHourlyEnergyForBuilding(req *pb.ByHour) (*pb.EnergyReportBuilding, error)
	GetDailyEnergyForBuilding(req *pb.ById) (*pb.EnergyReportBuilding, error)
}
