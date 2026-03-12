package user

import "database/sql"

type Repository struct {
	Client *sql.DB
}
