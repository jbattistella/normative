package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var testQueries *Queries

var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql:///test_normative?sslmode=disable"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("dropdb", "--if-exists", "test_normative")
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to drop")
		log.Fatal(err)
	}
	cmd = exec.Command("createdb", "test_normative")
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to create")
		log.Fatal(err)
	}

	mig, err := migrate.New(
		"file://../migrations",
		dbSource)
	if err != nil {
		log.Fatal(err)
	}
	if err := mig.Up(); err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
