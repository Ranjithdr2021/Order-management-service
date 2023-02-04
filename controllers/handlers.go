package controllers

func orders(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, orders)
}