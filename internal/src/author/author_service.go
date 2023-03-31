package author

import (
	"context"
	proto "github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
	"time"

	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/unique"

	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/repositories"
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
	hashPassword, _ := helpers.HashPassword(payload.Password)

	user := &entity.Author{
		ID:        uuid.NewString(),
		Name:      payload.Name,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  hashPassword,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err = s.repo.Author.SaveNewAuthor(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindAuthor represents algorithm to find author by id
func (s *service) FindAuthor(ctx context.Context, userID string) (*entity.Author, error) {
	if err := unique.ValidateUUID(userID); err != nil {
		return nil, entity.ErrUserNotExist
	}

	user, err := s.repo.Author.FindAuthorByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
