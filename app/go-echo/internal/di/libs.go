package di

import (
	"github.com/google/wire"

	"mickamy.com/playground/internal/lib/oauth"
)

type Libs struct {
	OAuth oauth.OAuth
}

var libs = wire.NewSet(
	oauth.New,
)
