package postgresdb

import "database/sql"

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) (*Postgres, error) {
	err := newAuthStmts(db)
	return &Postgres{
		db: db,
	}, err
}
