package entity

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
