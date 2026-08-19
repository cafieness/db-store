package queries

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetTopProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"total_sold",
	}).
		AddRow(1, "Product 1", 15).
		AddRow(2, "Product 2", 10)

	mock.ExpectQuery(
		"select id, name, total_sold from top_products_30ds",
	).WillReturnRows(rows)

	products, err := GetTopProducts(db)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}

	if products[0].ID != 1 {
		t.Errorf("expected id 1, got %d", products[0].ID)
	}

	if products[0].Name != "Product 1" {
		t.Errorf("expected Product 1, got %s", products[0].Name)
	}

	if products[0].TotalSold != 15 {
		t.Errorf("expected total sold 15, got %d", products[0].TotalSold)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sql expectations: %v", err)
	}
}

func TestGetRevenue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"total_revenue",
		"total_products_sold",
		"total_orders",
	}).AddRow(15000.50, 25, 8)

	mock.ExpectQuery(
		"select total_revenue, total_products_sold, total_orders from revenue_snapshots_30ds",
	).WillReturnRows(rows)

	revenue, err := GetRevenue(db)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(revenue) != 1 {
		t.Fatalf("expected 1 revenue snapshot, got %d", len(revenue))
	}

	if revenue[0].TotalRevenue != 15000.50 {
		t.Errorf(
			"expected revenue 15000.50, got %f",
			revenue[0].TotalRevenue,
		)
	}

	if revenue[0].TotalSoldProducts != 25 {
		t.Errorf(
			"expected 25 sold products, got %d",
			revenue[0].TotalSoldProducts,
		)
	}

	if revenue[0].TotalOrders != 8 {
		t.Errorf(
			"expected 8 orders, got %d",
			revenue[0].TotalOrders,
		)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sql expectations: %v", err)
	}
}

func TestGetTopProductsQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer db.Close()

	expectedErr := errors.New("database error")

	mock.ExpectQuery(
		"select id, name, total_sold from top_products_30ds",
	).WillReturnError(expectedErr)

	products, err := GetTopProducts(db)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if products != nil {
		t.Errorf("expected nil products, got %v", products)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sql expectations: %v", err)
	}
}
