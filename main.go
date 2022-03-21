package main

import (
	"dog-app/controllers"
	"dog-app/middlewares"
	"dog-app/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.InitDB()
	r := gin.Default()
	models.GetDB()

	r.Use(CORSMiddleware())
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/worker")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protected.POST("/dog/register", controllers.RegisterDog)
	protected.POST("/dog/:dni/upload", controllers.FileUpload)
	protected.GET("/dog/:name/getByName", controllers.FindDogByName)

	r.Run(":8080")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
