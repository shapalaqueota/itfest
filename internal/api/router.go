package api

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router.POST("/signup", Signup)
	router.POST("/login", Login)

	//router.POST("/send-confirmation-email", SendConfirmationEmailHandler)
	//router.GET("/search", SearchFilmsHandler)
	//router.POST("/create-invoice", CreateInvoice)
	//authorized := router.Group("/")
	//authorized.Use(middleware.TokenAuthMiddleware())
	//{
	//	authorized.GET("/users/:id", GetUser)
	//	authorized.POST("/upload/file", UploadFile)
	//	authorized.POST("/upload/episode", UploadEpisode)
	//	authorized.GET("/confirm_email", ConfirmEmailHandler)
	//	authorized.GET("/generate_token/:userID/:email", GenerateTokenHandler)
	//}
}
