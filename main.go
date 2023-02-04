package main

func main(){
	router := gin.Default()
    router.GET("/orders", controllers.orders)

    router.Run("localhost:8080")
}