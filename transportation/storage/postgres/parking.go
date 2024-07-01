package postgres

import (
    "database/sql"
    pb "gitlab.com/acumen/smart_city/transportation/genproto/transport"

	_ "github.com/lib/pq"
	"fmt"
	"strings"

)

type ParkingRepo struct {
    db *sql.DB
}


func NewParkingRepo(db *sql.DB) *ParkingRepo {
	return &ParkingRepo{
        db: db,
    }
}

func (r *ParkingRepo) CreateParkingSpcae(req *pb.ParkingSpaceCreate) (*pb.ParkingSpace, error) {
	res := &pb.ParkingSpace{}

    query := `INSERT INTO parking_spaces (lot_id, space_number, is_occupied)
    VALUES ($1, $2, $3) RETURNING id`

    var id string
    err := r.db.QueryRow(query, req.LotId, req.SpacesNumber, req.IsOccupied).Scan(&id)
    if err!= nil {
        return nil, err
    }

    res, err = r.GetParkingSpace(&pb.ById{Id: id})
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *ParkingRepo) GetParkingSpace(id *pb.ById) (*pb.ParkingSpace, error) {
	res := &pb.ParkingSpace{}

    query := `SELECT id, lot_id, space_number, is_occupied FROM parking_spaces WHERE id = $1`

    err := r.db.QueryRow(query, id.Id).Scan(
		&res.SpaceId, 
		&res.LotId, 
		&res.SpacesNumber, 
		&res.IsOccupied,
	)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func (r *ParkingRepo) GetAllParkingSpaces(req *pb.ParkingSpaceFilter) (*pb.ParkingSpaceList, error) {
	res := &pb.ParkingSpaceList{}

    query := `SELECT id, lot_id, space_number, is_occupied FROM parking_spaces where deleted_at is null`

    var args []interface{}
	var conditions []string

	if req.LotId != "" {
		args = append(args, req.LotId)
		conditions = append(conditions, fmt.Sprintf("lot_id = $%d", len(args)))
	}

	if req.LotId != "" {
		args = append(args, req.SpacesNumber)
		conditions = append(conditions, fmt.Sprintf("space_number = $%d", len(args)))
	}

	if req.LotId != "" {
		args = append(args, req.IsOccupied)
		conditions = append(conditions, fmt.Sprintf("is_occupied = $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var defaultLimit int32
	err := r.db.QueryRow("SELECT COUNT(1) FROM parking_spaces WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		return nil, err
	}
	if req.Filter.Limit == 0 {
		req.Filter.Limit = defaultLimit
	}

	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


    for rows.Next() {
        parkingSpace := &pb.ParkingSpace{}

        err := rows.Scan(
			&parkingSpace.SpaceId, 
			&parkingSpace.LotId, 
			&parkingSpace.SpacesNumber, 
			&parkingSpace.IsOccupied,
		)
        if err != nil {
            return nil, err
        }

        res.ParkingSpaces = append(res.ParkingSpaces, parkingSpace)
    }

    return res, nil
}

func (r *ParkingRepo) UpdateParkingSpace(req *pb.ParkingSpaceUpdate) (*pb.ParkingSpace, error) {
	res := &pb.ParkingSpace{}

    query := `UPDATE parking_spaces SET lot_id = $1, space_number = $2, is_occupied = $3, updated_at=now() WHERE id = $4 RETURNING id`

    var id string

    err := r.db.QueryRow(query, req.LotId, req.SpacesNumber, req.IsOccupied, req.SpaceId).Scan(&id)
    if err!= nil {
        return nil, err
    }

    res, err = r.GetParkingSpace(&pb.ById{Id: id})
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *ParkingRepo) DeleteParkingSpace(id *pb.ById) (*pb.Void, error) {
	res := &pb.Void{}

    query := `UPDATE parking_spaces SET deleted_at=EXTRACT(EPOCH FROM now() WHERE id = $1`

    _, err := r.db.Exec(query, id.Id)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *ParkingRepo) CreateParkingLot(req *pb.ParkingLotCreate) (*pb.ParkingLot, error) {
	res := &pb.ParkingLot{}

    query := `INSERT INTO parking_lots (lot_name, total_spaces, free_space, occupied_space, address)
    VALUES ($1) RETURNING id`

	tr, err := r.db.Begin()
	if err!= nil {
        return nil, err
    }

	var id string
	err = tr.QueryRow(query, req.LotName, req.TotalSpaces, req.FreeSpaces, req.OccupiedSpaces, req.Address).Scan(&id)
	if err!= nil {
		tr.Rollback()
        return nil, err
    }

	var i int32
	for i = 1; i <= req.TotalSpaces; i++ {
		query := `INSERT INTO parking_spaces (lot_id, space_number, is_occupied)
        VALUES ($1, $2, $3)`

        _, err = tr.Exec(query, id, i, false)
        if err!= nil {
            tr.Rollback()
            return nil, err
        }
	}

	res, err = r.GetParkingStatus(&pb.ById{Id: id})
	if err!= nil {
        tr.Rollback()
        return nil, err
    }

	tr.Commit()
	return res, nil
}

func (r *ParkingRepo) UpdateParkingLot(req *pb.ParkingLotUpdate) (*pb.ParkingLot, error) {
	res := &pb.ParkingLot{}

    query := `UPDATE parking_lots SET lot_name = $1, total_spaces = $2, free_space = $3, occupied_space = $4, address = $5, updated_at=now() WHERE id = $6 RETURNING id`

    var id string

    err := r.db.QueryRow(query, req.LotName, req.TotalSpaces, req.FreeSpaces, req.OccupiedSpaces, req.Address, req.LotId).Scan(&id)
    if err!= nil {
        return nil, err
    }

    res, err = r.GetParkingStatus(&pb.ById{Id: id})
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *ParkingRepo) GetParkingStatus(id *pb.ById) (*pb.ParkingLot, error) {
	res := &pb.ParkingLot{}

	tr, err := r.db.Begin()

	var total, occuoied, free int32
	err = tr.QueryRow("SELECT COUNT(1) FROM parking_spaces WHERE deleted_at=0 AND lot_id = $1", id.Id).Scan(&total)
	if err!= nil {
		tr.Rollback()
		return nil, err
    }

	err = tr.QueryRow("SELECT COUNT(1) FROM parking_spaces WHERE deleted_at=0 AND lot_id = $1 AND is_occupied = true", id.Id).Scan(&occuoied)
	if err!= nil {
		tr.Rollback()
		return nil, err
    }

	free = total - occuoied

	query := `UPDATE total_spaces=$1, free_space=$2, occupied_space=$3 FROM parking_lots WHERE id=$4`
	_, err = tr.Exec(query, total, free, occuoied, id.Id)

	query = `SELECT id, lot_name, total_spaces, free_space, occupied_space, address FROM parking_lots WHERE id=$1`

	err = tr.QueryRow(query, id.Id).Scan(
		&res.LotId, 
		&res.LotName, 
		&res.TotalSpaces,
		 &res.FreeSpaces, 
		 &res.OccupiedSpaces, 
		 &res.Address,
		)
	if err!= nil {
		tr.Rollback()
        return nil, err
    }

	err = tr.Commit()
	if err!= nil {
		tr.Rollback()
        return nil, err
    }

	return res, nil
}

func (r *ParkingRepo) DeleteParkingLot(id *pb.ById) (*pb.Void, error) {
	res := &pb.Void{}

	tr, err := r.db.Begin()
	if err!= nil {
        return nil, err
    }
    query := `UPDATE parking_lots SET deleted_at=EXTRACT(EPOCH FROM now() WHERE id = $1`

    _, err = tr.Exec(query, id.Id)
    if err!= nil {
		tr.Rollback()
        return nil, err
    }

	query = `UPDATE parking_spaces SET deleted_at=EXTRACT(EPOCH FROM now() WHERE lot_id=$1`
	_, err = tr.Exec(query, id.Id)
	if err!= nil {
        tr.Rollback()
        return nil, err
    }
	err = tr.Commit()
	if err!= nil {
		tr.Rollback()
	}

    return res, nil
}

func (r *ParkingRepo) GetAllParkingLots(req *pb.ParkingLotFilter) (*pb.PakingLotList, error) {
	res := &pb.PakingLotList{}

    query := `SELECT id, lot_name, total_spaces, free_space, occupied_space, address FROM parking_lots WHERE deleted_at=0`

    var conditions []string
    var args []interface{}

    if req.MinTotalSpcaes!= 0 {
        args = append(args, req.MinTotalSpcaes)
        conditions = append(conditions, fmt.Sprintf("total_spaces >= $%d", len(args)))
    }
	if req.MinFreeSpcaes!= 0 {
        args = append(args, req.MinFreeSpcaes)
        conditions = append(conditions, fmt.Sprintf("free_space >= $%d", len(args)))
    }
	if req.MaxTotalSpcaes!= 0 {
        args = append(args, req.MaxTotalSpcaes)
        conditions = append(conditions, fmt.Sprintf("total_spaces <= $%d", len(args)))
    }
	if req.MaxFreeSpcaes!= 0 {
        args = append(args, req.MaxFreeSpcaes)
        conditions = append(conditions, fmt.Sprintf("free_space <= $%d", len(args)))
    }

    if len(conditions) > 0 {
        query += " AND " + strings.Join(conditions, " AND ")
    }

    var defaultLimit int32
    err := r.db.QueryRow("SELECT COUNT(1) FROM parking_lots WHERE deleted_at=0").Scan(&defaultLimit)
    if err != nil {
        return nil, err
    }
	if req.Filter.Limit == 0 {
        req.Filter.Limit = defaultLimit
    }

	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	defer rows.Close()

	if err!= nil {
        return nil, err
    }

	for rows.Next() {
		lot := &pb.ParkingLot{}
        err = rows.Scan(&lot.LotId, &lot.LotName, &lot.TotalSpaces, &lot.FreeSpaces, &lot.OccupiedSpaces, &lot.Address)
        if err!= nil {
            return nil, err
        }
        res.ParkingLots = append(res.ParkingLots, lot)
	}

	res.Count = int32(len(res.ParkingLots))

	return res, err
}

func (r *ParkingRepo) ReserveParking(req *pb.ParkingReservationCreate) (*pb.ParkingReservation, error) {
	res := &pb.ParkingReservation{}

    tr, err := r.db.Begin()
    if err!= nil {
        return nil, err
    }

    var id string
    err = tr.QueryRow("INSERT INTO parking_reservations (parking_space_id, citizen_id, time_from, time_to, status) VALUES ($1, $2, $3, $4, $5) RETURNING id", req.ParkingSpaceId, req.CitizenId, req.TimeFrom, req.TimeTo, req.Status).Scan(&id)
    if err!= nil {
        tr.Rollback()
        return nil, err
    }

    query := `UPDATE parking_spaces SET is_occupied = true WHERE id=$1`
    _, err = tr.Exec(query, req.ParkingSpaceId)
	if err!= nil {
		tr.Rollback()
        return nil, err
    }

	res, err = r.GetReserveParking(&pb.ById{Id: id})
	if err!= nil {
        tr.Rollback()
        return nil, err
    }

	err = tr.Commit()
	if err!= nil {
        tr.Rollback()
        return nil, err
    }

	return res, nil
}

func (r *ParkingRepo) GetReserveParking(id *pb.ById) (*pb.ParkingReservation, error) {
	res := &pb.ParkingReservation{}

	query := `SELECT id, parking_space_id, citizen_id, time_from, time_to, status FROM parking_reservations WHERE id = $1`
	err := r.db.QueryRow(query, id.Id).Scan(
		&res.ReservationId,
		&res.ParkingSpaceId,
		&res.CitizenId,
		&res.TimeFrom,
		&res.TimeTo,
        &res.Status,
	)

	if err!= nil {
        return nil, err
    }
	return res, nil
}

func (r *ParkingRepo) GetAllReserveParking(req *pb.ParkingReservationFilter) (*pb.ParkingReservationList, error) {
	res := &pb.ParkingReservationList{}

    query := `SELECT id, parking_space_id, citizen_id, time_from, time_to, status FROM parking_reservations WHERE deleted_at=0`

    var conditions []string
    var args []interface{}

    if req.ParkingSpaceId!= "" {
        args = append(args, req.ParkingSpaceId)
        conditions = append(conditions, fmt.Sprintf("parking_space_id = $%d", len(args)))
    }
    if req.CitizenId!= "" {
        args = append(args, req.CitizenId)
        conditions = append(conditions, fmt.Sprintf("citizen_id = $%d", len(args)))
    }
	if req.Time != "" {
		args = append(args, req.Time, req.Time)
		conditions = append(conditions, fmt.Sprintf("$%d BETWEEN time_from AND time_to", len(args)))
	}	
    if req.Status!= "" {
        args = append(args, fmt.Sprintf("%" + req.Status + "%"))
        conditions = append(conditions, fmt.Sprintf("status ILIKE $%d", len(args)))
    }

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var defaultLimit int32
	err := r.db.QueryRow("SELECT COUNT(1) FROM parking_reservations WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		return nil, err
	}
	if req.Filter.Limit == 0 {
		req.Filter.Limit = defaultLimit
	}

	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	defer rows.Close()
	if err!= nil {
        return nil, err
    }

	for rows.Next() {
		reservation := &pb.ParkingReservation{}
        err = rows.Scan(&reservation.ReservationId, &reservation.ParkingSpaceId, &reservation.CitizenId, &reservation.TimeFrom, &reservation.TimeTo, &reservation.Status)
        if err!= nil {
            return nil, err
        }
        res.ParkingReservations = append(res.ParkingReservations, reservation)
    }

	res.Count = int32(len(res.ParkingReservations))

	return res, err
}

func (r *ParkingRepo) ReserveParkingUpdate(req *pb.ParkingReservationUpdate) (*pb.ParkingReservation, error) {
	res := &pb.ParkingReservation{}

    tr, err := r.db.Begin()
    if err!= nil {
        return nil, err
    }

	var space_id string
	err = tr.QueryRow("SELECT parking_space_id FROM parking_reservations WHERE id=$1", req.ReservationId).Scan(&space_id)
	if err!= nil {
		tr.Rollback()
		return nil, err
	}
	query := `UPDATE parking_spaces SET is_occupied = false WHERE id=$1`
    _, err = tr.Exec(query, space_id)

    var id string
    err = tr.QueryRow("UPDATE parking_reservations SET parking_space_id=$1, citizen_id=$2, time_from=$3, time_to=$4, status=$5 WHERE id=$6 RETURNING id", req.ParkingSpaceId, req.CitizenId, req.TimeFrom, req.TimeTo, req.Status, req.ReservationId).Scan(&id)
    if err!= nil {
        tr.Rollback()
        return nil, err
    }

    query = `UPDATE parking_spaces SET is_occupied = true WHERE id=$1`
    _, err = tr.Exec(query, req.ParkingSpaceId)

	res, err = r.GetReserveParking(&pb.ById{Id: id})
	if err!= nil {
        tr.Rollback()
        return nil, err
    }
	err = tr.Commit()

	return res, nil
}

func (r *ParkingRepo) ReserveParkingDelete(id *pb.ById) (*pb.Void, error) {
	res := &pb.Void{}

    tr, err := r.db.Begin()
    if err!= nil {
        return nil, err
    }

    var space_id string
    err = tr.QueryRow("SELECT parking_space_id FROM parking_reservations WHERE id=$1", id.Id).Scan(&space_id)
    if err!= nil {
        tr.Rollback()
        return nil, err
    }
    query := `UPDATE parking_spaces SET is_occupied = false WHERE id=$1`
    _, err = tr.Exec(query, space_id)

    _, err = tr.Exec("UPDATE parking_reservations SET deleted_at=NOW() WHERE id=$1", id.Id)
    if err!= nil {
        tr.Rollback()
        return nil, err
    }

    err = tr.Commit()

	return res, nil
}

