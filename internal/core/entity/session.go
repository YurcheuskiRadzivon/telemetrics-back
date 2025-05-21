package entity

type Session struct {
	PhoneNumber   string `json:"phone_number"`
	PhoneCodeHash string `json:"phone_code_hash"`
	Status        string `json:"status"`
}
