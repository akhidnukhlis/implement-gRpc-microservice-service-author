package author

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/repositories"
	proto "github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
	"time"

	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/unique"

	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"
	"github.com/google/uuid"
)

type service struct {
	repo *repositories.Repository
	err  *errorcodehandling.CodeError
}

func NewService(repo *repositories.Repository) *service {
	return &service{
		repo: repo,
	}
}

// CreateNewAuthor represents algorithm to register new author
func (s *service) CreateNewAuthor(ctx context.Context, payload *proto.CreateAuthorRequest) (*entity.Author, error) {

	err := entity.AuthorRequestValidate(payload)
	if err != nil {
		return nil, err
	}

	author := &entity.Author{
		ID:        uuid.NewString(),
		Name:      payload.Name,
		Email:     payload.Email,
		Nickname:  payload.Nickname,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err = s.repo.Author.SaveNewAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

// FindAuthor represents algorithm to find author by id
func (s *service) FindAuthor(ctx context.Context, authorID string) (*entity.Author, error) {
	if err := unique.ValidateUUID(authorID); err != nil {
		return nil, entity.ErrAuthorNotExist
	}

	author, err := s.repo.Author.FindAuthorByID(ctx, authorID)
	if err != nil {
		return nil, err
	}

	return author, nil
}
