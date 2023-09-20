// db_test.go
package db

import (
	_ "database/sql"
	"github.com/shirrashko/BuildingAServer-step2/config"
	"testing"
)

func Test_NewDBClient(t *testing.T) {
	// Create a mock configuration for testing
	mockConfig := config.DBConfig{
		// Fill in the mock configuration values
		Host:         "localhost",
		Port:         6432,
		User:         "srashkovits",
		Password:     "password",
		DatabaseName: "postgres",
	}

	// Call the NewDBClient function with the mock configuration
	dbClient, err := NewDBClient(mockConfig)
	if err != nil {
		t.Errorf("Error creating DB client: %v", err)
	}

	// Check that the returned dbClient is not nil
	if dbClient == nil {
		t.Errorf("Expected non-nil dbClient, got nil")
	}

}
