package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
}
type UserResponseDelete struct {
	ID int `json:"id"`
}
type UserResponseUpdate struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}
