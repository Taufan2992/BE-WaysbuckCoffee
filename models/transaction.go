package models

import "time"

// User model struct
type Transaction struct {
	ID int `json:"id" gorm:"primary_key:auto_increment"`
	// ProductID int                      `json:"product_id"`
	// Product   ProductUserResponse      `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int                      `json:"user_id"`
	User      UsersTransactionResponse `json:"user"`
	Price     int                      `json:"price"`
	Carts     []Cart                   `json:"carts"`
	Status    string                   `json:"status"  gorm:"type:varchar(25)"`
	CreatedAt time.Time                `json:"-"`
	UpdatedAt time.Time                `json:"-"`
}

type TransactionResponse struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Status    string               `json:"status" gorm:"type: varchar(255)"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user"`
	ProductID int                  `json:"product_id"`
	Product   ProductResponse      `json:"product"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
