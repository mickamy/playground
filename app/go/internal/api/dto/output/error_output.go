package output

import (
	"mickamy.com/playground/internal/lib/validator"
)

type errors = map[string][]string

type Error struct {
	Message string `json:"message,omitempty"`
	Errors  errors `json:"errors,omitempty"`
}

var ErrorDefault = Error{
	Message: "500 Internal Server Error: 時間をおいて再度お試しください。",
}

func NewErrorMessage(err error) Error {
	return Error{
		Message: err.Error(),
	}
}

func NewErrorMessages(messages validator.ErrorMessagesType) Error {
	return Error{
		Errors: messages,
	}
}
