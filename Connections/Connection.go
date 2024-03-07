package connections

import "database/sql"

type Database struct {
	SqlDb *sql.DB
}