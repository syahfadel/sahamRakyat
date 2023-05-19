package services

import (
	"errors"
	"fmt"
	"techTest/entities"

	"gorm.io/gorm"
)

type TechTestService struct {
	DB *gorm.DB
}

func (ts *TechTestService) InsertOrderItem(orderItem entities.OrderItem) (entities.OrderItem, error) {
	if err := ts.DB.Debug().Create(&orderItem).Error; err != nil {
		return entities.OrderItem{}, err
	}
	return orderItem, nil
}

func (ts *TechTestService) GetAllOrderItem(page, limit int) ([]entities.OrderItem, error) {
	var result []entities.OrderItem

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := ts.DB.Debug().Offset(offset).Limit(limit).Order("id asc").Find(&result).Error; err != nil {
		return []entities.OrderItem{}, err
	}
	return result, nil
}

func (ts *TechTestService) GetOrderItemById(id int) (entities.OrderItem, error) {
	var result entities.OrderItem
	if err := ts.DB.Debug().Where("id = ?", id).Take(&result).Error; err != nil {
		return entities.OrderItem{}, err
	}
	return result, nil
}

func (ts *TechTestService) UpdateOrderItem(id int, updateOrderItem entities.OrderItem) (entities.OrderItem, error) {
	var currentData entities.OrderItem

	if err := ts.DB.Debug().Where("id = ?", id).Take(&currentData).Error; err != nil {
		return entities.OrderItem{}, err
	}

	updateOrderItem.ID = uint64(id)
	updateOrderItem.CreatedAt = currentData.CreatedAt

	if updateOrderItem.Name == "" {
		updateOrderItem.Name = currentData.Name
	}

	if updateOrderItem.Price == 0 {
		updateOrderItem.Price = currentData.Price
	}

	res := ts.DB.Debug().Model(&updateOrderItem).Where("id = ?", id).Updates(&updateOrderItem)
	if res.Error != nil {
		return entities.OrderItem{}, res.Error
	}

	if res.RowsAffected == 0 {
		return entities.OrderItem{}, errors.New("no data updated")
	}

	return updateOrderItem, nil
}

func (ts *TechTestService) DeleteOrderItem(id int) error {
	res := ts.DB.Debug().Where("id = ?", id).Delete(&entities.OrderItem{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("order item with id %v not available", id))
	}
	return nil
}

func (ts *TechTestService) InsertUser(user entities.User) (entities.User, error) {
	if err := ts.DB.Debug().Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ts *TechTestService) GetAllUser(page, limit int) ([]entities.User, error) {
	var result []entities.User

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := ts.DB.Debug().Offset(offset).Limit(limit).Order("id asc").Find(&result).Error; err != nil {
		return []entities.User{}, err
	}
	return result, nil
}

func (ts *TechTestService) GetUserById(id int) (entities.User, error) {
	var result entities.User
	if err := ts.DB.Debug().Where("id = ?", id).Take(&result).Error; err != nil {
		return entities.User{}, err
	}
	return result, nil
}

func (ts *TechTestService) UpdateUser(id int, updateUser entities.User) (entities.User, error) {
	var currentData entities.User

	if err := ts.DB.Debug().Where("id = ?", id).Take(&currentData).Error; err != nil {
		return entities.User{}, err
	}

	updateUser.ID = uint64(id)
	updateUser.CreatedAt = currentData.CreatedAt

	if updateUser.FullName == "" {
		updateUser.FullName = currentData.FullName
	}

	if updateUser.FirstOrder == nil {
		updateUser.FirstOrder = currentData.FirstOrder
	}

	res := ts.DB.Debug().Model(&updateUser).Where("id = ?", id).Updates(&updateUser)
	if res.Error != nil {
		return entities.User{}, res.Error
	}

	if res.RowsAffected == 0 {
		return entities.User{}, errors.New("no data updated")
	}

	return updateUser, nil
}

func (ts *TechTestService) DeleteUser(id int) error {
	res := ts.DB.Debug().Where("id = ?", id).Delete(&entities.User{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("user with id %v not available", id))
	}
	return nil
}

func (ts *TechTestService) InsertOrderHistory(orderHistory entities.OrderHistory) (entities.OrderHistory, error) {
	if err := ts.DB.Debug().Create(&orderHistory).Error; err != nil {
		return entities.OrderHistory{}, err
	}
	return orderHistory, nil
}

func (ts *TechTestService) GetAllOrderHistory(userId, page, limit int) ([]entities.OrderHistory, error) {
	var result []entities.OrderHistory

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := ts.DB.Debug().Where("user_id = ?", userId).Offset(offset).Limit(limit).Order("id asc").Find(&result).Error; err != nil {
		return []entities.OrderHistory{}, err
	}
	return result, nil
}

func (ts *TechTestService) GetOrderHistoryById(id int) (entities.OrderHistory, error) {
	var result entities.OrderHistory
	if err := ts.DB.Debug().Where("id = ?", id).Take(&result).Error; err != nil {
		return entities.OrderHistory{}, err
	}
	return result, nil
}

func (ts *TechTestService) UpdateOrderHistory(id int, updateOrderHistory entities.OrderHistory) (entities.OrderHistory, error) {
	var currentData entities.OrderHistory

	if err := ts.DB.Debug().Where("id = ?", id).Take(&currentData).Error; err != nil {
		return entities.OrderHistory{}, err
	}

	updateOrderHistory.ID = uint64(id)
	updateOrderHistory.CreatedAt = currentData.CreatedAt

	if updateOrderHistory.UserID == 0 {
		updateOrderHistory.UserID = currentData.UserID
	}

	if updateOrderHistory.UserID == 0 {
		updateOrderHistory.OrderItemID = currentData.OrderItemID
	}

	if updateOrderHistory.Descriptions == "" {
		updateOrderHistory.Descriptions = currentData.Descriptions
	}

	res := ts.DB.Debug().Model(&updateOrderHistory).Where("id = ?", id).Updates(&updateOrderHistory)
	if res.Error != nil {
		return entities.OrderHistory{}, res.Error
	}

	if res.RowsAffected == 0 {
		return entities.OrderHistory{}, errors.New("no data updated")
	}

	return updateOrderHistory, nil
}

func (ts *TechTestService) DeleteOrderHistory(id int) error {
	res := ts.DB.Debug().Where("id = ?", id).Delete(&entities.OrderHistory{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("order history with id %v not available", id))
	}
	return nil
}
