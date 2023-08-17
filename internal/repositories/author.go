package repositories

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/entity"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/repositories/author"

	"github.com/jinzhu/gorm"
)

type Author interface {
	SaveNewAuthor(ctx context.Context, payload *entity.Author) error
	FindAuthorByID(ctx context.Context, authorID string) (*entity.Author, error)
}

func NewAuthor(db *gorm.DB) Author {
	return author.NewAuthorRepository(db)
}
