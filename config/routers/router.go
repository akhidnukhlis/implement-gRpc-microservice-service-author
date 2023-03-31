package routers

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/handler"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/src/author"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/repositories"
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
	AuthorHandler := handler.NewAuthorHandler(authorService)

	pb.RegisterAuthorServiceServer(grpcServer, AuthorHandler)
	//=========================================================

}
