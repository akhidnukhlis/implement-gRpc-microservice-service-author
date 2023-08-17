package entity

import (
	proto "github.com/akhidnukhlis/implement-gRpc-proto-bank/grpc/pb"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Author -> initialization author entity
type Author struct {
	ID        string    `gorm:"size:36;not null;unique index;primaryKey"`
	Name      string    `gorm:"size:255;"`
	Nickname  string    `gorm:"size:100;unique"`
	Email     string    `gorm:"size:100;unique"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// AuthorData represent body for get data from
type AuthorData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"size:100;unique"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthorRequest is payload for register author
type AuthorRequest struct {
	Name      string `json:"name"`
	Nickname  string `json:"size:100;unique"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthorRequestValidate is to validate input request
func AuthorRequestValidate(ar *proto.CreateAuthorRequest) error {
	err := validation.Errors{
		"name":     validation.Validate(&ar.Name, validation.Required, validation.Length(2, 40)),
		"nickname": validation.Validate(&ar.Nickname, validation.Required),
		"email":    validation.Validate(&ar.Email, validation.Required),
	}

	return err.Filter()
}
