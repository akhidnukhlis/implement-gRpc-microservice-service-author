package usecase

import (
	"context"

	"github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/services/author"
)

type AuthorHandler struct {
	pb.UnimplementedAuthorServiceServer
	service author.Service
}

func NewAuthorHandler(service author.Service) *AuthorHandler {
	return &AuthorHandler{
		service: service,
	}
}

// ServiceRegisterAuthor is handler function to Handle author registration
func (ah *AuthorHandler) ServiceRegisterAuthor(ctx context.Context, payload *pb.CreateAuthorRequest) (*pb.AuthorStatusResponse, error) {

	_, err := ah.service.CreateNewAuthor(ctx, payload)
	if err != nil {

		return nil, err
	}

	response := &pb.AuthorStatusResponse{
		Status:  "Success",
		Message: "Success to register author",
	}

	return response, nil
}

// ServiceFindAuthorById is handler function to Handle find author
func (ah *AuthorHandler) ServiceFindAuthorById(ctx context.Context, proto *pb.FindAuthorByIdRequest) (*pb.AuthorStatusResponse, error) {

	findAuthor, err := ah.service.FindAuthor(ctx, proto.Id)
	if err != nil {

		return nil, err

	}

	author := &pb.AuthorResponse{
		Id:        findAuthor.ID,
		Name:      findAuthor.Name,
		Nickname:  findAuthor.Nickname,
		Email:     findAuthor.Email,
		CreatedAt: findAuthor.CreatedAt.String(),
		UpdatedAt: findAuthor.UpdatedAt.String(),
	}

	response := &pb.AuthorStatusResponse{
		Status:  "Success",
		Message: "Author Found",
		Data:    author,
	}

	return response, nil
}
