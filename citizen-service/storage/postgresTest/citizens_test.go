package postgrestest

import (
	"citizen/genproto/citizen"
	"citizen/storage/postgres"
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

// Helper function to initialize PostgreSQL connection
func setupDBConn(t *testing.T) *postgres.Citizen {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"postgres",
		"root",
		"localhost",
		5432,
		"citizen_db",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &postgres.Citizen{Db: db}
}

// Helper function to create a new citizen object for testing
func newCitizenTest() *citizen.CitizenCreate {
	return &citizen.CitizenCreate{
		CitizenId:   uuid.NewString(),
		UserId:      uuid.NewString(),
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1990-01-01",
		Address:     "123 Main St",
		PhoneNumber: "123-456-7890",
		Email:       "john.doyasdsa@example.com",
	}
}

// Test function to register a citizen
func TestRegisterCitizen(t *testing.T) {
	citiDB := setupDBConn(t)
	citizenTest := newCitizenTest()

	citiRes, err := citiDB.RegisterCitizen(context.Background(), citizenTest)
	if err != nil {
		t.Fatalf("Error registering citizen: %v", err)
	}

	assert.Equal(t, citizenTest.FirstName, citiRes.FirstName)
	assert.Equal(t, citizenTest.LastName, citiRes.LastName)
	assert.Equal(t, citizenTest.DateOfBirth, citiRes.DateOfBirth)
	assert.Equal(t, citizenTest.Address, citiRes.Address)
	assert.Equal(t, citizenTest.PhoneNumber, citiRes.PhoneNumber)
	assert.Equal(t, citizenTest.Email, citiRes.Email)
	assert.NotEmpty(t, citiRes.CitizenId)
}


// Test function to update a citizen profile
func TestUpdateCitizenProfile(t *testing.T) {
	citiDB := setupDBConn(t)
	citizenTest := newCitizenTest()

	// Register a new citizen to update
	citiRes, err := citiDB.RegisterCitizen(context.Background(), citizenTest)
	if err != nil {
		t.Fatalf("Error registering citizen: %v", err)
	}

	// Update fields
	citizenTest.FirstName = "Jane"
	citizenTest.LastName = "Smith"
	citizenTest.DateOfBirth = "1992-01-02"
	citizenTest.Address = "456 Another St"
	citizenTest.PhoneNumber = "987-654-3210"
	citizenTest.Email = "jane.smithoab@example.com"
	citizenTest.CitizenId = citiRes.CitizenId

	updatedRes, err := citiDB.UpdateCitizenProfile(context.Background(), citizenTest)
	if err != nil {
		t.Fatalf("Error updating citizen profile: %v", err)
	}

	assert.Equal(t, citizenTest.FirstName, updatedRes.FirstName)
	assert.Equal(t, citizenTest.LastName, updatedRes.LastName)
	assert.Equal(t, citizenTest.DateOfBirth, updatedRes.DateOfBirth)
	assert.Equal(t, citizenTest.Address, updatedRes.Address)
	assert.Equal(t, citizenTest.PhoneNumber, updatedRes.PhoneNumber)
	assert.Equal(t, citizenTest.Email, updatedRes.Email)
}


// Test function to delete a citizen profile
func TestDeleteCitizenProfile(t *testing.T) {
	citiDB := setupDBConn(t)

	// Create a citizen for testing
	citizenTest := newCitizenTest()
	citiRes, err := citiDB.RegisterCitizen(context.Background(), citizenTest)
	if err != nil {
		t.Fatalf("Error registering citizen: %v", err)
	}

	// Delete the citizen profile
	req := &citizen.ById{Id: citiRes.CitizenId}
	_, err = citiDB.DeleteCitizenProfile(context.Background(), req)
	if err != nil {
		t.Fatalf("Error deleting citizen profile: %v", err)
	}

	// Try to retrieve the deleted citizen profile
	_, err = citiDB.GetCitizenProfile(context.Background(), req)
	assert.Error(t, err, "expected error for deleted profile")
}

