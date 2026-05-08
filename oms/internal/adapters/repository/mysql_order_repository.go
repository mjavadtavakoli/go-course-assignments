package repository

import (
	"database/sql"
	"errors"

	"oms/internal/core/domain"
)

type MySQLOrderRepository struct {
	db *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) *MySQLOrderRepository {
	return &MySQLOrderRepository{db: db}
}

func (r *MySQLOrderRepository) Create(order *domain.Order) error {
	query := `INSERT INTO orders (customer_name, status, created_at) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, order.CustomerName, order.Status, order.CreatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	order.ID = id
	return nil
}

func (r *MySQLOrderRepository) GetByID(id int64) (*domain.Order, error) {
	query := `SELECT id, customer_name, status, created_at FROM orders WHERE id = ?`

	row := r.db.QueryRow(query, id)

	order := &domain.Order{}
	err := row.Scan(&order.ID, &order.CustomerName, &order.Status, &order.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return order, nil
}
