package main

import (
	"finalproject/configs"
	"finalproject/controllers"
	"finalproject/middlewares"
	"finalproject/repositories"
	"finalproject/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.ConnectDB()

	//users
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	//photo
	photoRepo := repositories.NewPhotoRepo(db)
	photoService := services.NewPhotoService(photoRepo, userRepo)
	photoController := controllers.NewPhotoController(photoService)

	//comment
	commentRepo := repositories.NewCommentRepo(db)
	commentService := services.NewCommentService(commentRepo, userRepo, photoRepo)
	commentController := controllers.NewCommentController(commentService)

	//social media
	socmedRepo := repositories.NewSocialMediaRepo(db)
	socmedService := services.NewSocialMediaService(socmedRepo, userRepo)
	socmedController := controllers.NewSocialMediaController(socmedService)

	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.RegisterUser)
		userRouter.POST("/login", userController.UserLogin)
	}

	userRouterWithAuth := router.Group("/users")
	{
		userRouterWithAuth.Use(middlewares.Auth())
		userRouterWithAuth.PUT("/:userId", userController.UpdateUser)
		userRouterWithAuth.DELETE("", userController.DeleteUser)
	}

	photoRouterWithAuth := router.Group("/photos")
	{
		photoRouterWithAuth.Use(middlewares.Auth())
		photoRouterWithAuth.POST("", photoController.SavePhoto)
		photoRouterWithAuth.GET("", photoController.GetPhotos)
		photoRouterWithAuth.DELETE("/:photoId", photoController.DeletePhoto)
		photoRouterWithAuth.PUT("/:photoId", photoController.UpdatePhoto)
	}

	commentRouterWithAuth := router.Group("/comments")
	{
		commentRouterWithAuth.Use(middlewares.Auth())
		commentRouterWithAuth.POST("", commentController.CreateComment)
		commentRouterWithAuth.DELETE("/:commentId", commentController.DeleteComment)
		commentRouterWithAuth.GET("", commentController.GetComments)
		commentRouterWithAuth.PUT("/:commentId", commentController.UpdateComment)
	}

	socmedRouterWithAuth := router.Group("/socialmedias")
	{
		socmedRouterWithAuth.Use(middlewares.Auth())
		socmedRouterWithAuth.POST("", socmedController.CreateSocialMedia)
		socmedRouterWithAuth.DELETE("/:socialMediaId", socmedController.DeleteSocialMedia)
		socmedRouterWithAuth.GET("", socmedController.GetSocmed)
		socmedRouterWithAuth.PUT("/:socialMediaId", socmedController.UpdateSocmed)
	}

	router.Run(configs.APP_PORT)
}
