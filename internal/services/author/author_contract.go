package author

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/entity"

	"github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"
)

type Service interface {
	CreateNewAuthor(ctx context.Context, payload *pb.CreateAuthorRequest) (*entity.Author, error)
	FindAuthor(ctx context.Context, authorID string) (*entity.Author, error)
}
