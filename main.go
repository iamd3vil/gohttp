package main

import (
	"net/http"

	"gohttp/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/ip", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ip": c.Request.RemoteAddr})
	})
	r.GET("/user-agent", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user-agent": c.Request.UserAgent})
	})
	r.Run("0.0.0.0:80")
}
