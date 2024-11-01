package test

import (
	"context"
	"fmt"
	"project-go/contracts/entities"
	"project-go/contracts/repositories"
	"project-go/libs/config/database"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := repositories.NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment := entities.Comment{
		Email:   "admin@gmail.com",
		Comment: "Hello World",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := repositories.NewCommentRepository(database.GetConnection())

	result, err := commentRepository.FindById(context.Background(), 25)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := repositories.NewCommentRepository(database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
