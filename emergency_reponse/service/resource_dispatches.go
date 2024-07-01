package service

import (
	er "emergency_response-service/genproto/emergency_response"
	st "emergency_response-service/storage/postgres"

	"context"
)

type ResourceDispatchesService struct {
	storage st.Storage
	er.UnimplementedResourceDispatcheServiceServer
}

func NewResourceDispatchesService(storage *st.Storage) *ResourceDispatchesService{
	return &ResourceDispatchesService{
		storage: *storage,
	}
}

func (s *ResourceDispatchesService) Create(ctx context.Context, req *er.DispatchesCreateReq)(*er.DispatchesRes, error){
	res, err := s.storage.DispatchS.Create(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ResourceDispatchesService) GetById(ctx context.Context, id *er.ById)(*er.DispatchesFullRes, error){
	res, err := s.storage.DispatchS.GetById(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ResourceDispatchesService) GetAll(ctx context.Context, filter *er.Filter)(*er.DispatchesGetAllRes, error){
	res, err := s.storage.DispatchS.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ResourceDispatchesService) Update(ctx context.Context, req *er.DispatchesUpdateReq)(*er.DispatchesRes, error){
	res, err := s.storage.DispatchS.Update(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ResourceDispatchesService) Delete(ctx context.Context, id *er.ById)(*er.Void, error){
	res, err := s.storage.DispatchS.Delete(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}