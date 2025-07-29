package schemas

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name           string           `json:"name" binding:"required"`
	Email          string           `json:"email" binding:"required"`
	Password       string           `json:"password" binding:"required"`
	Balance        float32          `json:"amount"`
	FinancialEvent []FinancialEvent `gorm:"foreignKey:ClientId" json:"-"`
}
