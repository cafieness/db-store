package worker

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"time"
)

func generateActivity(db *sql.DB) {
	userID := rand.IntN(10) + 1
	amount := rand.IntN(5000) + 100

	statuses := []string{
		"delivered",
		"processing",
		"pending",
		"received",
	}

	status := statuses[rand.IntN(len(statuses))]

	var orderID int
	err := db.QueryRow(
		"insert into orders (amount, order_status, user_id) values ($1, $2, $3) RETURNING id",
		amount, status, userID,
	).Scan(&orderID)

	if err != nil {
		return
	}

	for i := 0; i < rand.IntN(3)+1; i++ {
		productID := rand.IntN(10) + 1

		db.Exec(
			"insert into order_items (product_id, order_id, unit_price, quantity) values ($1, $2, $3, $4)",
			productID,
			orderID,
			rand.IntN(100)+1,
			rand.IntN(3)+1,
		)
	}

	for i := 0; i < rand.IntN(5)+1; i++ {
		db.Exec(
			"insert into product_view (product_id, user_id) values ($1, $2)",
			rand.IntN(10)+1,
			userID,
		)
	}

	fmt.Println("generated fake activity")
}

func StartSimulator(db *sql.DB) {
	go func() {

		ticker := time.NewTicker(20 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			generateActivity(db)
		}
	}()
}
