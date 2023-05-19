package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type OrderHistory struct {
	ID           uint64         `gorm:"primaryKey" json:"id"`
	UserID       uint64         `gorm:"not null" valid:"required" json:"user_id"`
	OrderItemID  uint64         `gorm:"not null" valid:"required" json:"order_item_id"`
	Descriptions string         `gorm:"not null" valid:"required" json:"descriptions"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"update_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (o *OrderHistory) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(o)
	if errCreate != nil {
		err = errCreate
		return err
	}

	return nil
}

func (o *OrderHistory) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(o)
	if errUpdate != nil {
		err = errUpdate
		return err
	}

	return nil
}
