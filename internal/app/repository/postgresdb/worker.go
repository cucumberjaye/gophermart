package postgresdb

import (
	"database/sql"
	"time"

	"github.com/cucumberjaye/gophermart/internal/app/models"
)

func (r *Postgres) GetWaitingOrders() ([]string, error) {
	query := "SELECT id FROM orders WHERE status<2"
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var orders = []string{}

	for row.Next() {
		var id string
		err := row.Scan(&id)
		if err != nil {
			return nil, err
		}

		orders = append(orders, id)
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	return orders, nil
}

func (r *Postgres) UpdateOrder(order models.Order) error {
	query := "UPDATE orders SET status=$1, accrual=$2, uploaded_at=$3 WHERE id=$4"
	_, err := r.db.Exec(query, order.Status, order.Accrual, time.Now(), order.ID)
	return err
}
