package storage

import (
	er "emergency_response-service/genproto/emergency_response"
)

type EmergencyIncidentsI interface {
	Create(*er.IncidentsCreateReq) (*er.IncidentsRes, error)
	GetById(*er.ById) (*er.IncidentsGetByIdRes, error)
	GetAll(*er.Filter) (*er.IncidentsGetAllRes, error)
	Update(*er.IncidentsUpdateReq) (*er.IncidentsRes, error)
	Delete(*er.ById) (*er.Void, error)
	GetActiveIncidents(*er.Filter) (*er.IncidentsGetAllRes, error)
}

type EmergencyResourcesI interface {
	Create(*er.ResourcesCreateReq) (*er.ResourcesRes, error)
	GetById(*er.ById) (*er.ResourcesGetByIdRes, error)
	GetAll(*er.Filter) (*er.ResourcesGetAllRes, error)
	Update(*er.ResourcesUpdateReq) (*er.ResourcesRes, error)
	Delete(*er.ById) (*er.Void, error)
	GetAvailableResources(filter *er.FilterType)(*er.ResourcesGetAllRes, error)
}

type ResourceDispatchesI interface {
	Create(*er.DispatchesCreateReq) (*er.DispatchesRes, error)
	GetById(*er.ById) (*er.DispatchesFullRes, error)
	GetAll(*er.Filter) (*er.DispatchesGetAllRes, error)
	Update(*er.DispatchesUpdateReq) (*er.DispatchesRes, error)
	Delete(*er.ById) (*er.Void, error)
}

type EmergencyAlertsI interface {
	Create(*er.AlertsCreateReq) (*er.AlertsRes, error)
	GetById(*er.ById) (*er.AlertsGetByIdRes, error)
	GetAll(*er.Filter) (*er.AlertsGetAllRes, error)
	Update(*er.AlertsUpdateReq) (*er.AlertsRes, error)
	Delete(*er.ById) (*er.Void, error)
}

type EvacuationRoutesI interface {
	Create(*er.RoutesCreateReq) (*er.RoutesRes, error)
	GetById(*er.ById) (*er.RoutesGetByIdRes, error)
	GetAll(*er.Filter) (*er.RoutesGetAllRes, error)
	Update(*er.RoutesUpdateReq) (*er.RoutesRes, error)
	Delete(*er.ById) (*er.Void, error)
}

type EmergencyDrillsI interface {
	Create(*er.DrillsCreateReq) (*er.DrillsRes, error)
	GetById(*er.ById) (*er.DrillsGetByIdRes, error)
	GetAll(*er.Filter) (*er.DrillsGetAllRes, error)
	Update(*er.DrillsUpdateReq) (*er.DrillsRes, error)
	Delete(*er.ById) (*er.Void, error)
}

type EmergencyFacilitiesI interface {
	Create(*er.FacilitiesCreateReq) (*er.FacilitiesRes, error)
	GetById(*er.ById) (*er.FacilitiesGetByIdRes, error)
	GetAll(*er.Filter) (*er.FacilitiesGetAllRes, error)
	Update(*er.FacilitiesUpdateReq) (*er.FacilitiesRes, error)
	Delete(*er.ById) (*er.Void, error)
	GetFacilityByType(*er.FilterType)(*er.FacilitiesGetAllRes, error)
}
