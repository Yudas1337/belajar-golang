package repositories

import (
	"context"
	"project-go/contracts/entities"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entities.Comment) (entities.Comment, error)
	FindById(ctx context.Context, id int32) (entities.Comment, error)
	FindAll(ctx context.Context) ([]entities.Comment, error)
}
