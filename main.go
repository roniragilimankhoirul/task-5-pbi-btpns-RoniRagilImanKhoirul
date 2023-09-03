package main

import (
	"github.com/gin-gonic/gin"
	database "task-5-pbi-btpns-RoniRagilImanKhoirul/database"
	controllers "task-5-pbi-btpns-RoniRagilImanKhoirul/controllers"
	middleware "task-5-pbi-btpns-RoniRagilImanKhoirul/middleware"
)

func main() {
	// Inisialisasi router Gin
	var r = gin.Default()

	// Menghubungkan ke database
	database.ConnectDb()

	// Route untuk endpoint beranda
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halo, Rakamin!",
		})
	})

	// Route untuk endpoint registrasi pengguna
	r.POST("/users/register", controllers.Register_user)

	// Route untuk endpoint login pengguna
	r.POST("/users/login", controllers.Login_user)

	// Route untuk endpoint pembaruan data pengguna
	r.PUT("/users/:userId", middleware.Require_Auth, controllers.Update_user)

	// Route untuk endpoint penghapusan data pengguna
	r.DELETE("/users/:userId", middleware.Require_Auth, controllers.Delete_user)



	r.POST("/photos", middleware.Require_Auth, controllers.Create_photo)
	r.GET("/photos", middleware.Require_Auth, controllers.Show_photo)
	r.PUT("/photos/:photoId", middleware.Require_Auth, controllers.Update_photo)
	r.DELETE("/photos/:photoId", middleware.Require_Auth, controllers.Delete_photo)


	// Menjalankan server web
	r.Run()
}
