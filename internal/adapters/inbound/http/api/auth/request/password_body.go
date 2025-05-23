package request

type PasswordBody struct {
	SessionID string `json:"session_id"`
	Password  string `json:"code"`
}
