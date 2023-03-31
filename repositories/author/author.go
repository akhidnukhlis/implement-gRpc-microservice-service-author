package author

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/errorcodehandling"
	"github.com/jinzhu/gorm"
)

type AuthorRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}
