package service

import (
	"context"
	"fmt"

	citizen "citizen/genproto/citizen"
	"citizen/storage"
	"log/slog"

	"github.com/rs/zerolog/log"
)

type ServiceRequestService struct {
	stg storage.StorageI
	citizen.UnimplementedCitizenRequestServiceServer
}

func NewServiceRequestService(stg storage.StorageI) *ServiceRequestService {
	if stg == nil {
		slog.Error("storage.StorageI is nil", stg)
	}
	return &ServiceRequestService{
		stg: stg,
	}
}

func (s *ServiceRequestService) CreateServiceRequest(ctx context.Context, req *citizen.ServiceReqCreate) (*citizen.ServiceReq, error) {
	log.Info().Msg("ServiceRequestService: CreateServiceRequest called")
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if s.stg == nil {
		return nil, fmt.Errorf("storage is nil")
	}
	if s.stg.ServiceRequest() == nil {
		return nil, fmt.Errorf("service request storage is nil")
	}
	resp, err := s.stg.ServiceRequest().CreateServiceRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}



func (s *ServiceRequestService) GetCitizenRequests(ctx context.Context, req *citizen.ById) (*citizen.ServiceReqList, error) {
	log.Info().Msg("ServiceRequestService: GetCitizenRequests called")
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if s.stg == nil {
		return nil, fmt.Errorf("storage is nil")
	}
	if s.stg.ServiceRequest() == nil {
		return nil, fmt.Errorf("service request storage is nil")
	}
	resp, err := s.stg.ServiceRequest().GetCitizenRequests(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceRequestService) UpdateServiceRequest(ctx context.Context, req *citizen.ServiceReqCreate) (*citizen.ServiceReq, error) {
	log.Info().Msg("ServiceRequestService: UpdateServiceRequest called")
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if s.stg == nil {
		return nil, fmt.Errorf("storage is nil")
	}
	if s.stg.ServiceRequest() == nil {
		return nil, fmt.Errorf("service request storage is nil")
	}
	resp, err := s.stg.ServiceRequest().UpdateServiceRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceRequestService) DeleteServiceRequest(ctx context.Context, req *citizen.ById) (*citizen.Void, error) {
	log.Info().Msg("ServiceRequestService: DeleteServiceRequest called")
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if s.stg == nil {
		return nil, fmt.Errorf("storage is nil")
	}
	if s.stg.ServiceRequest() == nil {
		return nil, fmt.Errorf("service request storage is nil")
	}
	resp, err := s.stg.ServiceRequest().DeleteServiceRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

