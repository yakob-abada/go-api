package repository

import "database/sql"

type DatabaseConnection interface {
	Connect() (*sql.DB, error)
}
