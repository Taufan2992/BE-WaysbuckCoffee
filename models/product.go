package models

import "time"

// User model struct
type Product struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	// UserID int    `json:"user_id" form:"user_id"`
	// User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Toping    []Toping  `json:"toping" gorm:"-"`
	TopingID  []int     `json:"-" form:"toping_id" gorm:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductResponse struct {
	ID       int                  `json:"id"`
	Title    string               `json:"title"`
	Price    int                  `json:"price"`
	Image    string               `json:"image"`
	Qty      int                  `json:"qty"`
	UserID   int                  `json:"user_id" form:"user_id"`
	User     UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Toping   []Toping             `json:"toping" gorm:"many2many:product_topings"`
	TopingID []int                `json:"-" form:"toping_id" gorm:"-"`
}

type ProductTransactionResponse struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Price    int      `json:"price"`
	Image    string   `json:"image"`
	UserID   int      `json:"-"`
	Toping   []Toping `json:"toping" gorm:"many2many:product_topings"`
	TopingID []int    `json:"-" form:"toping_id" gorm:"-"`
}

type ProductUserResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}

func (ProductTransactionResponse) TableName() string {
	return "products"
}
