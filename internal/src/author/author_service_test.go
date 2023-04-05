package author

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/external/faker"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/external/seeders"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/external/tests"
	"github.com/akhidnukhlis/implement-gRpc-microservice/grpc/pb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_service_CreateNewAuthor(t *testing.T) {
	tt := struct {
		fields tests.Fields
		args   tests.Args
	}{}

	seedAuthor, _ := seeders.SeedAuthor(tt.args.DB)

	t.Run("if author was successfully recorded, it should not return error", func(t *testing.T) {
		s := &service{
			repo: tt.fields.Repo,
			err:  tt.fields.Err,
		}

		result, err := s.CreateNewAuthor(tt.args.Ctx, &pb.CreateAuthorRequest{
			Name:     seedAuthor.Name,
			Nickname: seedAuthor.Nickname,
			Email:    seedAuthor.Email,
		})
		assert.NoError(t, err)
		assert.Equal(t, result, 1)
	})
}

func Test_service_FindAuthor(t *testing.T) {
	fakerAuthor := faker.FakeAuthor()

	tt := struct {
		fields tests.Fields
		args   tests.Args
	}{}

	t.Run("if author was exists, it should not return error", func(t *testing.T) {
		s := &service{
			repo: tt.fields.Repo,
			err:  tt.fields.Err,
		}

		result, err := s.FindAuthor(tt.args.Ctx, fakerAuthor.ID)
		assert.NoError(t, err)
		assert.Equal(t, fakerAuthor.ID, result.ID)
	})
}
