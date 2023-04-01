package tests

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"
	"time"
)

// FakeAuthor initiate fake data author
func FakeAuthor() *entity.Author {
	timeNow := time.Now()

	return &entity.Author{
		ID:        AuthorID001,
		Name:      AuthorUsername001,
		Nickname:  AuthorNickname001,
		Email:     AuthorEmail001,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

// // FakeUserInput initiate fake data for payload create or register new author
// func FakeUserInput() *userentity.UserRequest {
// 	fakeUser := &userentity.UserRequest{
// 		Firstname: AuthorFirstName001,
// 		Lastname:  AuthorLastName001,
// 		Phone:     AuthorPhoneNumber001,
// 		Avatar:    AuthorAvatar001,
// 		Email:     AuthorEmail001,
// 		Username:  AuthorUsername001,
// 		Password:  AuthorID001,
// 		Status:    AuthorID001,
// 	}

// 	return fakeUser
// }

// // FakeUserUpdate initiate fake payload for update data author
// func FakeUserUpdate() *userentity.UserData {
// 	t := time.Now()
// 	fakeUser := &userentity.UserData{
// 		ID:        AuthorID001,
// 		Firstname: UserFirstName002,
// 		Lastname:  UserLastName002,
// 		Phone:     AuthorPhoneNumber001,
// 		Avatar:    UserAvatar003,
// 		Email:     AuthorEmail001,
// 		Username:  UserUsername002,
// 		Status:    AuthorStatus001,
// 		UpdatedAt: t.String(),
// 	}

// 	return fakeUser
// }
