package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/todo-api/logger"
)

func main() {
	l, err := logger.New("todo-api.log")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.String(200, "Hello world")
		})
	}

	l.Info("Starting server..")
	r.Run()
}
