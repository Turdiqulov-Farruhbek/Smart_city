package postgres

import (
	"database/sql"
	pb "gitlab.com/acumen/smart_city/transportation/genproto/transport"

	_ "github.com/lib/pq"
)

type IncidentsRepo struct {
	db *sql.DB
}

func NewIncidentsRepo(db *sql.DB) *IncidentsRepo {
	return &IncidentsRepo{
		db: db,
	}
}

func (r *IncidentsRepo) CreateIncident(req *pb.IncidentCreate) (*pb.Incident, error) {
    res := &pb.Incident{}

	query := `INSERT INTO incidents (road_id, description, status, reported_at)
				VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var id string

	err := r.db.QueryRow(query, req.RoadId, req.Description, req.String(), req.Time).Scan(&id)
	if err!= nil {
        return nil, err
    }

	res, err = r.GetIncident(&pb.ById{Id: id})
	if err!= nil {
        return nil, err
    }

	return res, nil
}

func (r *IncidentsRepo) GetIncident(id *pb.ById) (*pb.Incident, error) {
	res := &pb.Incident{}

    query := `SELECT id, road_id, description, status, reported_at FROM incidents WHERE id = $1`

    err := r.db.QueryRow(query, id.Id).Scan(&res.IncidentId, &res.RoadId, &res.Description, &res.Status, &res.Time)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func (r *IncidentsRepo) GetAllIncedents(req *pb.IncidentFilter) (*pb.IncidentList, error) {
	res := &pb.IncidentList{}

    query := `SELECT id, road_id, description, status, reported_at FROM incidents where deleted_at is null`

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        incident := &pb.Incident{}

        err := rows.Scan(&incident.IncidentId, &incident.RoadId, &incident.Description, &incident.Status, &incident.Time)
        if err != nil {
            return nil, err
        }

        res.Incidents = append(res.Incidents, incident)
    }

    return res, nil
}

func (r *IncidentsRepo) UpdateIncident(req *pb.IncidentUpdate) (*pb.Incident, error) {
	res := &pb.Incident{}

    query := `UPDATE incidents SET road_id = $1, description = $2, status = $3, reported_at = $4, updated_at=now() WHERE id = $5 RETURNING id`

    var id string

    err := r.db.QueryRow(query, req.RoadId, req.Description, req.Status, req.Time, req.IncidentId).Scan(&id)
    if err!= nil {
        return nil, err
    }

    res, err = r.GetIncident(&pb.ById{Id: id})
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *IncidentsRepo) DeleteIncident(id *pb.ById) (*pb.Void, error) {
	res := &pb.Void{}

    query := `UPDATE incidents SET deleted_at=EXTRACT(EPOCH FROM now()) WHERE id = $1`

    _, err := r.db.Exec(query, id.Id)
    if err!= nil {
        return nil, err
    }

    return res, nil
}