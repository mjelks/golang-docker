package main

import (
	"net/http"
	"os"

	"github.com/gorm-starter/controllers"
	"github.com/gorm-starter/initializers"
	"github.com/gorm-starter/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.User{})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	r.GET("/ping", controllers.Ping)
	r.GET("/htmx", controllers.HtmxIndex)
	r.POST("/htmx", controllers.HtmxCreate)

	r.Run(":" + os.Getenv("API_SERVER_PORT"))
}
