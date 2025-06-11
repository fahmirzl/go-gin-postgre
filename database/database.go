package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host		= "localhost"
	port		= 5432
	user		= "postgres"
	password	= "root"
	dbname		= "go-sql"
)

var (
	db		*sql.DB
	err		error
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Koneksi database sukses ...")
	return db, err
}