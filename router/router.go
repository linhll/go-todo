package router

import (
	"github.com/gin-gonic/gin"
)

// Router return gin router
func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "It works!")
	})
	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.String(200, "/api")
	})
	return r
}
