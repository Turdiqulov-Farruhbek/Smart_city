package service

import (
	"context"

	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
	"gitlab.com/acumen931026/energy_management/storage"
)

type BuildingsStorage struct {
	storage storage.StorageI
	pb.UnimplementedBuildingServiceServer
}

func NewBuildingsStorage(storage storage.StorageI) *BuildingsStorage {
	return &BuildingsStorage{
		storage: storage,
	}
}

func (s *BuildingsStorage) CreateBuilding(c context.Context, req *pb.BuildingCreate) (*pb.Building, error) {
	return s.storage.Building().CreateBuilding(req)
}
func (s *BuildingsStorage) UpdateBuilding(c context.Context, req *pb.BuildingUpdate) (*pb.Building, error) {
	return s.storage.Building().UpdateBuilding(req)
}
func (s *BuildingsStorage) DeleteBuilding(c context.Context, id *pb.ById) (*pb.Void, error) {
	return s.storage.Building().DeleteBuilding(id)
}

func (s *BuildingsStorage) GetBuilding(c context.Context, id *pb.ById) (*pb.Building, error) {
	return s.storage.Building().GetBuilding(id)
}
