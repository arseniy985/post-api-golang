package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Database struct {
	dialect  string
	host     string
	name     string
	user     string
	password string
	port     int
}

var DB *sql.DB

func NewDatabase(dialect, host, name, user, password string, port int) *Database {
	return &Database{
		dialect:  dialect,
		host:     host,
		name:     name,
		user:     user,
		password: password,
		port:     port,
	}
}

func (database *Database) OpenConnection() error {
	db, err := sql.Open(database.dialect, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", database.user, database.password, database.host, database.port, database.name))
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
	return nil
}

func (database *Database) CloseConnection() error {
	if err := DB.Close(); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
