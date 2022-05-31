package main

import (
	"final-project/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	_, err := database.ConnectDB("mysql")
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	route := gin.Default()
	// User
	// route.POST("/users/register", Register)
	// route.POST("/users/login", Login)
	// route.PUT("/users", UpdateUser)
	// route.DELETE("/users", DeleteUser)

	// Photo
	// route.POST("/photos", CreatePhoto)
	// route.GET("/photos", GetPhotos)
	// route.PUT("/photos/:photoId", UpdatePhoto)
	// route.DELETE("/photos/:photoId", DeletePhoto)

	// Comment
	// route.POST("/comments", CreateComment)
	// route.GET("/comments", GetComments)
	// route.PUT("/comments/:commentId", UpdateComment)
	// route.DELETE("/comments/:commentId", DeleteComment)

	// Social Media
	// route.POST("/socialmedias", CretaeSocialMedia)
	// route.GET("/socialmedias", GetSocialMedias)
	// route.PUT("/socialmedias/:socialMediaId", UpdateSocialMedia)
	// route.DELETE("/socialmedias/:socialMediaId", DeleteSocialMedia)

	route.Run(database.APP_PORT)
}
