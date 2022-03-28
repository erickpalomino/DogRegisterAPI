package controllers

import (
	"dog-app/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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

func FindDogByDNI(c *gin.Context) {
	dogDNI := c.Params.ByName("dni")
	dniNum, br := strconv.ParseUint(dogDNI, 10, 64)
	if br != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("parse to int err: #{br.Error()}"))
		return
	}
	dog, err := models.GetDogByDni(dniNum)
	dog.Diagnostics = dog.GetDiagnosticsFromDog()
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
	c.JSON(http.StatusOK, gin.H{"message": "dog registration success"})
}

func FileUpload(c *gin.Context) {
	dni := c.Params.ByName("dni")
	dniNum, br := strconv.ParseUint(dni, 10, 64)
	if br != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("parse to int err: %s", br.Error()))
	}
	dog, err := models.GetDogByDni(dniNum)
	fmt.Print(dog)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	dir, erroros := os.Getwd()
	if erroros != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	filepath := "./public/" + filename
	url, _ := CloudinaryUpload(filepath, dni)
	dog.Pic = url
	fmt.Print(dog)
	models.GetDB().Save(&dog)
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
