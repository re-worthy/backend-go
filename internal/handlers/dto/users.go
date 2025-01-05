package dto

type TGetUserRs struct {
	Username string `json:"username" validate:"required" `
	Image    string `json:"image" validate:"required" `
	Id       int64  `json:"id" validate:"required" `
	Balance  int64  `json:"balace" validate:"required" `
}
