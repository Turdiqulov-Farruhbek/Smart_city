package storage

import (
	"context"
	citi "citizen/genproto/citizen"
)	

type StorageI interface {
	Citizen() CitizenI
	Document() DocumentI
	Notification() NotificationI
	ServiceRequest() ServiceRequestI
}

type CitizenI interface {
	RegisterCitizen(ctx context.Context, req *citi.CitizenCreate) (*citi.Citizen, error)
	UpdateCitizenProfile(ctx context.Context, req *citi.CitizenCreate) (*citi.Citizen, error)
	GetCitizenProfile(ctx context.Context, req *citi.ById) (*citi.Citizen, error)
	DeleteCitizenProfile(ctx context.Context, req *citi.ById) (*citi.Void, error)
}


type DocumentI interface {
	UploadDocument(ctx context.Context, req *citi.DocumentCreate) (*citi.Document, error)
	GetCitizenDocuments(ctx context.Context, req *citi.ById) (*citi.DocumentList, error)	
}


type NotificationI interface {
	SetNotificationPref(ctx context.Context, req *citi.NotificationPref) (*citi.Void, error)
    UpdateNotificationPref(ctx context.Context, req *citi.NotificationPref) (*citi.Void, error)
    DeleteNotificationPref(ctx context.Context, req *citi.ById) (*citi.Void, error)
}


type ServiceRequestI interface {
	CreateServiceRequest(ctx context.Context, req *citi.ServiceReqCreate) (*citi.ServiceReq, error)
	GetCitizenRequests(ctx context.Context, req *citi.ById) (*citi.ServiceReqList, error)
	UpdateServiceRequest(ctx context.Context, req *citi.ServiceReqCreate) (*citi.ServiceReq, error)
	DeleteServiceRequest(ctx context.Context, req *citi.ById) (*citi.Void, error)
}