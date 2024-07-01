package service

import (
	er "emergency_response-service/genproto/emergency_response"
	st "emergency_response-service/storage/postgres"

	"context"
)

type EmergencyIncidentsService struct {
	storage st.Storage
	er.UnimplementedEmergencyIncidentServiceServer
}

func NewEmergencyIncidentsService(storage *st.Storage) *EmergencyIncidentsService{
	return &EmergencyIncidentsService{
		storage: *storage,
	}
}

func (s *EmergencyIncidentsService) Create(ctx context.Context, req *er.IncidentsCreateReq)(*er.IncidentsRes, error){
	res, err := s.storage.IncidentS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyIncidentsService) GetById(ctx context.Context, id *er.ById)(*er.IncidentsGetByIdRes, error){
	res, err := s.storage.IncidentS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyIncidentsService) GetAll(ctx context.Context, filter *er.Filter)(*er.IncidentsGetAllRes, error){
	res, err := s.storage.IncidentS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyIncidentsService) Update(ctx context.Context, req *er.IncidentsUpdateReq)(*er.IncidentsRes, error){
	res, err := s.storage.IncidentS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyIncidentsService) Delete(ctx context.Context, id *er.ById)(*er.Void, error){
	res, err := s.storage.IncidentS.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyIncidentsService) GetActiveIncidents(ctx context.Context, filter *er.Filter)(*er.IncidentsGetAllRes, error){
	res, err := s.storage.IncidentS.GetActiveIncidents(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}