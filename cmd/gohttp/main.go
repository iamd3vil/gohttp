package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"gohttp/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")

	headers := make(map[string]string)

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

	r.GET("/headers", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			headers[k] = v[0]
		}
		c.JSON(http.StatusOK, gin.H{"headers": headers})
	})

	r.GET("/status/:code", func(c *gin.Context) {
		code := c.Param("code")
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid code"})
		}
		c.JSON(codeInt, gin.H{"status_code": codeInt})

	})

	r.GET("/get", func(c *gin.Context) {
		queries := c.Request.URL.Query()
		fmt.Println(queries)
		args := make(map[string]string)
		for k, v := range queries {
			args[k] = v[0]
		}
		for k, v := range c.Request.Header {
			headers[k] = v[0]
		}
		c.JSON(http.StatusOK, gin.H{"args": args, "headers": headers,
			"ip": c.Request.Header["X-Forwarded-For"][0], "url": c.Request.URL.String()})
	})

	r.Run(":" + port)
}
