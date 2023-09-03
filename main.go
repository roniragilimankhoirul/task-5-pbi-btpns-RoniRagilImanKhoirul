package main

import (
	"github.com/gin-gonic/gin"
	database "task-5-pbi-btpns-RoniRagilImanKhoirul/database"
	controllers "task-5-pbi-btpns-RoniRagilImanKhoirul/controllers"
	middleware "task-5-pbi-btpns-RoniRagilImanKhoirul/middleware"
)

func main() {
	var r = gin.Default()
	database.ConnectDb()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Rakamin!",
		})
	})

	r.POST("/users/register", controllers.Register_user)
	r.POST("/users/login", controllers.Login_user)
	r.PUT("/users/:userId", middleware.Require_Auth, controllers.Update_user)
	r.DELETE("/users/:userId", middleware.Require_Auth, controllers.Delete_user)



	r.Run()
}
