package service

import (
	er "emergency_response-service/genproto/emergency_response"
	st "emergency_response-service/storage/postgres"

	"context"
)

type EmergencyFacilitiesService struct {
	storage st.Storage
	er.UnimplementedEmergencyFacilityServiceServer
}

func NewEmergencyFacilitiesService(storage *st.Storage) *EmergencyFacilitiesService{
	return &EmergencyFacilitiesService{
		storage: *storage,
	}
}

func (s *EmergencyFacilitiesService) Create(ctx context.Context, req *er.FacilitiesCreateReq)(*er.FacilitiesRes, error){
	res, err := s.storage.FacilityS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyFacilitiesService) GetById(ctx context.Context, id *er.ById)(*er.FacilitiesGetByIdRes, error){
	res, err := s.storage.FacilityS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyFacilitiesService) GetAll(ctx context.Context, filter *er.Filter)(*er.FacilitiesGetAllRes, error){
	res, err := s.storage.FacilityS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyFacilitiesService) Update(ctx context.Context, req *er.FacilitiesUpdateReq)(*er.FacilitiesRes, error){
	res, err := s.storage.FacilityS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyFacilitiesService) Delete(ctx context.Context, id *er.ById)(*er.Void, error){
	res, err := s.storage.FacilityS.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyFacilitiesService) GetFacilityByType(ctx context.Context , filter *er.FilterType)(*er.FacilitiesGetAllRes, error){
	res, err := s.storage.FacilityS.GetFacilityByType(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}
