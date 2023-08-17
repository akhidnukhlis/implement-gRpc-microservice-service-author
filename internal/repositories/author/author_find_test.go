package author_test

import (
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/external/seeders"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/external/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorRepository_FindAuthorByID(t *testing.T) {
	tt := struct {
		fields tests.Fields
		args   tests.Args
	}{}

	seedAuthor, _ := seeders.SeedAuthor(tt.args.DB)

	t.Run("if author was exists, it should not return error", func(t *testing.T) {
		result, err := tt.fields.Repo.Author.FindAuthorByID(tt.args.Ctx, seedAuthor.ID)
		assert.NoError(t, err)
		assert.Equal(t, seedAuthor.ID, result.ID)
	})
}
