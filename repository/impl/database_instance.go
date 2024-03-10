package impl

import "database/sql"

var databaseInst *sql.DB

func SetDatabaseInstance(dbInst *sql.DB) {
	databaseInst = dbInst
}
