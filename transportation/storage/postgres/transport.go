package postgres

import (
    "database/sql"
    pb "gitlab.com/acumen/smart_city/transportation/genproto/transport"

	_ "github.com/lib/pq"
	"time"
)

type TransportRepo struct {
    db *sql.DB
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{
        db: db,
    }
}

func (r *TransportRepo) CreateVehicle(req *pb.VehicleCreate) (*pb.Vehicle, error) {
	res := &pb.Vehicle{}

    query := `INSERT INTO vehicles (vehicle_type, capacity, current_location, status)
    VALUES ($1, $2, $3) RETURNING id`

    var id string
    err := r.db.QueryRow(query, req.VehicleType, req.Capacity, req.CurrentLocation, req.Status).Scan(&id)
    if err!= nil {
        return nil, err
    }

    res, err = r.GetVehicle(&pb.ById{Id: id})
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *TransportRepo) GetVehicle(id *pb.ById) (*pb.Vehicle, error) {
	res := &pb.Vehicle{}

    query := `SELECT id, vehicle_type, capacity, current_location, status FROM vehicles WHERE id = $1`

    err := r.db.QueryRow(query, id.Id).Scan(&res.VehicleId, &res.VehicleType, &res.Capacity, &res.CurrentLocation, &res.Status)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *TransportRepo) UpdateVehicle(req *pb.VehicleUpdate) (*pb.Vehicle, error) {
	res := &pb.Vehicle{}

    query := `UPDATE vehicles SET vehicle_type = $1, capacity = $2, current_location = $3, status = $4, updated_at=now() WHERE id = $5 RETURNING id`

    var id string
    err := r.db.QueryRow(query, req.VehicleType, req.Capacity, req.CurrentLocation, req.Status, req.VehicleId).Scan(&id)
    if err!= nil {
        return nil, err
    }

    res, err = r.GetVehicle(&pb.ById{Id: id})
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *TransportRepo) DeleteVehicle(id *pb.ById) (*pb.Void, error) {
	res := &pb.Void{}

    query := `UPDATE vehicles SET deleted_at=now() WHERE id = $1`
	
	  _, err := r.db.Exec(query, id.Id)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *TransportRepo) GetLocationOfTrasnport(req *pb.ById) (*pb.Location, error) {
	res := &pb.Location{}

    query := `SELECT current_location FROM vehicles WHERE id = $1`

    err := r.db.QueryRow(query, req.Id).Scan(&res.Location)
    if err!= nil {
        return nil, err
    }
	res.Time = time.Now().Format("2006-01-02 15:04:05")

    return res, nil
}

