package routers

import (
	"mygram/controllers"
	"mygram/middlewere"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewere.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewere.Authentication(), controllers.DeleteUser)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewere.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:photoId", middlewere.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewere.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewere.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComments)
		commentRouter.PUT("/:commentId", middlewere.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewere.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewere.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewere.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewere.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}
	return r
}
