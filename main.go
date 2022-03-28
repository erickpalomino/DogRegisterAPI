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
	public.GET("/getUser", controllers.CurrentUser)

	protected := r.Group("/api/worker")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/getDoctors", controllers.GetDoctors)
	protected.GET("/user", controllers.CurrentUser)
	protected.POST("/dog/register", controllers.RegisterDog)
	protected.POST("/dog/:dni/upload", controllers.FileUpload)
	protected.GET("/dog/:name/getByName", controllers.FindDogByName)
	protected.GET("/dog/getDog/:dni", controllers.FindDogByDNI)
	protected.POST("/dog/diagnostic/newDiagnostic", controllers.RegisterDiagnostic)
	protected.POST("/dog/diagnostic/:id/uploadFiles", controllers.UploadXrayBloodResult)
	public.POST("/dog/date/register", controllers.RegisterDate)
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
