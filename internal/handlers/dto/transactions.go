package dto

type TTransactionRs struct {
	Description string `json:"description" validate:"required" `
	Currency    string `json:"currency" validate:"required" `
	ID          int    `json:"id" validate:"required" `
	OwnerID     int    `json:"ownerID" validate:"required" `
	Amount      int    `json:"amount" validate:"required" `
	IsIncome    int    `json:"IsIncome" validate:"required" `
	Createdat   int    `json:"createdAt" validate:"required" `
}
type TTransactionWTagsRs struct {
	Tags []string `json:"tags" validate:"required"`
	TTransactionRs
}

type TTransactionRq struct {
	Description string   `json:"description" validate:"required" `
	Tags        []string `json:"tags" validate:"required"`
	Amount      int      `json:"amount" validate:"required" `
	IsIncome    int      `json:"isIncome" validate:"required" `
}
