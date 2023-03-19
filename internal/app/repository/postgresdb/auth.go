package postgresdb

import (
	"database/sql"
	"errors"

	"github.com/cucumberjaye/gophermart/internal/app/handler"
	"github.com/cucumberjaye/gophermart/internal/app/models"
)

var (
	selectStmt *sql.Stmt
	insertStmt *sql.Stmt
)

func newAuthStmts(db *sql.DB) error {
	var err error

	insertStmt, err = db.Prepare("INSERT INTO users (id, login, password_hash) values ($1, $2, $3)")
	if err != nil {
		return err
	}

	selectStmt, err = db.Prepare("SELECT id, login FROM users WHERE login=$1 AND password_hash=$2")
	if err != nil {
		return err
	}

	return nil
}

func (r *Postgres) GetUser(loginUser models.LoginUser) (models.User, error) {
	var user models.User

	row := selectStmt.QueryRow(loginUser.Login, loginUser.Password)
	err := row.Scan(&user.ID, &user.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, handler.ErrorWrongLoginOrPassword
		}
		return user, err
	}
	return user, nil
}

func (r *Postgres) CreateUser(id string, user models.RegisterUser) error {
	_, err := insertStmt.Exec(id, user.Login, user.Password)
	if err != nil {
		if !errors.Is(err, sql.ErrConnDone) {
			return handler.ErrorLoginExists
		}
		return err
	}

	return nil
}
