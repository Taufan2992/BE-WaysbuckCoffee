package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart, ID int) (models.Cart, error)
	FindCartTopings(TopingID []int) ([]models.Toping, error)
	FindUserCarts(ID int) ([]models.Cart, error)
	GetTransactionID() (models.Transaction, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

// GET ALL CARTS
func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Toping").Preload("Product").Preload("User").Find(&carts).Error

	return carts, err
}

// GET BY ID
func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Toping").Preload("Product").First(&cart, ID).Error

	return cart, err
}

// CREATE CART
func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Preload("Toping").Create(&cart).Error

	return cart, err
}

// UPDATE CART
func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

// DELETE
func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

// CART TOPING
func (r *repository) FindCartTopings(TopingID []int) ([]models.Toping, error) {
	var topings []models.Toping
	err := r.db.Debug().Find(&topings, TopingID).Error

	return topings, err
}

// CART USER
func (r *repository) FindUserCarts(UserID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Toping").Preload("Product").Find(&carts, "user_id  = ?", UserID).Error

	return carts, err
}

func (r *repository) GetTransactionID() (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Find(&transaction, "status = ?", "waiting").Error
	return transaction, err
}
