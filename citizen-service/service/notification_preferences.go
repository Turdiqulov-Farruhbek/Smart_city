package service

import (
	"citizen/genproto/citizen"
	"citizen/storage"
	"context"
	"fmt"
	"log/slog"

	"github.com/rs/zerolog/log"
)

type NotificationService struct {
	stg storage.StorageI
	citizen.UnimplementedCitizenNotificationServiceServer
}

func NewNotificationService(Stg storage.StorageI) *NotificationService {
	if Stg == nil {
		slog.Error("storage.StorageI is nil", Stg)
	}
	return &NotificationService{
		stg: Stg,
	}
}

func (s *NotificationService) SetNotificationPref(ctx context.Context, req *citizen.NotificationPref) (*citizen.Void, error) {
	log.Info().Msg("NotificationService: SetNotificationPref called")
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }
    resp, err := s.stg.Notification().SetNotificationPref(ctx, req)
    if err!= nil {
        return nil, err
    }
    return resp, nil

}

func (s *NotificationService) UpdateNotificationPref(ctx context.Context, req *citizen.NotificationPref) (*citizen.Void, error){
	log.Info().Msg("NotificationService: UpdateNotificationPref called")
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }
    resp, err := s.stg.Notification().UpdateNotificationPref(ctx, req)
    if err!= nil {
        return nil, err
    }
    return resp, nil
}

func (s *NotificationService) DeleteNotificationPref(ctx context.Context, req *citizen.ById) (*citizen.Void, error){
	log.Info().Msg("NotificationService: DeleteNotificationPref called")
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }
    resp, err := s.stg.Notification().DeleteNotificationPref(ctx, req)
    if err!= nil {
        return nil, err
    }
    return resp, nil
}
