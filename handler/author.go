package handler

import (
	"context"

	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/src/author"
	"github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
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

	findUser, err := ah.service.FindAuthor(ctx, proto.Id)
	if err != nil {

		return nil, err

	}

	user := &pb.AuthorResponse{
		Id:        findUser.ID,
		Name:      findUser.Name,
		Email:     findUser.Email,
		Username:  findUser.Username,
		Password:  findUser.Password,
		CreatedAt: findUser.CreatedAt.String(),
		UpdatedAt: findUser.UpdatedAt.String(),
	}

	response := &pb.AuthorStatusResponse{
		Status:  "Success",
		Message: "Author Found",
		Data:    user,
	}

	return response, nil
}
