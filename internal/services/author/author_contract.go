package author

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"

	proto "github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
)

type Service interface {
	CreateNewAuthor(ctx context.Context, payload *proto.CreateAuthorRequest) (*entity.Author, error)
	FindAuthor(ctx context.Context, authorID string) (*entity.Author, error)
}
