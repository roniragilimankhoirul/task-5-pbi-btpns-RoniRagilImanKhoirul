package main

import (
	"github.com/gin-gonic/gin"
	database "task-5-pbi-btpns-RoniRagilImanKhoirul/database"
)

func main() {
	var r = gin.Default()
	database.ConnectDb()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Rakamin!",
		})
	})

	r.Run()
}
