package author

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"

	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/config/app"
	"github.com/pkg/errors"
)

// FindAuthorByID is used to run query find author
func (r *AuthorRepository) FindAuthorByID(ctx context.Context, authorID string) (*entity.Author, error) {
	var author entity.Author
	err := r.db.First(&author, "id = ?", authorID).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return nil, app.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to find author from database")
		}
	}
	return &author, nil
}
