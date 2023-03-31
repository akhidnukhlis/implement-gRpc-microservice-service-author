package entity

type Error string

// Declare error message
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")

	ErrUserNotExist            = Error("domain.author.error.not_exist")
	ErrUserAlreadyExist        = Error("domain.author.error.email_or_username_alredy_exist")
	ErrUsersCredentialNotExist = Error("domain.author.error.credential_not_exist")
)

func (e Error) Error() string {
	return string(e)
}
