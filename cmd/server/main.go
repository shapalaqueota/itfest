package main

import (
	"github.com/gin-gonic/gin"
	"itfest/internal/api"
	"itfest/internal/db"
	"itfest/internal/middleware"
	"itfest/internal/utils"
)

func main() {

	db.ConnectDB()
	defer db.DB.Close()

	utils.InitVKCloudService()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	api.Router(router)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}

}
