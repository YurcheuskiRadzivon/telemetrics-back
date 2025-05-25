package request

type CodeBody struct {
	ManageSessionID string `json:"session_id"`
	Code            string `json:"code"`
}
