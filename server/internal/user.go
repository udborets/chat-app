package internal

type UserSignUpInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserSignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
