package input

type UserSignUp struct {
	Provider string `json:"provider" validate:"required"`
	IDToken  string `json:"id_token" validate:"required"`
}
