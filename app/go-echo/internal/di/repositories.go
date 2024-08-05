package di

import (
	"github.com/google/wire"

	"mickamy.com/playground/internal/repository"
)

type Repositories struct {
	User        repository.User
	UserAccount repository.UserAccount
	UserAvatar  repository.UserAvatar
	UserProfile repository.UserProfile
}

var repositories = wire.NewSet(
	repository.NewUser,
	repository.NewUserAccount,
	repository.NewUserAvatar,
	repository.NewUserProfile,
)
