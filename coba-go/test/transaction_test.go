package test

import (
	"context"
	"fmt"
	"project-go/libs/config/database"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	// do transaction
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "yudas" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke: " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		lastId, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("InsertId: ", lastId)
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
