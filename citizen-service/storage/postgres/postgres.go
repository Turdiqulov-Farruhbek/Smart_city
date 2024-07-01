package postgres

import (
	"citizen/config"
	"context"
	"fmt"
	"log/slog"
	"citizen/storage"

	"github.com/jackc/pgx/v5"
)

type StorageStruct struct {
	DB *pgx.Conn
	Citizen_S storage.CitizenI
	Document_S storage.DocumentI
	Notification_S storage.NotificationI
	ServiceRequest_S storage.ServiceRequestI
}

func DbCon() (*StorageStruct, error) {
	var (
		db  *pgx.Conn
		err error
	)
	cfg := config.Load()
	dbcon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err = pgx.Connect(context.Background(), dbcon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
		return nil, err
	}

	err = db.Ping(context.Background())
	if err!= nil {
        slog.Warn("Unable to ping database:", err)
        return nil, err
    }

	return &StorageStruct{
        DB: db,
    }, nil
}

func (s *StorageStruct) Citizen() storage.CitizenI{
	if s.Citizen_S == nil {
		s.Citizen_S = NewCitizen(s.DB)
	}

    return s.Citizen_S
}


func (s *StorageStruct) Document() storage.DocumentI{
	if s.Document_S == nil {
        s.Document_S = NewDocument(s.DB)
    }

	return s.Document_S
}

func (s *StorageStruct) Notification() storage.NotificationI{
	if s.Notification_S == nil {
        s.Notification_S = NewNotification(s.DB)
    }

	return s.Notification_S
}

func (s *StorageStruct) ServiceRequest() storage.ServiceRequestI{
	if s.ServiceRequest_S == nil {
        s.ServiceRequest_S = NewServiceRequest(s.DB)
    }

	return s.ServiceRequest_S
}