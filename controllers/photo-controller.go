package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "task-5-pbi-btpns-RoniRagilImanKhoirul/database"
	helpers "task-5-pbi-btpns-RoniRagilImanKhoirul/helpers"
	models "task-5-pbi-btpns-RoniRagilImanKhoirul/models"
)

func Create_photo(c *gin.Context) {
	var photo models.Photo
	photo.Id = uuid.New().String()
	userid, _ := c.Get("userid")
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messsage": err.Error()})
		return
	}

	if database.DB.Create(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to add item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": photo})

}


func Show_photo(c *gin.Context) {
	var photos []models.Photo
	userid, _ := c.Get("userid")

	database.DB.Where("userid = ?", userid).Find(&photos)
	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func Update_photo(c *gin.Context) {
	var photo models.Photo
	photo.Id = c.Param("photoId")
	userid, _ := c.Get("userid")
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Perform the update operation
	result := database.DB.Model(&photo).Where("userid = ?", photo.Userid).Updates(&photo)

	// Check if the update was successful
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not found"})
		return
	}

	// Fetch the updated photo data from the database
	database.DB.First(&photo, "id = ?", photo.Id)

	c.JSON(http.StatusOK, gin.H{"message": "data has been updated", "data": photo})
}
