package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID             uint64         `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"not null" valid:"required" json:"name"`
	Price          uint64         `gorm:"type:bigint not null" valid:"required" json:"price"`
	ExpiredAt      time.Time      `json:"expired_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	OrderHistories []OrderHistory
}

func (o *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	o.ExpiredAt = time.Now().Add(14 * 24 * time.Hour)

	_, errCreate := govalidator.ValidateStruct(o)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (o *OrderItem) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExpiredAt = time.Now().Add(12 * 24 * time.Hour)

	var currentData OrderItem
	tx.Debug().Where("id = ?", o.ID).Take(&currentData)

	if o.Name == "" {
		o.Name = currentData.Name
	}

	if o.Price == 0 {
		o.Price = currentData.Price
	}

	_, errUpdate := govalidator.ValidateStruct(o)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
