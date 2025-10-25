package routes

import (
	"github.com/donny-c-1/sorcemoola/server/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handlers.HomeHandler)
}