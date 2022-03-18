package controllers

import (
	"dog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DogSearchInput struct {
	Name string `json:"name" binding:"required"`
}

func FindDogByName(c *gin.Context) {
	dogName := c.Params.ByName("name")
	dog, err := models.GetDogByName(dogName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "succes", "data": dog})
}

func RegisterDog(c *gin.Context) {
	var dog models.Dog
	if err := c.ShouldBindJSON(&dog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := dog.SaveDog()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
