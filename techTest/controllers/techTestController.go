package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"techTest/entities"
	"techTest/services"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TechTestController struct {
	DB              *gorm.DB
	TechTestService services.TechTestService
}

type orderItemRequest struct {
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

type userRequest struct {
	Fullname   string `json:"full_name"`
	FirstOrder bool   `json:"first_order"`
}

type orderHistoryRequest struct {
	UserID       uint64 `json:"user_id"`
	OrderItemID  uint64 `json:"order_item_id"`
	Descriptions string `json:"descriptions"`
}

type json map[string]interface{}

// Order Item
func (tc *TechTestController) CreateOrderItem(ctx echo.Context) error {
	var orderItemRequest orderItemRequest
	if err := ctx.Bind(&orderItemRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	orderItem := entities.OrderItem{
		Name:  orderItemRequest.Name,
		Price: orderItemRequest.Price,
	}

	res, err := tc.TechTestService.InsertOrderItem(orderItem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) GetAllOrderItem(ctx echo.Context) error {
	rawPage := ctx.Param("page")
	page, err := strconv.Atoi(rawPage)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("page %v not an integer", rawPage),
		})

	}
	rawLimit := ctx.Param("limit")
	limit, err := strconv.Atoi(rawLimit)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("limit %v not an integer", rawLimit),
		})
	}

	res, err := tc.TechTestService.GetAllOrderItem(page, limit)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) GetOrderItemById(ctx echo.Context) error {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	res, err := tc.TechTestService.GetOrderItemById(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) UpdateOrderItem(ctx echo.Context) error {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})
	}

	var orderItemRequest orderItemRequest
	if err := ctx.Bind(&orderItemRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	orderItem := entities.OrderItem{
		Name:  orderItemRequest.Name,
		Price: orderItemRequest.Price,
	}

	res, err := tc.TechTestService.UpdateOrderItem(id, orderItem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) DeleteOrderItem(ctx echo.Context) error {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	err = tc.TechTestService.DeleteOrderItem(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
	})
}

// User
func (tc *TechTestController) CreateUser(ctx echo.Context) error {
	var userRequest userRequest
	if err := ctx.Bind(&userRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	var firstOrderTime *time.Time = nil

	if userRequest.FirstOrder {
		now := time.Now()
		firstOrderTime = &now
		fmt.Println(*firstOrderTime)
	}

	user := entities.User{
		FullName:   userRequest.Fullname,
		FirstOrder: firstOrderTime,
	}

	data, err := tc.TechTestService.InsertUser(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
		"data":   data,
	})
}

func (tc *TechTestController) GetAllUser(ctx echo.Context) error {
	rawPage := ctx.Param("page")
	page, err := strconv.Atoi(rawPage)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("page %v not an integer", rawPage),
		})

	}
	rawLimit := ctx.Param("limit")
	limit, err := strconv.Atoi(rawLimit)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("limit %v not an integer", rawLimit),
		})
	}

	res, err := tc.TechTestService.GetAllUser(page, limit)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) GetUserById(ctx echo.Context) error {
	rawId := ctx.Param("id")
	userId, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	res, err := tc.TechTestService.GetUserById(userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) UpdateUser(ctx echo.Context) error {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	var userRequest userRequest
	if err := ctx.Bind(&userRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	var firstOrderTime *time.Time = nil

	if userRequest.FirstOrder {
		now := time.Now()
		firstOrderTime = &now
		fmt.Println(*firstOrderTime)
	}

	user := entities.User{
		FullName:   userRequest.Fullname,
		FirstOrder: firstOrderTime,
	}

	res, err := tc.TechTestService.UpdateUser(id, user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) DeleteUser(ctx echo.Context) error {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	err = tc.TechTestService.DeleteUser(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
	})
}

// Order History
func (tc *TechTestController) CreateOrderHistory(ctx echo.Context) error {
	var orderHistoryRequest orderHistoryRequest
	if err := ctx.Bind(&orderHistoryRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	orderHistory := entities.OrderHistory{
		UserID:       orderHistoryRequest.UserID,
		OrderItemID:  orderHistoryRequest.OrderItemID,
		Descriptions: orderHistoryRequest.Descriptions,
	}

	data, err := tc.TechTestService.InsertOrderHistory(orderHistory)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
		"data":   data,
	})
}

func (tc *TechTestController) GetAllOrderHistory(ctx echo.Context) error {
	rawUserId := ctx.Param("userId")
	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("userId %v not an integer", rawUserId),
		})

	}

	rawPage := ctx.Param("page")
	page, err := strconv.Atoi(rawPage)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("page %v not an integer", rawPage),
		})

	}
	rawLimit := ctx.Param("limit")
	limit, err := strconv.Atoi(rawLimit)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("limit %v not an integer", rawLimit),
		})
	}

	res, err := tc.TechTestService.GetAllOrderHistory(userId, page, limit)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) GetOrderHistoryById(ctx echo.Context) error {
	rawId := ctx.Param("id")
	userId, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	res, err := tc.TechTestService.GetOrderHistoryById(userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) UpdateOrderHistory(ctx echo.Context) error {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	var orderHistoryRequest orderHistoryRequest
	if err := ctx.Bind(&orderHistoryRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	orderHistory := entities.OrderHistory{
		UserID:       orderHistoryRequest.UserID,
		OrderItemID:  orderHistoryRequest.OrderItemID,
		Descriptions: orderHistoryRequest.Descriptions,
	}

	res, err := tc.TechTestService.UpdateOrderHistory(id, orderHistory)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
		"data":   res,
	})
}

func (tc *TechTestController) DeleteOrderHistory(ctx echo.Context) error {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": fmt.Sprintf("id %v not an integer", rawId),
		})

	}

	err = tc.TechTestService.DeleteOrderHistory(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, json{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, json{
		"status": "success",
	})
}
