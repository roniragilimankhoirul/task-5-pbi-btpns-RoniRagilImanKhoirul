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
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

func Register_user(c *gin.Context) {
	var user_reg app.AuthRegister
	user_reg.Id = uuid.New().String()
	if err := c.ShouldBindJSON(&user_reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validation
	if err := helpers.Validation(c, user_reg); err != nil {
		return
	}
	//hash password
	hashedPass := helpers.Encrypt_password(c, user_reg.Password)
	if hashedPass == "" {
		return
	}

	user := models.User{
		Id:        user_reg.Id,
		Username:  user_reg.Username,
		Email:     user_reg.Email,
		Password:  hashedPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user. please change the email or the password",
		})
		return
	}
	// Create the response struct with "message" and "data" fields
	response := ResponseData{
		Message: "Successfully registered",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
		},
	}

	// Return the structured response
	c.JSON(http.StatusOK, response)

}