package dto

type TTransactionRs struct {
	Description string `json:"description" validate:"required" `
	Currency    string `json:"currency" validate:"required" `
	ID          int64  `json:"id" validate:"required" `
	OwnerID     int64  `json:"ownerID" validate:"required" `
	Amount      int64  `json:"amount" validate:"required" `
	IsIncome    int64  `json:"IsIncome" validate:"required" `
	Createdat   int64  `json:"createdAt" validate:"required" `
}
type TTransactionWTagsRs struct {
	Tags []string `json:"tags" validate:"required"`
	TTransactionRs
}

type TTransactionRq struct {
	Description string   `json:"description" validate:"required" `
	Tags        []string `json:"tags" validate:"required"`
	Amount      int64    `json:"amount" validate:"required" `
	IsIncome    int64    `json:"isIncome" validate:"required" `
}
