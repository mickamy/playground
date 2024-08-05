package validator_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"mickamy.com/playground/internal/lib/validator"
)

func TestValidator_Struct(t *testing.T) {
	t.Parallel()

	// arrange
	testCases := []struct {
		name     string
		s        interface{}
		messages map[string][]string
	}{
		{
			name: "Invalid required field",
			s: struct {
				Field string `validate:"required"`
			}{
				Field: "",
			},
			messages: map[string][]string{
				"Field": {"is required."},
			},
		},
		{
			name: "Valid required field",
			s: struct {
				Field string `validate:"required"`
			}{
				Field: "required",
			},
			messages: nil,
		},
		{
			name: "Invalid min field",
			s: struct {
				Field string `validate:"min=2"`
			}{
				Field: "1",
			},
			messages: map[string][]string{
				"Field": {fmt.Sprintf("is too short. (min=%d)", 2)},
			},
		},
		{
			name: "Valid min field",
			s: struct {
				Field string `validate:"min=2"`
			}{
				Field: "12",
			},
			messages: nil,
		},
		{
			name: "Invalid max field",
			s: struct {
				Field string `validate:"max=1"`
			}{
				Field: "12",
			},
			messages: map[string][]string{
				"Field": {fmt.Sprintf("is too long. (max=%d)", 1)},
			},
		},
		{
			name: "Valid max field",
			s: struct {
				Field string `validate:"max=1"`
			}{
				Field: "1",
			},
			messages: nil,
		},
		{
			name: "Invalid email field",
			s: struct {
				Field string `validate:"email"`
			}{
				Field: "invalid-email",
			},
			messages: map[string][]string{
				"Field": {"is not a valid email."},
			},
		},
		{
			name: "Valid email field",
			s: struct {
				Field string `validate:"email"`
			}{
				Field: "valid@example.com",
			},
			messages: nil,
		},
		{
			name: "Invalid url field",
			s: struct {
				Field string `validate:"url"`
			}{
				Field: "invalid-url",
			},
			messages: map[string][]string{
				"Field": {"is not a valid URL."},
			},
		},
		{
			name: "Valid url field",
			s: struct {
				Field string `validate:"url"`
			}{
				Field: "https://example.com",
			},
			messages: nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			// act
			messages := validator.Struct(context.Background(), testCase.s)

			// assert
			assert.Equal(t, testCase.messages, messages)
		})
	}
}
