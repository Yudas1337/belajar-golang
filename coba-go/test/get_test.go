package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"project-go/libs/config/database"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id,name FROM customers"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)

	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM customer")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil {
			//panic(err)
			log.Fatal(err)
		}

		fmt.Println("===========================")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date: ", birthDate.Time)

		}
		fmt.Println("Created At: ", createdAt)
		fmt.Println("Married: ", married)
	}

}
