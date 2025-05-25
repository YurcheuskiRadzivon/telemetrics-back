package response

type PasswordResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}
