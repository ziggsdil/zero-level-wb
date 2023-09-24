package db

const (
	// сюда будет приходить uid из фронта
	selectDataByOrderId = `
		SELECT * FROM orders
			WHERE order_uid=$1
	`
)
