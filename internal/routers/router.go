package routers

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/repositories"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/services/author"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/usecase"
	"github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	r := repositories.NewRepository(se.DB)

	//======================== ROUTER ========================
	//Setting Services
	//Setting Author Service

	//=== AUTHOR ===
	authorService := author.NewService(r)
	AuthorHandler := usecase.NewAuthorHandler(authorService)

	pb.RegisterAuthorServiceServer(grpcServer, AuthorHandler)
	//=========================================================
}
