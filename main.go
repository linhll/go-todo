package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/api", func(c *gin.Context) {
		c.String(200, "It's work!")
	})

	router.Run()
}
