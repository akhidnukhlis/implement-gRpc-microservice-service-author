package author

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/external/seeders"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/external/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_authorRepository_SaveNewAuthor(t *testing.T) {
	tt := struct {
		fields tests.Fields
		args   tests.Args
	}{}

	seedAuthor, _ := seeders.SeedAuthor(tt.args.DB)

	t.Run("if author was successfully recorded, it should not return error", func(t *testing.T) {
		err := tt.fields.Repo.Author.SaveNewAuthor(tt.args.Ctx, seedAuthor)
		assert.NoError(t, err)
	})
}
