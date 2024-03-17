package controllers

import (
	"net/http"

	"githib.com/Waris-Shaik/go-crud/initializers"
	"githib.com/Waris-Shaik/go-crud/models"
	"github.com/gin-gonic/gin"
)

// homeController
func HomeController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Welcome to golang server",
	})
}

func CreatePost(ctx *gin.Context) {
	// GET the data off req.body
	// Parse the request body to get post data
	var post models.Post
	if err := ctx.Bind(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": "false",
			"error":   err.Error(),
		})
		return
	}

	// Check if required field's are empty
	if post.Title == "" || post.Description == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Please fill all required field's",
		})
		return
	}

	// Store the data in the database
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create post",
		})
		return
	}

	// Return the created post
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post successfully created",
		"post":    post,
	})
}

func GetAllPosts(ctx *gin.Context) {
	var posts []models.Post

	// Retrieve all the posts from the database
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to retrieve posts",
		})
		return
	}

	// Return the retrieved posts
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"posts":   posts,
	})

}

func GetSinglePost(ctx *gin.Context) {

	// Get postID from URL parameter
	postID := ctx.Param("id")

	// Check if postID is valid
	if postID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Post ID is required",
		})
		return
	}

	// Retrieve post from the database
	var post models.Post
	result := initializers.DB.First(&post, postID)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	// Return the retreived post
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"post":    post,
	})

}

func UpdatePost(ctx *gin.Context) {
	// Get post ID from URL parameter
	postID := ctx.Param("id")

	// Check if post ID is valid
	if postID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Post ID is required",
		})
		return
	}

	var post models.Post
	// Retreive the existing post from the database
	result := initializers.DB.First(&post, postID)
	// Return if post is not found
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	// Get the data off req.body
	//	Parse the req,body to get data
	if err := ctx.Bind(&post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Update post
	initializers.DB.Model(&post).Updates(models.Post{Title: post.Title, Description: post.Description})

	// Return the updated post
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post successfully updated",
		"post":    post,
	})

}

func DeletePost(ctx *gin.Context) {
	// Get post ID from URL parameter
	postID := ctx.Param("id")

	//	Check is postID is valid
	if postID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Post ID is required",
		})
		return
	}

	// Retreive the existing post from the database
	var post models.Post
	result := initializers.DB.First(&post, postID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	// Delete the post
	initializers.DB.Delete(&post)

	// Return the respond
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post successfully deleted",
		"post":    post,
	})

}
