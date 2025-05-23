package response

const (
	StatusCodeRequested  = "CODE_REQUESTED"
	StatusInvalidRequest = "INVALID_REQUEST"
)

type StartResponse struct {
	SessionID string `json:"session_id"`
	Status    string `json:"status`
}
