package request

type CodeBody struct {
	SessionID string `json:"session_id"`
	Code      string `json:"code"`
}
