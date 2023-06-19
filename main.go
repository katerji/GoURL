package main

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/db"
	"github.com/katerji/UserAuthKit/handler"
	"github.com/katerji/UserAuthKit/middleware"
)

func main() {
	initDB()
	initWebServer()
}

func initDB() {
	client := db.GetDbInstance()
	err := client.Ping()
	if err != nil {
		panic(err)
	}
}

func initWebServer() {
	router := gin.Default()
	api := router.Group("/api")

	api.GET(handler.LandingPath, handler.LandingController)
	api.GET(handler.RouterPath, handler.RouterHandler)

	auth := api.Group("/auth")

	auth.POST(handler.RegisterPath, handler.RegisterHandler)
	auth.POST(handler.LoginPath, handler.LoginHandler)
	auth.POST(handler.RefreshTokenPath, handler.RefreshTokenHandler)

	api.Use(middleware.GetAuthMiddleware())

	api.GET(handler.UserInfoPath, handler.UserInfoHandler)
	api.POST(handler.URLCreatePath, handler.URLCreateHandler)
	api.DELETE(handler.URLDeletePath, handler.URLDeleteHandler)

	err := router.Run(":85")
	if err != nil {
		panic(err)
	}
}
