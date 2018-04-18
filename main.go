package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.String(200, "Hello world")
		})
	}

	r.Run()
}
