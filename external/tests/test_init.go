package tests

import (
	"context"
	proto "github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/repositories"
	"github.com/jinzhu/gorm"
)

type Fields struct {
	Repo *repositories.Repository
	Err  *errorcodehandling.CodeError
}

type Args struct {
	Ctx     context.Context
	Payload *proto.CreateAuthorRequest
	DB      *gorm.DB
}
