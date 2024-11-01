package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func OpenConnection() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/belajar_golang")
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
}

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
