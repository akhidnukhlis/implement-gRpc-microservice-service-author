package author

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/errorcodehandling"
	"github.com/jinzhu/gorm"
)

type authorRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{
		db: db,
	}
}
