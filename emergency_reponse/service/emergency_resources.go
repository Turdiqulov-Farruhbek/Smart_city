package service

import (
	er "emergency_response-service/genproto/emergency_response"
	st "emergency_response-service/storage/postgres"

	"context"
)

type EmergencyResourcesService struct {
	storage st.Storage
	er.UnimplementedEmergencyResourceServiceServer
}

func NewEmergencyResourcesService(storage *st.Storage) *EmergencyResourcesService{
	return &EmergencyResourcesService{
		storage: *storage,
	}
}

func (s *EmergencyResourcesService) Create(ctx context.Context, req *er.ResourcesCreateReq)(*er.ResourcesRes, error){
	res, err := s.storage.ResourceS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyResourcesService) GetById(ctx context.Context, id *er.ById)(*er.ResourcesGetByIdRes, error){
	res, err := s.storage.ResourceS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyResourcesService) GetAll(ctx context.Context, filter *er.Filter)(*er.ResourcesGetAllRes, error){
	res, err := s.storage.ResourceS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyResourcesService) Update(ctx context.Context, req *er.ResourcesUpdateReq)(*er.ResourcesRes, error){
	res, err := s.storage.ResourceS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyResourcesService) Delete(ctx context.Context, id *er.ById)(*er.Void, error){
	res, err := s.storage.ResourceS.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyResourcesService) GetAvailableResources(ctx context.Context, filter *er.FilterType)(*er.ResourcesGetAllRes, error){
	res, err := s.storage.ResourceS.GetAvailableResources(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}