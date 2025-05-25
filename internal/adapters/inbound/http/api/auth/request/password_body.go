package request

type PasswordBody struct {
	ManageSessionID string `json:"session_id"`
	Password        string `json:"password"`
}
