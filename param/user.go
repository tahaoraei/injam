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

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"access_token"`
}

type UserProfileRequest struct {
	ID int `json:"id"`
}

type UserProfileResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
