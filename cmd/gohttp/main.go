package main

import (
	"log"
	"net/http"
	"os"

	"gohttp/cmd/gohttp/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/ip", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ip": c.Request.Header["X-Forwarded-For"][0]})
	})
	r.GET("/user-agent", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user-agent": c.Request.UserAgent()})
	})
	r.Run(":" + port)
}
