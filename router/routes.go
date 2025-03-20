package routes

import (
	"net/http"
	websock "tic-tac-toe/websocket"

	"github.com/gin-gonic/gin"
)

func Steproutes(router *gin.Engine, wsHandler *websock.Handler) {
	//HealthCheckRoutes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "website is working fine"})
	})

	router.GET("/ws", wsHandler.Websockethandler)

	//fallback for invalid routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Route not found"})
	})
}
