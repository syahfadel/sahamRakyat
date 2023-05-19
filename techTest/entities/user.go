package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID             uint64         `gorm:"primaryKey" json:"id"`
	FullName       string         `gorm:"not null" valid:"required" json:"full_name"`
	FirstOrder     *time.Time     `json:"first_order"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	OrderHistories []OrderHistory
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	var currentData User
	tx.Debug().Where("id = ?", u.ID).Take(&currentData)
	if u.FullName == "" {
		u.FullName = currentData.FullName
	}

	_, errUpdate := govalidator.ValidateStruct(u)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
