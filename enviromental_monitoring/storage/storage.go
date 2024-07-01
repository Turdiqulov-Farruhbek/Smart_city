package storage

import 	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"

type StorageI interface {
	AirStation() AirStationI
}
type AirStationI interface {
	CreateAirStation(req *pb.Station) error
	UpdateAirStation(req *pb.Station) error
	DeleteStation(id *pb.ById) error
	GetStation(id *pb.ById) (*pb.Station, error)
}