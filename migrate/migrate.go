package main

import (
	"githib.com/Waris-Shaik/go-crud/initializers"
	"githib.com/Waris-Shaik/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
