package tests

import (
	"context"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/repositories"
	proto "github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
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
