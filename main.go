package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cafieness/db-store.git/queries"
	"github.com/cafieness/db-store.git/worker"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://dev_user:dev_pass_301@localhost:5432/app_db?sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Ping error:", err)
	}
	fmt.Println("Connected to database!")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server Started!")
	})

	mux.HandleFunc("/analytics/top-products", func(w http.ResponseWriter, r *http.Request) {
		topProducts, err := queries.GetTopProducts(db)
		if err != nil {
			log.Println(err)
			http.Error(w, "Top Products Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(topProducts); err != nil {
			log.Println("JSON encode error:", err)
		}
	})
	mux.HandleFunc("/analytics/revenue", func(w http.ResponseWriter, r *http.Request) {
		revenue, err := queries.GetRevenue(db)
		if err != nil {
			log.Println(err)
			http.Error(w, "Revenue Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(revenue); err != nil {
			log.Println("JSON encode error:", err)
		}
	})
	mux.HandleFunc("/analytics/productview", func(w http.ResponseWriter, r *http.Request) {
		productViews, err := queries.GetProductViews(db)
		if err != nil {
			log.Println(err)
			http.Error(w, "Product View Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(productViews); err != nil {
			log.Println("JSON encode error:", err)
		}
	})
	worker.StartWorker(db)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
