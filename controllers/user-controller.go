package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"task-5-pbi-btpns-RoniRagilImanKhoirul/app"
	database "task-5-pbi-btpns-RoniRagilImanKhoirul/database"
	helpers "task-5-pbi-btpns-RoniRagilImanKhoirul/helpers"
	models "task-5-pbi-btpns-RoniRagilImanKhoirul/models"
)

type ResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Register_user(c *gin.Context) {
	var user_reg app.AuthRegister
	user_reg.Id = uuid.New().String()
	if err := c.ShouldBindJSON(&user_reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data yang diterima
	if err := helpers.Validation(c, user_reg); err != nil {
		return
	}

	// Enkripsi kata sandi
	hashedPass := helpers.Encrypt_password(c, user_reg.Password)
	if hashedPass == "" {
		return
	}

	// Membuat objek pengguna baru
	user := models.User{
		Id:        user_reg.Id,
		Username:  user_reg.Username,
		Email:     user_reg.Email,
		Password:  hashedPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Membuat pengguna baru di database
	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user. please change the email or the password",
		})
		return
	}

	// Menyiapkan respons dengan "message" dan "data" fields
	response := ResponseData{
		Message: "Successfully registered",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
		},
	}

	// Mengembalikan respons yang telah disiapkan
	c.JSON(http.StatusOK, response)
}

func Login_user(c *gin.Context) {
	var user_login app.AuthLogin
	var user models.User

	// Mengaitkan data login dengan konteks
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data yang diterima
	if err := helpers.Validation(c, user_login); err != nil {
		return
	}

	// Mencari pengguna berdasarkan alamat email
	database.DB.First(&user, "email = ?", user_login.Email)

	// Memeriksa alamat email dan kata sandi yang tidak valid
	if user.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Memeriksa kata sandi
	if err := helpers.Check_password(user.Password, user_login.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Membuat token
	tokenStr, err := helpers.Initialize_token(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Menyiapkan respons dengan token dan data pengguna
	response := ResponseData{
		Message: "Successfully login",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
			"token":    tokenStr,
		},
	}

	// Set token sebagai cookie
	expTime := 86400 // detik (1 hari)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, expTime, "", "", false, true)

	// Mengembalikan respons yang telah disiapkan
	c.JSON(http.StatusOK, response)
}

func Update_user(c *gin.Context) {
	type Input_user struct {
		Id        string    `valid:"required" json:"id"`
		Username  string    `json:"username"`
		Email     string    `valid:"email" json:"email"`
		Password  string    `valid:"minstringlength(6)" json:"password"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	var input_user Input_user
	var user models.User
	input_user.Id = c.Param("userId")
	input_user.UpdatedAt = time.Now()

	if err := c.ShouldBindJSON(&input_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data yang diterima
	if err := helpers.Validation(c, input_user); err != nil {
		return
	}

	// Enkripsi kata sandi jika ada
	if input_user.Password != "" {
		hashedPass := helpers.Encrypt_password(c, input_user.Password)
		if hashedPass == "" {
			return
		}
		input_user.Password = hashedPass
	}

	// Memperbarui data pengguna di database
	if database.DB.Model(&user).Where("id = ?", input_user.Id).Updates(&input_user).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update. Please change the email or the password.",
			"data":    nil, // Anda dapat mengatur data ke nil jika pembaruan gagal.
		})
		return
	}

	// Menyiapkan respons dengan "message" dan "data" fields
	response := ResponseData{
		Message: "Successfully updated",
		Data: gin.H{
			"id":       input_user.Id,
			"username": input_user.Username,
			"email":    input_user.Email,
		},
	}

	// Mengembalikan respons yang telah disiapkan
	c.JSON(http.StatusOK, response)
}

func Delete_user(c *gin.Context) {
	var user models.User
	var photo models.Photo
	var deletedPhotos bool
	userid := c.Param("userId")

	// Menghapus foto-foto yang terkait dengan pengguna
	if err := database.DB.Where("userid = ?", userid).First(&photo).Error; err == nil {
		database.DB.Delete(&models.Photo{}, "userid = ?", userid)
		deletedPhotos = !deletedPhotos
	}

	// Menghapus pengguna berdasarkan ID
	if err := database.DB.Where("id = ?", userid).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	database.DB.Where("id = ?", userid).Delete(&user)

	// Menyiapkan respons
	if deletedPhotos {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully deleted photos and account",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}
