package order

import (
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(order *Order) (*Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return order, err
}

func (r *repository) Get() (*[]Order, error) {
	var orders []Order

	err := r.db.Preload("Item").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (r *repository) Delete(orderId uint) error {
	var order Order

	if err := r.db.First(&order, orderId).Error; err != nil {
		return err
	}

	// Delete record item
	r.db.Where("order_id=?", orderId).Delete(&Item{})
	err := r.db.Delete(&order)
	if err.RowsAffected == 0 {
		return fmt.Errorf("null")
	}

	return nil
}

func (r *repository) Update(orderId uint, order *Order) (*Order, error) {

}
