package db

import (
	"database/sql"
	"fmt"
	"github.com/shirrashko/BuildingAServer-step2/config"
	"os"
)

func NewDBClient(connectionInfo config.DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connectionInfo.Host,
		connectionInfo.Port, connectionInfo.User, connectionInfo.Password, connectionInfo.DatabaseName)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}

	// Execute the table creation script
	script, err := os.ReadFile("/Users/srashkovits/repos/onboarding/schema/create_table.sql")
	if err != nil {
		fmt.Printf("Error reading create_tables.sql: %v\n", err)
		return nil, err
	}

	_, err = db.Exec(string(script)) // execute the script create a table named userProfiles with columns names
	// according to the api.model
	if err != nil {
		fmt.Printf("Error executing create_tables.sql: %v\n", err)
		return nil, err
	}

	return db, nil
}
