package manage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbHost   = "localhost"
	dbName   = "BOOTBot_db"
	dbPass   = "BOOTBot_pass"
	dbUser   = "BOOTBot_user"
)

func Connection() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", dbUser, dbName, dbPass, dbHost))
	if err != nil {
		return nil, err
	}
	return db, nil
}
