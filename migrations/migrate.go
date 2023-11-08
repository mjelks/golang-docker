package main

import (
	"github.com/gorm-starter/initializers"
	"github.com/gorm-starter/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
