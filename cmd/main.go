package main

import (
	"net/http"
	"os"

	"gohttp/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/ip", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ip": c.Request.RemoteAddr})
	})
	r.GET("/user-agent", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user-agent": c.Request.UserAgent})
	})
	r.Run(":" + port)
}
