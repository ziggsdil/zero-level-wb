package db

import (
	"context"
	"database/sql"
	"encoding/json"
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

func (db *Database) GetOrder(ctx context.Context, orderId string) (map[string]interface{}, error) {
	row := db.client.QueryRowContext(ctx, selectDataByOrderId, orderId)

	jsonData := make([]byte, 0)
	err := row.Scan(&jsonData)
	if err != nil {
		return nil, err
	}

	dataMap := make(map[string]interface{})

	err = json.Unmarshal(jsonData, &dataMap)
	if err != nil {
		return nil, err
	}

	return dataMap, err
}

func (db *Database) InsertData(ctx context.Context, orderId string, data []byte) error {
	_, err := db.client.ExecContext(ctx, saveMessage, orderId, data)
	return err
}
