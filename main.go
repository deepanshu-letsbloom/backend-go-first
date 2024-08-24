package main

import (
	"backend/app/router"
	"backend/config"
	"os"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()

	// gin.SetMode(gin.ReleaseMode)
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
