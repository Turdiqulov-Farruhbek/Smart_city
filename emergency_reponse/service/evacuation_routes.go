package service

import (
	er "emergency_response-service/genproto/emergency_response"
	st "emergency_response-service/storage/postgres"

	"context"
)

type EvacuationRoutesService struct {
	storage st.Storage
	er.UnimplementedEvacuationRouteServiceServer
}

func NewEvacuationRoutesService(storage *st.Storage) *EvacuationRoutesService{
	return &EvacuationRoutesService{
		storage: *storage,
	}
}

func (s *EvacuationRoutesService) Create(ctx context.Context, req *er.RoutesCreateReq)(*er.RoutesRes, error){
	res, err := s.storage.RouteS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EvacuationRoutesService) GetById(ctx context.Context, id *er.ById)(*er.RoutesGetByIdRes, error){
	res, err := s.storage.RouteS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EvacuationRoutesService) GetAll(ctx context.Context, filter *er.Filter)(*er.RoutesGetAllRes, error){
	res, err := s.storage.RouteS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EvacuationRoutesService) Update(ctx context.Context, req *er.RoutesUpdateReq)(*er.RoutesRes, error){
	res, err := s.storage.RouteS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EvacuationRoutesService) Delete(ctx context.Context, id *er.ById)(*er.Void, error){
	res, err := s.storage.RouteS.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}