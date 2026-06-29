package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://dev_user:dev_pass_301@localhost:5432/app_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 1; i <= 10; i++ {
		_, err := db.Exec(
			"insert into users (fullname, email, phone, password_hash, address) values ($1, $2, $3, $4, $5)",
			fmt.Sprintf("User %d", i),
			fmt.Sprintf("user%d@mail.com", i),
			fmt.Sprintf("+755512%04d", i),
			"hashed_password",
			fmt.Sprintf("adress %d", i),
		)
		if err != nil {
			log.Fatal("users insert error:", err)
		}
	}
	fmt.Println("users seed created")

	db.Exec("insert into categories (name) VALUES ($1)", "Electronics")
	db.Exec("insert into categories (name) VALUES ($1)", "People")
	db.Exec("insert into categories (name) VALUES ($1)", "Bananas")

	db.Exec("insert into categories (name, parent_id) VALUES ($1, $2)", "phones", 1)
	db.Exec("insert into categories (name, parent_id) VALUES ($1, $2)", "Cavendish", 3)
	fmt.Println("categories seed created")

	for i := 1; i <= 10; i++ {
		n := rand.IntN(3) + 1

		_, err := db.Exec(
			"insert into products (name, description, price, category_id, image_url) values ($1, $2, $3, $4, $5)",
			fmt.Sprintf("Product %d", i),
			fmt.Sprintf("Description %d", i),
			float64(i*1000),
			n,
			fmt.Sprintf("https://example.com/images/product%d.jpg", i),
		)
		if err != nil {
			log.Fatal("products insert error:", err)
		}
	}

	fmt.Println("products seed created")

	for i := 1; i <= 20; i++ {
		n := rand.IntN(5) + 1
		userID := rand.IntN(10) + 1
		s := map[int]string{
			1: "delivered",
			2: "received",
			3: "processing",
			4: "canceled",
			5: "pending",
		}
		_, err := db.Exec(
			"insert into orders (amount, order_status, user_id) values ($1, $2, $3)",
			fmt.Sprintf("%d0000", i),
			s[n],
			userID,
		)
		if err != nil {
			log.Fatal("orders insert error:", err)
		}
	}

	fmt.Println("orders seed created")

	for i := 1; i <= 10; i++ {
		orderID := rand.IntN(20) + 1
		_, err := db.Exec(
			"insert into order_items (product_id, order_id, unit_price, quantity) values ($1, $2, $3, $4)",
			i,
			orderID,
			float64(i*1000),
			i,
		)
		if err != nil {
			log.Fatal("order_items insert error:", err)
		}
	}

	fmt.Println("order_items seed created")

	for i := 1; i <= 1000; i++ {
		productID := rand.IntN(10) + 1
		userID := rand.IntN(10) + 1
		_, err := db.Exec(
			"insert into product_view (product_id, user_id) values ($1, $2)",
			productID,
			userID,
		)
		if err != nil {
			log.Fatal("product_view insert error:", err)
		}
	}

	fmt.Println("product_view seed created")

	for i := 1; i <= 10; i++ {
		changedAt := time.Now().AddDate(0, 0, -rand.IntN(30))
		_, err := db.Exec(
			"insert into price_history (product_id, price, changed_at) values ($1, $2, $3)",
			i,
			float64(i*1000),
			changedAt,
		)
		if err != nil {
			log.Fatal("price_history insert error:", err)
		}
	}

	fmt.Println("price_history seed created")

}
