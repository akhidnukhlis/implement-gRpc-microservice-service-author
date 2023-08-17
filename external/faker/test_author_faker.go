package faker

import (
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/entity"
	"time"
)

// FakeAuthor initiate fake data author
// FakeAuthor initiate fake data author
func FakeAuthor() *entity.Author {
	timeNow := time.Now()
	return &entity.Author{
		ID:        AuthorID,
		Name:      AuthorFirstName + AuthorLastName,
		Nickname:  AuthorNickname,
		Email:     AuthorEmail,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

// FakeAuthorRequest initiate fake data author request
func FakeAuthorRequest() *entity.AuthorRequest {
	return &entity.AuthorRequest{
		Name:     AuthorFirstName + AuthorLastName,
		Nickname: AuthorNickname,
		Email:    AuthorEmail,
	}
}
