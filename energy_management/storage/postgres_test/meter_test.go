package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
	"gitlab.com/acumen931026/energy_management/storage/postgres"
)

func TestCreateMeterReading(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewMeterReadingsStorage(db)

	req := &pb.MeterReading{
		MeterId:      "some-meter-uuid",
		ReadingValue: 123.45,
		Time:         "2024-06-28T00:00:00Z",
	}

	mock.ExpectExec(`INSERT INTO meter_readings`).
		WithArgs(sqlmock.AnyArg(), req.MeterId, req.ReadingValue, req.Time).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := store.CreateMeterReading(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMeterReading(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewMeterReadingsStorage(db)

	req := &pb.ById{Id: "some-reading-uuid"}

	
	mock.ExpectQuery(`SELECT m.reading_id, m.meter_id, m.reading_value, e.meter_id, e.building_id, e.meter_type, e.created_at FROM meter_readings AS m JOIN energy_meters AS e ON m.meter_id = e.meter_id WHERE m.reading_id = \$1`).
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"reading_id", "m.meter_id", "reading_value", "time", "e.meter_id", "building_id", "created_at"}).
			AddRow(req.Id, "some-meter-uuid", 123.45, "2024-06-28T00:00:00Z", "some-meter-uuid", "some-building-uuid", "2024-06-28T00:00:00Z"))

	resp, err := store.GetMeterReading(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, req.Id, resp.ReadingId)
	assert.Equal(t, "some-meter-uuid", resp.MeterId)
	assert.Equal(t, float32(123.45), resp.ReadingValue)
	assert.Equal(t, "2024-06-28T00:00:00Z", resp.Time)
	assert.Equal(t, "some-meter-uuid", resp.Energy.MeterId)
	assert.Equal(t, "some-building-uuid", resp.Energy.BuildingId)
	assert.Equal(t, "2024-06-28T00:00:00Z", resp.Energy.CreatedAt)

	assert.NoError(t, mock.ExpectationsWereMet())
}



func TestDeleteMeterReading(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewMeterReadingsStorage(db)

	req := &pb.ById{Id: "some-meter-uuid"}

	mock.ExpectExec(`UPDATE meter_readings SET deleted_at = EXTRACT\(EPOCH FROM NOW\(\)\) WHERE meter_id = \$1`).
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := store.DeleteMeterReading(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestGetHourlyEnergyForBuilding(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewMeterReadingsStorage(db)

	req := &pb.ByHour{
		BuildingId: "some-building-uuid",
		Time:       "2024-06-28T00:00:00Z",
	}

	mock.ExpectQuery(`SELECT meter_id, COALESCE\(SUM\(reading_value\), 0\) / 24 AS average_hourly_energy_value FROM meter_readings WHERE meter_id = \$1 GROUP BY meter_id`).
		WithArgs(req.BuildingId).
		WillReturnRows(sqlmock.NewRows([]string{"meter_id", "average_hourly_energy_value"}).
			AddRow(req.BuildingId, 44.44))

	resp, err := store.GetHourlyEnergyForBuilding(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.BuildingId, resp.BuildingId)
	assert.Equal(t, float32(44.44), resp.EnergyValue) // Tekshirishda float32 dan foydalaning
	assert.Equal(t, "2024-06-28T00:00:00Z", resp.Time)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDailyEnergyForBuilding(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := postgres.NewMeterReadingsStorage(db)

	req := &pb.ById{Id: "some-building-uuid"}

	mock.ExpectQuery(`SELECT meter_id, COALESCE\(SUM\(reading_value\), 0\) as energy_value FROM meter_readings WHERE meter_id = \$1 GROUP BY meter_id`).
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"meter_id", "energy_value"}).
			AddRow(req.Id, 1067.89))

	resp, err := store.GetDailyEnergyForBuilding(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.BuildingId)
	assert.Equal(t, float32(1067.89), resp.EnergyValue) // Tekshirishda float32 dan foydalaning

	assert.NoError(t, mock.ExpectationsWereMet())
}
