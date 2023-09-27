package param

type UserRegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type UserRegisterResponse struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
