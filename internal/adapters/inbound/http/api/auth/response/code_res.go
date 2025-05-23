package response

const (
	Status2FANeeded    = "2FA REQUIRED"
	StatusSuccessfully = "SUCCESSFULLY"
	StatusInvalidCode  = "INVALID_CODE"
)

type CodeRespose struct {
	Status string `json:"status"`
}
