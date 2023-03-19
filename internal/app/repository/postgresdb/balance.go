package postgresdb

import (
	"database/sql"
	"errors"
	"time"

	"github.com/cucumberjaye/gophermart/internal/app/models"
)

func (r *Postgres) GetBalance(userID string) (models.Balance, error) {
	var balance models.Balance

	var accrualSql sql.NullInt32
	accrualQuery := "SELECT SUM(accrual) FROM orders WHERE accrual > 0 and user_id=$1"
	err := r.db.QueryRow(accrualQuery, userID).Scan(&accrualSql)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return balance, err
	}
	if accrualSql.Valid {
		balance.Current = int(accrualSql.Int32)
	}

	var withdrawSql sql.NullInt32
	withdrawQuery := "SELECT SUM(accrual) FROM orders WHERE accrual < 0 and user_id=$1"
	err = r.db.QueryRow(withdrawQuery, userID).Scan(&withdrawSql)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return balance, err
	}
	if withdrawSql.Valid {
		balance.Withdrawn = int(withdrawSql.Int32)
	}

	return balance, nil
}

func (r *Postgres) Withdraw(userID string, withdraw models.Withdraw) error {
	query := "INSERT INTO orders (id, user_id, status, accrual, uploaded_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, withdraw.Order, userID, models.Invalid, withdraw.Sum, time.Now())

	return err
}

func (r *Postgres) GetWithdrawals(userID string) ([]models.Withdraw, error) {

	query := "SELECT id, accrual, uploaded_at FROM orders WHERE user_id=$1 AND accrual<0 ORDER BY uploaded_at DESC"
	row, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var withdraw []models.Withdraw = []models.Withdraw{}

	for row.Next() {
		var tmp models.Withdraw
		var sumSQl sql.NullInt32
		err := row.Scan(&tmp.Order, &sumSQl, &tmp.ProcessedAt)
		if err != nil {
			return nil, err
		}
		if sumSQl.Valid {
			tmp.Sum = int(sumSQl.Int32)
		}

		withdraw = append(withdraw, tmp)
	}

	if len(withdraw) == 0 {
		return nil, sql.ErrNoRows
	}

	return withdraw, nil
}