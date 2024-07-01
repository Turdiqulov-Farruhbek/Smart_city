package service

import (
	er "emergency_response-service/genproto/emergency_response"
	st "emergency_response-service/storage/postgres"

	"context"
)

type EmergencyDrillsService struct {
	storage st.Storage
	er.UnimplementedEmergencyDrillServiceServer
}

func NewEmergencyDrillsService(storage *st.Storage) *EmergencyDrillsService{
	return &EmergencyDrillsService{
		storage: *storage,
	}
}

func (s *EmergencyDrillsService) Create(ctx context.Context, req *er.DrillsCreateReq)(*er.DrillsRes, error){
	res, err := s.storage.DrillS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyDrillsService) GetById(ctx context.Context, id *er.ById)(*er.DrillsGetByIdRes, error){
	res, err := s.storage.DrillS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyDrillsService) GetAll(ctx context.Context, filter *er.Filter)(*er.DrillsGetAllRes, error){
	res, err := s.storage.DrillS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyDrillsService) Update(ctx context.Context, req *er.DrillsUpdateReq)(*er.DrillsRes, error){
	res, err := s.storage.DrillS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyDrillsService) Delete(ctx context.Context, id *er.ById)(*er.Void, error){
	res, err := s.storage.DrillS.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
