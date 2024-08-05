package validator

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	lib "github.com/go-playground/validator/v10"
)

var validator = lib.New()

type ErrorMessagesType = map[string][]string

func Struct(ctx context.Context, s interface{}) ErrorMessagesType {
	if err := validator.StructCtx(ctx, s); err != nil {
		errs := lib.ValidationErrors{}
		if errors.As(err, &errs) {
			slog.Error("validationErrors", "err", errs)
			return mapErrorMessages(errs)
		}
		panic(fmt.Errorf("unknown validation error %w", err))
	}
	return nil
}

func mapErrorMessages(errs lib.ValidationErrors) ErrorMessagesType {
	messages := map[string][]string{}
	for _, err := range errs {
		var message string
		switch err.ActualTag() {
		case "required":
			message = fmt.Sprintf("is required.")
		case "min":
			message = fmt.Sprintf("is too short. (min=%s)", err.Param())
		case "max":
			message = fmt.Sprintf("is too long. (max=%s)", err.Param())
		case "email":
			message = fmt.Sprintf("is not a valid email.")
		case "url":
			message = fmt.Sprintf("is not a valid URL.")
		default:
			slog.Warn("failing back to default error message.", "ActualTag", err.ActualTag(), "Tag", err.Tag(), "Param", err.Param())
			message = fmt.Sprintf("is invalid.")
		}
		field := err.Field()
		if val, ok := messages[field]; ok {
			// Currently, validator package does not provide feature to return multiple violations at once.
			// But here's a PR to achieve it.
			// https://github.com/go-playground/validator/pull/1288
			messages[field] = append(val, message)
		} else {
			messages[field] = []string{message}
		}
	}
	return messages
}
