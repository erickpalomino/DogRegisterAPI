package controllers

import (
	"dog-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterDiagnostic(c *gin.Context) {
	var diag models.Diagnostic
	if err := c.ShouldBindJSON(&diag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := diag.SaveDiagnostic()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "diagnostic registration success"})
}
