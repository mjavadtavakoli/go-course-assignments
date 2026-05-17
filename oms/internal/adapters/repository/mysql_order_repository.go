package repository

import (
	"database/sql"
	"oms/internal/domain"
	"oms/internal/logger"
)

type MySQLOrderRepository struct {
	db *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) *MySQLOrderRepository {
	return &MySQLOrderRepository{
		db: db,
	}
}

func (r *MySQLOrderRepository) Create(order *domain.Order) error {

	query := `
	INSERT INTO orders(customer_name, amount)
	VALUES (?, ?)
	`

	result, err := r.db.Exec(
		query,
		order.CustomerName,
		order.Amount,
	)

	if err != nil {
		logger.Log.Printf(
			"failed to insert order into mysql: error=%v",
			err,
		)

		return err
	}

	id, err := result.LastInsertId()
	if err != nil {

		logger.Log.Printf(
			"failed to get last insert id: error=%v",
			err,
		)

		return err
	}

	order.ID = id

	logger.Log.Printf(
		"order created successfully in mysql: id=%d customer=%s amount=%.2f",
		order.ID,
		order.CustomerName,
		order.Amount,
	)

	return nil
}

func (r *MySQLOrderRepository) GetByID(id int64) (*domain.Order, error) {

	query := `
	SELECT id, customer_name, amount
	FROM orders
	WHERE id = ?
	`

	row := r.db.QueryRow(query, id)

	order := &domain.Order{}

	err := row.Scan(
		&order.ID,
		&order.CustomerName,
		&order.Amount,
	)

	if err != nil {
		return nil, err
	}

	order.Statuses = []string{"CREATED"}

	return order, nil
}

func (r *MySQLOrderRepository) GetAll() ([]domain.Order, error) {
	query := `
	SELECT id, customer_name, amount
	FROM orders
	ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.Amount); err != nil {
			return nil, err
		}
		o.Statuses = []string{"CREATED"}
		orders = append(orders, o)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *MySQLOrderRepository) Update(order *domain.Order) error {

	query := `
	UPDATE orders
	SET customer_name = ?, amount = ?
	WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		order.CustomerName,
		order.Amount,
		order.ID,
	)

	if err != nil {

		logger.Log.Printf(
			"failed to update order: id=%d error=%v",
			order.ID,
			err,
		)

		return err
	}

	logger.Log.Printf(
		"order updated successfully: id=%d customer=%s amount=%.2f",
		order.ID,
		order.CustomerName,
		order.Amount,
	)

	return nil
}
