package main

import (
	"fmt"
	"log"
	"os"

	"githib.com/Waris-Shaik/go-crud/controllers"
	"githib.com/Waris-Shaik/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
		log.Fatal("Error PORT is not defined in .env FILE")
	}

	router := gin.Default()

	// routes
	router.GET("/", controllers.HomeController)
	router.GET("/api/v1/posts", controllers.GetAllPosts)
	router.GET("/api/v1/posts/:id", controllers.GetSinglePost)
	router.POST("/api/v1/posts/new", controllers.CreatePost)
	router.PUT("/api/v1/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/v1/posts/:id", controllers.DeletePost)

	// server
	fmt.Println("Server is listening on PORT:", PORT)

	// server error
	if err := router.Run(); err != nil {
		log.Fatal("Error connecting server", err.Error())
	}

}
