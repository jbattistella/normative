package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var dbQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:pJGlBJilIdmLHvJIIFfq@containers-us-west-107.railway.app:6131/railway"
)

var DB *sql.DB

func ConnectDB() (*Queries, error) {

	DB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return nil, err

	}

	dbQueries = New(DB)

	return dbQueries, nil
}
