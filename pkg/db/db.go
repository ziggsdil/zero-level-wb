package db

import (
	"database/sql"
	"fmt"
)

type Database struct {
	client *sql.DB
}

func NewDatabase(config Config) (*Database, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)

	client, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}

	return &Database{client: client}, nil
}
