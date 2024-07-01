package service

import (
	"citizen/genproto/citizen"
	"citizen/storage"
	"context"
	"fmt"
	"log/slog"

	"github.com/rs/zerolog/log"
)

type DocumentsService struct {
	stg storage.StorageI
	citizen.UnimplementedCitizenDocumentServiceServer
}

func NewDocumentsService(Stg storage.StorageI) *DocumentsService {
	if Stg == nil {
		slog.Error("storage.StorageI is nil", Stg)
	}
	return &DocumentsService{
		stg: Stg,
	}
}



func (s *DocumentsService) UploadDocument(ctx context.Context, req *citizen.DocumentCreate) (*citizen.Document, error) {
	log.Info().Msg("DocumentsService: UploadDocument called")
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }
    resp, err := s.stg.Document().UploadDocument(ctx, req)
    if err!= nil {
        return nil, err
    }
    return resp, nil
}

func (s *DocumentsService) GetCitizenDocuments(ctx context.Context, req *citizen.ById) (*citizen.DocumentList, error) {
	log.Info().Msg("DocumentsService: GetCitizenDocuments called")
    if req == nil {
        return nil, fmt.Errorf("request is nil")
    }
    if s.stg == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if s.stg.Citizen() == nil {
        return nil, fmt.Errorf("citizen storage is nil")
    }
    resp, err := s.stg.Document().GetCitizenDocuments(ctx, req)
    if err!= nil {
        return nil, err
    }
    return resp, nil
}