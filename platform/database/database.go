package database

import (
	"database/sql"
	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

func GetOrderDataFromDB(orderID string) (string, error) {
	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", "user=username dbname=yourdb sslmode=disable")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Выполняем SQL-запрос для получения данных заказа
	var orderData string
	err = db.QueryRow("SELECT order_data FROM orders WHERE order_id = $1", orderID).Scan(&orderData)
	if err != nil {
		return "", err
	}

	return orderData, nil
}
