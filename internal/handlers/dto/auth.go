package dto

type TLoginRq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TRegisterRq struct {
	Username string `json:"username" validate:"required" `
	Image    string `json:"image" `
	Password string `json:"password" validate:"required" `
}

type TAuthRs struct {
	Token string `json:"token" validate:"required"`
}
