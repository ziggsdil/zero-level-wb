package db

const (
	// сюда будет приходить uid из фронта
	selectDataByOrderId = `
		SELECT * FROM orders
			WHERE order_uid=$1
	`

	saveMessage = `
		INSERT INTO orders(order_uid, data)
			VALUES ($1, $2)
	`
)
