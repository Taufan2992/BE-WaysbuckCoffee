package cartdto

import "time"

type CartRequest struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	// TransactionID int       `json:"transaction_id"`
	Qty         int       `json:"qty" form:"qty"`
	Image       string    `json:"image" form:"image"`
	ProductName string    `json:"product_name" form:"product_name"`
	SubAmount   int       `json:"subamount"`
	TopingID    []int     `json:"toping_id" form:"toping_id" gorm:"-"`
	UserID      int       `json:"user_id" gorm:"type: int"`
	CreatedAt   time.Time `json:"created_at"`
}
