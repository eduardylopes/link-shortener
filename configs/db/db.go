package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase() (*Database, error) {
	connURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connURL)

	if err != nil {
		return nil, err
	}

	return &Database{Db: db}, nil
}

func (d *Database) Close() {
	d.Db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.Db
}
