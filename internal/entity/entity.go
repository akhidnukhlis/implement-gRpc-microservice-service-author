package entity

type Error string

// Declare error message
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")

	ErrAuthorNotExist            = Error("domain.author.error.not_exist")
	ErrAuthorAlreadyExist        = Error("domain.author.error.email_or_username_alredy_exist")
	ErrAuthorsCredentialNotExist = Error("domain.author.error.credential_not_exist")
)

func (e Error) Error() string {
	return string(e)
}
