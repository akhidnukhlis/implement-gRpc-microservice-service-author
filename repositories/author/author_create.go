package author

import (
	"context"

	"github.com/pkg/errors"

	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/config/app"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"
)

// SaveNewAuthor is used to run query insert
func (r *AuthorRepository) SaveNewAuthor(ctx context.Context, payload *entity.Author) error {

	err := r.db.Create(payload).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return entity.ErrAuthorNotExist
		case app.ErrUniqueViolation:
			return entity.ErrAuthorAlreadyExist
		default:
			return errors.Wrap(parsed, "build statement query to insert author from database")
		}
	}
	return nil
}
