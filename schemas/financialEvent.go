package schemas

import "gorm.io/gorm"

type FinancialEvent struct {
	gorm.Model
	FinancialName string  `json:"financial_name" binding:"required"`
	Amount        float32 `json:"amount" binding:"required"`
	ClientId      uint    `json:"client_id" binding:"required"`
	Client        Client  `gorm:"foreignKey:ClientId" json:"-"`
}
