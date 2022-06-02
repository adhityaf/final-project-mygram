package main

import (
	"final-project/controllers"
	"final-project/database"
	"final-project/middleware"
	"final-project/repositories"
	"final-project/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	db, err := database.ConnectDB("mysql")
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	socialMediaRepo := repositories.NewSocialMediaRepo(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	route := gin.Default()
	// User
	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userController.Register)
		userRoute.POST("/login", userController.Login)

		userRoute.Use(middleware.Auth())
		userRoute.PATCH("/:userId", userController.UpdateUser)
		userRoute.DELETE("/:userId", userController.DeleteUser)
	}

	// Photo
	// route.POST("/photos", CreatePhoto)
	// route.GET("/photos", GetPhotos)
	// route.PATCH("/photos/:photoId", UpdatePhoto)
	// route.DELETE("/photos/:photoId", DeletePhoto)

	// Comment
	// route.POST("/comments", CreateComment)
	// route.GET("/comments", GetComments)
	// route.PATCH("/comments/:commentId", UpdateComment)
	// route.DELETE("/comments/:commentId", DeleteComment)

	// Social Media
	socialMediaRoute := route.Group("/socialmedias")
	{
		socialMediaRoute.Use(middleware.Auth())
		socialMediaRoute.POST("", socialMediaController.CreateSocialMedia)
		socialMediaRoute.GET("", socialMediaController.GetSocialMedias)
		socialMediaRoute.PATCH("/:socialMediaId", socialMediaController.UpdateSocialMedia)
		socialMediaRoute.DELETE("/:socialMediaId", socialMediaController.DeleteSocialMedia)
	}

	route.Run(database.APP_PORT)
}
