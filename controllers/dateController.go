package controllers

import (
	"dog-app/models"
	"dog-app/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterDate(c *gin.Context) {
	var date models.Date
	if err := c.ShouldBindJSON(&date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := models.GetUserByID(user_id)
	newDate, err := date.SaveDate()
	newDate.Owner = u.Username
	models.GetDB().Save(&newDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "date registration success", "data": newDate})
}
