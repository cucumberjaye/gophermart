package postgresdb

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/cucumberjaye/gophermart/internal/app/handler"
	"github.com/cucumberjaye/gophermart/internal/app/models"
)

func (r *Postgres) SetOrder(order models.Order) error {
	var id string
	userQuery := "SELECT id FROM orders WHERE id=$1 and user_id=$2"
	row := r.db.QueryRow(userQuery, order.ID, order.UserID)
	err := row.Scan(&id)
	if err == nil {
		return handler.ErrUserOrderExists
	}
	if errors.Is(err, sql.ErrNoRows) {
		allQuery := "SELECT id FROM orders WHERE id=$1"
		row = r.db.QueryRow(allQuery, order.ID)
		err = row.Scan(&id)
		if err == nil {
			return handler.ErrOrderExists
		}
		if errors.Is(err, sql.ErrNoRows) {
			insertQuery := "INSERT INTO orders (id, user_id, status, uploaded_at) VALUES ($1, $2, $3, $4)"
			_, err = r.db.Exec(insertQuery, order.ID, order.UserID, order.Status, order.UploadedAt)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return err
}

func (r *Postgres) GetOrders(userID string) ([]models.Order, error) {
	query := "SELECT * FROM orders WHERE user_id=$1 ORDER BY uploaded_at DESC"
	row, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var orders = []models.Order{}

	for row.Next() {
		var tmp models.Order
		var accrualSQL sql.NullInt32
		err := row.Scan(&tmp.ID, &tmp.UserID, &tmp.Status, &accrualSQL, &tmp.UploadedAt)
		if err != nil {
			return nil, err
		}

		if accrualSQL.Valid {
			tmp.Accrual = int(accrualSQL.Int32)
		}

		fmt.Println("---", tmp)
		orders = append(orders, tmp)
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}

	return orders, nil
}
