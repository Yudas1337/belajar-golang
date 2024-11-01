package test

import (
	"context"
	"fmt"
	"project-go/libs/config/database"
	"testing"
)

func TestInjectionSafe(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? and password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParam(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()
	username := "admin'; #"
	password := "salah"

	ctx := context.Background()

	query := "INSERT INTO user (username, password) VALUES (?, ?)"

	_, err := db.ExecContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses tambah data")
}
