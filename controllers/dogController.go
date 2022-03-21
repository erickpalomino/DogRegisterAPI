package controllers

import (
	"dog-app/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"context"
    "github.com/cloudinary/cloudinary-go"
    "github.com/cloudinary/cloudinary-go/api/uploader"

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

func FileUpload(c *gin.Context) {
	dni := c.Params.ByName("dni")
	dog, err := models.GetDogByDni(dni)
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
 url:=CloudinaryUpload(filepath,dni)
	dog.Pic = url
	fmt.Print(dog)
	models.GetDB().Save(&dog)

	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func CloudinaryUpload(path string,dni string) (url string){
	cld, _ := cloudinary.NewFromParams("dziaapbmr", "922581187159196", "sY44Tzpsnok0L-SSYx3JhtbF73I")
ctx := context.Background()

resp, err := cld.Upload.Upload(ctx, path, uploader.UploadParams{PublicID: dni,
    Transformation: "c_crop,g_center/q_auto/f_auto", Tags: []string{"fruit"}})

		my_image, err := cld.Image(dni)
		if err != nil {
				fmt.Println("error")
		}

		url, errstring := my_image.String()
		fmt.Print(resp)
if errstring != nil {
    fmt.Println("error")
}
fmt.Print(url)
		return url
	}
