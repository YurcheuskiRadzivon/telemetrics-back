package response

const (
	StatusCodeRequested = "CODE_REQUESTED"
)

type StartResponse struct {
	ManageSessionID string `json:"session_id"`
	Status          string `json:"status`
}
