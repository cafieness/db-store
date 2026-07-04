package worker

import (
	"database/sql"
	"log"
	"time"
)

func StartWorker(db *sql.DB) {
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			refreshViews(db)
		}
	}()
}

func refreshViews(db *sql.DB) {
	_, err := db.Exec("refresh materialized view top_products_30ds")
	if err != nil {
		log.Println("top_products refresh error:", err)
	}

	_, err = db.Exec("refresh materialized view revenue_snapshots_30ds")
	if err != nil {
		log.Println("revenue refresh error:", err)
	}

	log.Println("materialized views refreshed")
}
