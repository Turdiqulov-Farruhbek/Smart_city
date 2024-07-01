package postgres

import (
    "database/sql"
    pb "gitlab.com/acumen/smart_city/transportation/genproto/transport"

	_ "github.com/lib/pq"
)

type RouteRepo struct {
    db *sql.DB
}

func NewRouteRepo(db *sql.DB) *RouteRepo {
	return &RouteRepo{
        db: db,
    }
}

func (r *RouteRepo) CreateTransportRoute(req *pb.RouteCreate) (*pb.Route, error) {
	res := &pb.Route{}

	query := `INSERT INTO routes (route_name, status, route_type, start_point, end_point)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var id string
	err := r.db.QueryRow(query, req.RouteName, req.Status, req.RouteType, req.StartPoint, req.EndPoint).Scan(&id)
	if err!= nil {
        return nil, err
    }

	res, err = r.GetTransportRoute(&pb.ById{Id: id})
	if err!= nil {
        return nil, err
    }

	return res, nil
}


func (r *RouteRepo) GetTransportRoute(id *pb.ById) (*pb.Route, error) {
	res := &pb.Route{}

    query := `SELECT id, route_name, route_type, start_point, end_point FROM routes WHERE id = $1`

    err := r.db.QueryRow(query, id.Id).Scan(&res.RouteId, &res.RouteName, &res.RouteType, &res.StartPoint, &res.EndPoint)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *RouteRepo) GetAllTrasnportRoutes(req *pb.RouteFilter) (*pb.RouteList, error) {
	res := &pb.RouteList{}

    query := `SELECT id, route_name, route_type, start_point, end_point FROM routes where deleted_at is null`

    rows, err := r.db.Query(query)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        route := &pb.Route{}

        err := rows.Scan(&route.RouteId, &route.RouteName, &route.RouteType, &route.StartPoint, &route.EndPoint)
        if err!= nil {
            return nil, err
        }

        res.Routes = append(res.Routes, route)
    }

    return res, nil
}

func (r *RouteRepo) CreateSchedule(req *pb.RouteScheduleCreate) (*pb.RouteSchedule, error) {
	res := &pb.RouteSchedule{}

	query := `INSERT INTO route_schedules (route_id, departure_time, arrival_time, day_of_week)
	VALUES ($1, $2, $3, $4) RETURNING id`

	var id string
	err := r.db.QueryRow(query, req.RouteId, req.DepartureTime, req.ArrivalTime, req.DayOfWeek).Scan(&id)
	if err!= nil {
        return nil, err
    }

	res, err = r.GetSchedule(&pb.ById{Id: id})
	if err!= nil {
        return nil, err
    }

	return res, nil
}

func (r *RouteRepo) GetSchedule(id *pb.ById) (*pb.RouteSchedule, error) {
	res := &pb.RouteSchedule{}

    query := `SELECT id, route_id, departure_time, arrival_time, day_of_week FROM route_schedules WHERE id = $1`

    err := r.db.QueryRow(query, id.Id).Scan(&res.ScheduleId, &res.RouteId, &res.DepartureTime, &res.ArrivalTime, &res.DayOfWeek)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *RouteRepo) GetScheduleForRoute(req *pb.ById) (*pb.RouteScheduleList, error) {
	res := &pb.RouteScheduleList{}

    query := `SELECT id, route_id, departure_time, arrival_time, day_of_week FROM route_schedules WHERE route_id = $1`

    rows, err := r.db.Query(query, req.Id)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        schedule := &pb.RouteSchedule{}

        err := rows.Scan(&schedule.ScheduleId, &schedule.RouteId, &schedule.DepartureTime, &schedule.ArrivalTime, &schedule.DayOfWeek)
        if err!= nil {
            return nil, err
        }

        res.Schedules = append(res.Schedules, schedule)
    }

    return res, nil
}