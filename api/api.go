package api

import (
	"net/http"

	"main/websocket"

	"github.com/gin-gonic/gin"
)

// GetSystemInfoHandler handles requests to get system information.
func GetSystemInfoHandler(c *gin.Context) {
	// Get system information using your function from websocket.go
	systemInfo := websocket.GetSystemInfo()

	c.JSON(http.StatusOK, systemInfo)
}
