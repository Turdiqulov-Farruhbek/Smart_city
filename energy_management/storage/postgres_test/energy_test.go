package postgres_test

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
	"gitlab.com/acumen931026/energy_management/storage/postgres"
)

func TestCreateEnergyMeter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewEnergyMetersStorage(db)

	req := &pb.EnergyMeterCreate{
		BuildingId: "some-building-uuid",
		MeterType:  "Electric",
	}

	mock.ExpectQuery(`INSERT INTO energy_meters`).
		WithArgs(sqlmock.AnyArg(), req.BuildingId, req.MeterType).
		WillReturnRows(sqlmock.NewRows([]string{"meter_id", "building_id", "meter_type"}).
			AddRow("some-meter-uuid", req.BuildingId, req.MeterType))

	resp, err := store.CreateEnergyMeter(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "some-meter-uuid", resp.MeterId)
	assert.Equal(t, req.BuildingId, resp.BuildingId)
	assert.Equal(t, req.MeterType, resp.MeterType)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateEnergyMeter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewEnergyMetersStorage(db)

	req := &pb.EnergyMeterUpdate{
		MeterId:    "some-meter-uuid",
		BuildingId: "some-building-uuid",
		MeterType:  "Water",
	}

	mock.ExpectQuery(`UPDATE energy_meters`).
		WithArgs(req.BuildingId, req.MeterType, req.MeterId).
		WillReturnRows(sqlmock.NewRows([]string{"meter_id", "building_id", "meter_type"}).
			AddRow(req.MeterId, req.BuildingId, req.MeterType))

	resp, err := store.UpdateEnergyMeter(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.MeterId, resp.MeterId)
	assert.Equal(t, req.BuildingId, resp.BuildingId)
	assert.Equal(t, req.MeterType, resp.MeterType)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetEnergyMeter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewEnergyMetersStorage(db)

	req := &pb.ById{Id: "some-meter-uuid"}

	mock.ExpectQuery(`SELECT e\.meter_id, e\.building_id, e\.meter_type, e\.created_at, b\.building_id, b\.address, b\.building_type, b\.total_area, b\.created_at FROM energy_meters AS e JOIN buildings AS b ON e\.building_id = b\.building_id WHERE e\.meter_id = \$1`).
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"meter_id", "building_id", "meter_type", "e_created_at", "building_id", "address", "building_type", "total_area", "b_created_at"}).
			AddRow(req.Id, "some-building-uuid", "Electric", "2024-06-28T00:00:00Z", "some-building-uuid", "123 Main St", "Residential", 2000, "2024-01-01T00:00:00Z"))

	resp, err := store.GetEnergyMeter(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.Meter.MeterId)
	assert.Equal(t, "some-building-uuid", resp.Meter.BuildingId)
	assert.Equal(t, "Electric", resp.Meter.MeterType)
	assert.Equal(t, "2024-06-28T00:00:00Z", resp.Meter.CreatedAt)
	assert.Equal(t, "some-building-uuid", resp.Building.BuildingId)
	assert.Equal(t, "123 Main St", resp.Building.Address)
	assert.Equal(t, "Residential", resp.Building.BuildingType)
	assert.Equal(t, "2000", resp.Building.TotalArea)
	assert.Equal(t, "2024-01-01T00:00:00Z", resp.Building.CreatedAt)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteEnergyMeter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewEnergyMetersStorage(db)

	req := &pb.ById{Id: "some-meter-uuid"}

	mock.ExpectExec(`UPDATE energy_meters SET deleted_at = EXTRACT\(EPOCH FROM NOW\(\)\) WHERE meter_id = \$1`).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := store.DeleteEnergyMeter(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.NoError(t, mock.ExpectationsWereMet())
}
