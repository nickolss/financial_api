package schemas

type ClientLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type FinancialEventRequest struct {
	FinancialName string  `json:"financial_name" binding:"required"`
	Amount        float32 `json:"amount" binding:"required"`
	ClientId      uint    `json:"client_id" binding:"required"`
}
