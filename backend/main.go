package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cafieness/db-store.git/backend/queries"
	"github.com/cafieness/db-store.git/backend/worker"

	_ "github.com/lib/pq"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Ping error:", err)
	}
	log.Println("Connected to database!")

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.PingContext(r.Context()); err != nil {
			http.Error(w, "database unavailable", http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		}); err != nil {
			log.Println("health response encode error:", err)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server Started!")
	})

	mux.HandleFunc("/analytics/top-products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
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

	mux.HandleFunc("/analytics/orders-summary", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		data, err := queries.GetOrdersSummary(db)
		if err != nil {
			log.Println(err)
			http.Error(w, "Product View Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println("JSON encode error:", err)
		}
	})
	worker.StartWorker(db)
	worker.StartSimulator(db)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
