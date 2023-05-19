package routers

import (
	"techTest/controllers"
	"techTest/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StartService(db *gorm.DB) *echo.Echo {
	techTestService := services.TechTestService{
		DB: db,
	}

	techTestController := controllers.TechTestController{
		DB:              db,
		TechTestService: techTestService,
	}

	app := echo.New()
	orderItem := app.Group("/orderItem")
	orderItem.POST("", techTestController.CreateOrderItem)
	orderItem.GET("/page/:page/limit/:limit", techTestController.GetAllOrderItem)
	orderItem.GET("/:id", techTestController.GetOrderItemById)
	orderItem.PUT("/:id", techTestController.UpdateOrderItem)
	orderItem.DELETE("/:id", techTestController.DeleteOrderItem)

	user := app.Group("/user")
	user.POST("", techTestController.CreateUser)
	user.GET("/page/:page/limit/:limit", techTestController.GetAllUser)
	user.GET("/:id", techTestController.GetUserById)
	user.PUT("/:id", techTestController.UpdateUser)
	user.DELETE("/:id", techTestController.DeleteUser)

	orderHistory := app.Group("/orderHistory")
	orderHistory.POST("", techTestController.CreateOrderHistory)
	orderHistory.GET("/userId/:userId/page/:page/limit/:limit", techTestController.GetAllOrderHistory)
	orderHistory.GET("/:id", techTestController.GetOrderHistoryById)
	orderHistory.PUT("/:id", techTestController.UpdateOrderHistory)
	orderHistory.DELETE("/:id", techTestController.DeleteOrderHistory)

	return app
}
