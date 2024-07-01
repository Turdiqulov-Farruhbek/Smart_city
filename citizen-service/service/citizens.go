package service

import (
	"citizen/genproto/citizen"
	"citizen/storage"
	"context"
	"fmt"
	"log/slog"

	"github.com/rs/zerolog/log"
)

type CitizensService struct {
	stg storage.StorageI
	citizen.UnimplementedCitizenServiceServer
}

func NewCitizenService(Stg storage.StorageI) *CitizensService {
	if Stg == nil {
		slog.Error("storage.StorageI is nil", Stg)
	}
	return &CitizensService{
		stg: Stg,
	}
}

func (s *CitizensService) RegisterCitizen(ctx context.Context, req *citizen.CitizenCreate) (*citizen.Citizen, error) {
	log.Info().Msg("CitizenService: CreateCitizen called")

	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if s.stg == nil {
		return nil, fmt.Errorf("storage is nil")
	}
	if s.stg.Citizen() == nil {
		return nil, fmt.Errorf("citizen storage is nil")
	}

	resp, err := s.stg.Citizen().RegisterCitizen(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (s *CitizensService) GetCitizenProfile(ctx context.Context, req *citizen.ById) (*citizen.Citizen, error) {
    log.Info().Msg("CitizenService: GetCitizenProfile called")
	if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
	if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
	if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }
	resp, err := s.stg.Citizen().GetCitizenProfile(ctx, req)
	if err!= nil {
        return nil, err
    }
	return resp, nil
}



func (s *CitizensService) UpdateCitizenProfile(ctx context.Context, req *citizen.CitizenCreate) (*citizen.Citizen, error) {
	log.Info().Msg("CitizenService: UpdateCitizenProfile called")
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }

    resp, err := s.stg.Citizen().UpdateCitizenProfile(ctx, req)
    if err!= nil {
        return nil, err
    }

    return resp, nil
}



func (s *CitizensService) DeleteCitizenProfile(ctx context.Context, req *citizen.ById) (*citizen.Void, error) {
	log.Info().Msg("CitizenService: DeleteCitizenProfile called")
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }

    resp, err := s.stg.Citizen().DeleteCitizenProfile(ctx, req)
    if err!= nil {
        return nil, err
    }

    return resp, nil
}