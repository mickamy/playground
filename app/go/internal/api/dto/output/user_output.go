package output

type UserAccountSignUp struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
	UID      string `json:"uid"`
}
