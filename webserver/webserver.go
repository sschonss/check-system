package webserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartWebServer starts the web server.
func StartWebServer() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		renderHTML(c)
	})

	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()
		c.Next()
	})

	r.Run(":8081")
}

func renderHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "static/index.html", nil)
}
