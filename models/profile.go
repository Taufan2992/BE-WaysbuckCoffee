package models

import "time"

type Profile struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Phone     string               `json:"phone" gorm:"type: varchar(255)"`
	Gender    string               `json:"gender" gorm:"type: varchar(255)"`
	Address   string               `json:"address" gorm:"type: text"`
	PostCode  string               `json:"post_code" gorm:"type: text"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

// for association relation with another table (user)
type ProfileResponse struct {
	ID       int                  `json:"id"`
	Phone    string               `json:"phone"`
	Gender   string               `json:"gender"`
	Address  string               `json:"address"`
	PostCode string               `json:"post_code"`
	UserID   int                  `json:"user_id"`
	User     UsersProfileResponse `json:"-"`
}

type ProfileResponseUser struct {
	ID       int                  `json:"id"`
	Phone    string               `json:"phone"`
	Gender   string               `json:"gender"`
	Address  string               `json:"address"`
	PostCode string               `json:"post_code"`
	UserID   int                  `json:"user_id"`
	User     UsersProfileResponse `json:"user"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
