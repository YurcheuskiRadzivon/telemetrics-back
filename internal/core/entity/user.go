package entity

type User struct {
	ID          int
	Username    string
	PhoneNumber string
	Session     []byte
}
