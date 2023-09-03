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


func Login_user(c *gin.Context) {
	var user_login app.AuthLogin
	var user models.User

	// Bind user_login to context
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if err := helpers.Validation(c, user_login); err != nil {
		return
	}

	database.DB.First(&user, "email = ?", user_login.Email)

	// Check for invalid email and password
	if user.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Check password
	if err := helpers.Check_password(user.Password, user_login.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a token
	tokenStr, err := helpers.Initialize_token(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Set token and user data in the response
	response := ResponseData{
		Message: "Successfully login",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
			"token": tokenStr,
		},
	}

	// Set token as a cookie
	expTime := 86400 // seconds
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, expTime, "", "", false, true)

	// Return the structured response
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
	if err := helpers.Validation(c, input_user); err != nil {
		return
	}
	if input_user.Password != "" {
		hashedPass := helpers.Encrypt_password(c, input_user.Password)
		if hashedPass == "" {
			return
		}
		input_user.Password = hashedPass
	}
	if database.DB.Model(&user).Where("id = ?", input_user.Id).Updates(&input_user).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update. Please change the email or the password.",
			"data":    nil, // You can set data to nil if the update fails.
		})
		return
	}
	// Create the response struct with "message" and "data" fields
	response := ResponseData{
		Message: "Successfully updated",
		Data: gin.H{
			"id":       input_user.Id,
			"username": input_user.Username,
			"email":    input_user.Email,
		},
	}

	// Return the structured response
	c.JSON(http.StatusOK, response)
}


