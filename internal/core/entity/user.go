package entity

type User struct {
	UserID      int
	Username    string
	PhoneNumber string
	Session     []byte
}
