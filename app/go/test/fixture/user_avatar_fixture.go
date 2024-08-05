package fixture

import (
	"mickamy.com/playground/config"
	"mickamy.com/playground/internal/model"
)

func UserAvatar(setter func(m *model.UserAvatar)) model.UserAvatar {
	m := model.UserAvatar{
		Bucket: config.AWS().S3Bucket,
	}
	if setter != nil {
		setter(&m)
	}
	return m
}
