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

	// Returns JSON containing origin of the request.
	// Here, as Heroku contains the original IP address in "X-Forwarded-For".
	// So, RemoteAddr is not used.
	r.GET("/ip", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"ip": c.Request.Header.Get("X-Forwarded-For")})
	})

	// Returns JSON containing user agent of the request.
	r.GET("/user-agent", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"user-agent": c.Request.UserAgent()})
	})

	r.GET("/headers", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			headers[k] = v[0]
		}
		c.IndentedJSON(http.StatusOK, gin.H{"headers": headers})
	})

	// Returns user specified status code and also a JSON if possible consisting of the code.
	r.GET("/status/:code", func(c *gin.Context) {
		code := c.Param("code")
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid code"})
		}
		c.IndentedJSON(codeInt, gin.H{"status_code": codeInt})

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
		c.IndentedJSON(http.StatusOK, gin.H{"args": args, "headers": headers,
			"ip": c.Request.Header.Get("X-Forwarded-For"), "url": c.Request.URL.String()})
	})

	// r.GET("/redirect/:n", func(c *gin.Context) {
	// 	redirectsStr := c.Param("n")
	// 	redirects, err := strconv.Atoi(redirectsStr)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid code"})
	// 	}
	// 	fmt.Println(redirects - 1)
	// 	if redirects == 1 {
	// 		c.Redirect(302, "/get")
	// 	} else {
	// 		c.Redirect(302, "/redirect/"+string(redirects-1))
	// 	}
	// })

	r.POST("/post", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		values := make(map[string]string)
		for k, v := range c.Request.PostForm {
			values[k] = v[0]
		}
		c.IndentedJSON(http.StatusCreated, gin.H{"values": values})
	})

	r.Run(":" + port)
}
