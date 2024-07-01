package storage

import pb "gitlab.com/acumen/smart_city/transportation/genproto/transport"

type StorageI interface {
	Incident() IncidentI

}

type IncidentI interface {
	CreateIncident(*pb.IncidentCreate) (*pb.Incident, error)
	UpdateIncident(*pb.IncidentCreate) (*pb.Incident, error)
	DeleteIncident(*pb.ById) (*pb.Void, error)
	GetIncident(*pb.ById) (*pb.Void, error)
	GetAllIncedents(*pb.IncidentFilter) (*pb.IncidentList, error)
}