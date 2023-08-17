package routers

import (
	"github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/repositories"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/services/author"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/usecase"
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
