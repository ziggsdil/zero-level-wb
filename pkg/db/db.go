package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

func (db *Database) GetAllData(ctx context.Context) (map[string][]byte, error) {
	rows, err := db.client.QueryContext(ctx, selectAllData)
	if err != nil {
		fmt.Printf("Failed to do query: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	data := make(map[string][]byte)

	for rows.Next() {
		var orderId string
		var jsonData []byte

		if err = rows.Scan(&orderId, &jsonData); err != nil {
			return data, err
		}
		data[orderId] = jsonData
	}

	if err = rows.Err(); err != nil {
		return data, err
	}

	return data, nil
}

func (db *Database) GetOrder(ctx context.Context, orderId string) ([]byte, error) {
	row := db.client.QueryRowContext(ctx, selectDataByOrderId, orderId)

	jsonData := make([]byte, 0)
	err := row.Scan(&jsonData)
	if err != nil {
		fmt.Printf("Failed to scan order: %s\n", err.Error())
		return nil, err
	}
	return jsonData, err
}

func (db *Database) InsertData(ctx context.Context, orderId string, data []byte) error {
	// todo: check if orderId already exist
	// we can use map for check
	_, err := db.client.ExecContext(ctx, saveMessage, orderId, data)
	if err == nil {
		log.Printf("Data with orderId: %s was success saved\n", orderId)
	}
	return err
}
