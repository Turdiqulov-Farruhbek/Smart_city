package service

import (
	er "emergency_response-service/genproto/emergency_response"
	st "emergency_response-service/storage/postgres"

	"context"
)

type EmergencyAlertsService struct {
	storage st.Storage
	er.UnimplementedAlertServiceServer
}

func NewEmergencyAlertsService(storage *st.Storage) *EmergencyAlertsService{
	return &EmergencyAlertsService{
		storage: *storage,
	}
}

func (s *EmergencyAlertsService) Create(ctx context.Context, req *er.AlertsCreateReq)(*er.AlertsRes, error){
	res, err := s.storage.AlertS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyAlertsService) GetById(ctx context.Context, id *er.ById)(*er.AlertsGetByIdRes, error){
	res, err := s.storage.AlertS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyAlertsService) GetAll(ctx context.Context, filter *er.Filter)(*er.AlertsGetAllRes, error){
	res, err := s.storage.AlertS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyAlertsService) Update(ctx context.Context, req *er.AlertsUpdateReq)(*er.AlertsRes, error){
	res, err := s.storage.AlertS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmergencyAlertsService) Delete(ctx context.Context, id *er.ById)(*er.Void, error){
	res, err := s.storage.AlertS.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
