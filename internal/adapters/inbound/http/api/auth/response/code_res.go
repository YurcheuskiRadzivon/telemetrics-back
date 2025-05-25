package response

const (
	Status2FANeeded          = "2FA REQUIRED"
	StatusSuccessfully       = "SUCCESSFULLY"
	StatusAuthRestart        = "AUTH_RESTART"
	StatusCodeExpired        = "PHONE_CODE_EXPIRED"
	StatusCodeEmpty          = "PHONE_CODE_EMPTY"
	StatusPhoneNumUnnocupied = "PHONE_NUMBER_UNOCCUPIED"
)

type CodeRespose struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}
