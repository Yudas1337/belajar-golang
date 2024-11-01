package test

import (
	"context"
	"fmt"
	"project-go/libs/config/database"
	"testing"
)

func TestInsert(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer (id, name) VALUES ('123', 'John Doe')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new customer")
}
