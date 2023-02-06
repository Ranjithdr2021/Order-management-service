package controllers

import (
	"Order-management/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	orderId = "O_id"
	jsn     = "application/json"
)

var order Order

// Orders gives you all orders in response
func Orders(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", jsn)
	db := database.ConnectToDB()
	defer db.Close()
	row, err := db.Query("SELECT * FROM orders")
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}

	orders := []Order{}
	for row.Next() {
		var order Order
		var items []byte
		if err := row.Scan(&order.Id, &order.Status, &order.Total, &order.CurrencyUnit, &items); err != nil {
			fmt.Fprint(c.Writer, err)
			return
		}
		if err := json.Unmarshal(items, &order.Items); err != nil {
			fmt.Fprint(c.Writer, err)
			return
		}
		orders = append(orders, order)
	}
	c.Writer.Header().Set("Content-Type", jsn)
	c.JSON(http.StatusOK, orders)
}

// GetOrdersonstatus will gives you orders in sorted order based on given status in response
func GetOrdersonstatus(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", jsn)
	db := database.ConnectToDB()
	defer db.Close()
	row, err := db.Query("SELECT * FROM orders")
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}

	orders := []Order{}
	for row.Next() {
		var order Order
		var items []byte
		if err := row.Scan(&order.Id, &order.Status, &order.Total, &order.CurrencyUnit, &items); err != nil {
			fmt.Fprint(c.Writer, err)

		}
		if err := json.Unmarshal(items, &order.Items); err != nil {
			fmt.Fprint(c.Writer, err)
			return
		}
		orders = append(orders, order)
	}
	status := c.Params.ByName("status")
	orderss := getOrderByStatus(orders, status)
	c.Writer.Header().Set("Content-Type", jsn)
	c.JSON(http.StatusOK, orderss)
}

// GetOrdersById will gives you order with requested order id in param in response
func GetOrdersById(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", jsn)
	db := database.ConnectToDB()
	defer db.Close()
	row, err := db.Query("SELECT * FROM orders")
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}

	orders := []Order{}
	for row.Next() {
		var order Order
		var items []byte
		if err := row.Scan(&order.Id, &order.Status, &order.Total, &order.CurrencyUnit, &items); err != nil {
			fmt.Fprint(c.Writer, err)

		}
		if err := json.Unmarshal(items, &order.Items); err != nil {
			fmt.Fprint(c.Writer, err)
			return
		}
		orders = append(orders, order)
	}
	id := c.Params.ByName("O_id")
	orderss := getOrderById(orders, id)
	c.Writer.Header().Set("Content-Type", jsn)
	c.JSON(http.StatusOK, orderss)
}

// AddOrders adds new orders into the system
// No duplicate order ids allowed
func AddOrders(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", jsn)
	items, _ := json.Marshal(order.Items)
	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		fmt.Fprint(c.Writer, "Failure")
		panic(err.Error())
	}
	db := database.ConnectToDB()
	defer db.Close()

	_, er := db.Query("insert into orders(O_id, Status, Total, CurrencyUnit, Items) values(?,?,?,?,?)", order.Id, order.Status, order.Total, order.CurrencyUnit, items)
	if er != nil {
		fmt.Fprint(c.Writer, "orders table -->", er)
		return
	}

	fmt.Fprint(c.Writer, "Success")
}

// UpdateStatus updates the status of the order as given in request body.
func UpdateStatus(c *gin.Context) {
	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	O_id := c.Param(orderId)
	db := database.ConnectToDB()
	defer db.Close()

	_, er := db.Exec("update orders set Status=? where O_id=?", order.Status, O_id)
	if er != nil {
		fmt.Fprint(c.Writer, er)
		return
	}
	fmt.Fprint(c.Writer, "Status updated Successfully...  ")
}
