package repositories

import "database/sql"

type SalesRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}
