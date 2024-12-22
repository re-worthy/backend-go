package dto

type THelloWorld struct {
	Hello string `json:"hello"`
}

type THelloWorldRq struct {
	Name string `json:"name" validate:"required" `
}

type THelloDB struct {
	Counter int32 `json:"counter"`
}
