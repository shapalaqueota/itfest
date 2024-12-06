package api

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router.POST("/signup", Signup)
	router.POST("/signin", Signin)

	authorized := router.Group("/")
	{
		authorized.GET("/users/:id", GetUser)
		router.POST("/upload", UploadImageHandler)
		router.POST("/items", CreateItemHandler)

	}
}
