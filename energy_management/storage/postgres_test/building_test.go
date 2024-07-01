package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/acumen931026/energy_management/storage/postgres"
	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
)

func TestCreateBuilding(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewBuildingsStorage(db)

	req := &pb.BuildingCreate{
		Address:      "123 Main St",
		BuildingType: "Residential",
		TotalArea:    "2000",
	}

	mock.ExpectQuery(`INSERT INTO buildings`).
		WithArgs(sqlmock.AnyArg(), req.Address, req.BuildingType, req.TotalArea).
		WillReturnRows(sqlmock.NewRows([]string{"building_id", "address", "building_type", "total_area"}).
			AddRow("some-uuid", req.Address, req.BuildingType, req.TotalArea))

	resp, err := store.CreateBuilding(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "some-uuid", resp.BuildingId)
	assert.Equal(t, req.Address, resp.Address)
	assert.Equal(t, req.BuildingType, resp.BuildingType)
	assert.Equal(t, req.TotalArea, resp.TotalArea)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateBuilding(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewBuildingsStorage(db)

	req := &pb.BuildingUpdate{
		Id:           "some-uuid",
		Address:      "456 Elm St",
		BuildingType: "Commercial",
		TotalArea:    "3000",
	}

	mock.ExpectQuery(`UPDATE buildings`).
		WithArgs(req.Address, req.BuildingType, req.TotalArea, req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"building_id", "address", "building_type", "total_area"}).
			AddRow(req.Id, req.Address, req.BuildingType, req.TotalArea))

	resp, err := store.UpdateBuilding(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.BuildingId)
	assert.Equal(t, req.Address, resp.Address)
	assert.Equal(t, req.BuildingType, resp.BuildingType)
	assert.Equal(t, req.TotalArea, resp.TotalArea)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetBuilding(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewBuildingsStorage(db)

	req := &pb.ById{Id: "some-uuid"}

	mock.ExpectQuery(`SELECT building_id, address, building_type, total_area FROM buildings`).
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"building_id", "address", "building_type", "total_area"}).
			AddRow(req.Id, "123 Main St", "Residential", 2000))

	resp, err := store.GetBuilding(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.BuildingId)
	assert.Equal(t, "123 Main St", resp.Address)
	assert.Equal(t, "Residential", resp.BuildingType)
	assert.Equal(t, "2000", resp.TotalArea)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteBuilding(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewBuildingsStorage(db)

	req := &pb.ById{Id: "some-uuid"}

	mock.ExpectExec(`UPDATE buildings SET deleted_at = EXTRACT\(EPOCH FROM NOW\(\)\) WHERE building_id = \$1`).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := store.DeleteBuilding(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.NoError(t, mock.ExpectationsWereMet())
}
