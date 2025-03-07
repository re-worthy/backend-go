package dto

type TGetUserRs struct {
	Username string `json:"username" validate:"required" `
	Image    string `json:"image" validate:"required" `
	Id       int  `json:"id" validate:"required" `
	Balance  int  `json:"balace" validate:"required" `
}
