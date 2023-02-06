package main

import (
	"Order-management/controllers"
	"Order-management/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	database.NewTable()

	router := gin.Default()

	router.POST("/add-order", controllers.AddOrders)

	router.PUT("/update-status/:O_id", controllers.UpdateStatus)

	router.GET("/GetOrdersonstatus/:status", controllers.GetOrdersonstatus)

	router.GET("/GetOrdersById/:O_id", controllers.GetOrdersById)

	router.GET("/orders", controllers.Orders)

	router.Run("localhost:8080")
}
