package queries

import (
	"database/sql"
)

type TopProduct struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TotalSold int    `json:"total_sold"`
}

func GetTopProducts(db *sql.DB) ([]TopProduct, error) {
	rows, err := db.Query("select id, name, total_sold from top_products_30ds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []TopProduct
	for rows.Next() {
		var p TopProduct
		if err := rows.Scan(&p.ID, &p.Name, &p.TotalSold); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil

}

type ProductViews struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	TotalViews int    `json:"total_views"`
}

func GetProductViews(db *sql.DB) ([]ProductViews, error) {
	rows, err := db.Query("select id, name, view_count from product_view_summary_30ds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []ProductViews
	for rows.Next() {
		var p ProductViews
		if err := rows.Scan(&p.ID, &p.Name, &p.TotalViews); err != nil {
			return nil, err
		}
		views = append(views, p)
	}

	return views, nil

}

type Revenue struct {
	TotalRevenue      float64 `json:"total_revenue"`
	TotalSoldProducts int     `json:"total_products_sold"`
	TotalOrders       int     `json:"total_orders"`
}

func GetRevenue(db *sql.DB) ([]Revenue, error) {
	rows, err := db.Query("select total_revenue, total_products_sold, total_orders from revenue_snapshots_30ds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snapshots []Revenue
	for rows.Next() {
		var p Revenue
		if err := rows.Scan(&p.TotalRevenue, &p.TotalSoldProducts, &p.TotalOrders); err != nil {
			return nil, err
		}
		snapshots = append(snapshots, p)
	}

	return snapshots, nil

}

type OrdersSummary struct {
	TotalOrders  int     `json:"total_orders"`
	TotalRevenue float64 `json:"total_revenue"`
}

func GetOrdersSummary(db *sql.DB) ([]OrdersSummary, error) {
	rows, err := db.Query("SELECT total_orders, total_revenue FROM orders_summary")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []OrdersSummary

	for rows.Next() {
		var o OrdersSummary
		rows.Scan(&o.TotalOrders, &o.TotalRevenue)
		res = append(res, o)
	}

	return res, nil
}
