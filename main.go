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

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/worker")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protected.POST("/dog/register", controllers.RegisterDog)
	protected.GET("/dog/:name/getByName", controllers.FindDogByName)

	r.Run(":8080")
}
