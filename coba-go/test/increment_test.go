package test

import (
	"context"
	"fmt"
	"project-go/libs/config/database"
	"strconv"
	"testing"
)

func TestIncrement(t *testing.T) {

	db := database.GetConnection()
	defer db.Close()

	email := "admin@gmail.com"
	comment := "Hello World"

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	InsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("InsertId: ", InsertId)
}

// query yang sama, namun data berbeda beda, bisa menggunakan prepared statement
func TestPreparedStatement(t *testing.T) {

	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	statement, err := db.PrepareContext(ctx, "INSERT INTO comments(email, comment) VALUES (?, ?)")

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "yudas" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke: " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("InsertId: ", id)
	}
}
