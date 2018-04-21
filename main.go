package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h00s/todo-api/config"
	"github.com/h00s/todo-api/db"
	"github.com/h00s/todo-api/list"
	"github.com/h00s/todo-api/logger"
)

func main() {
	c, err := config.Load("configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.New(c.Log.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	db, err := db.Connect(c.Database)
	if err != nil {
		l.Fatal(err.Error())
	}

	lc := list.NewController(db, l)

	if c.Router.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/lists", lc.GetLists)
	}

	l.Info("Starting server..")
	r.Run(c.Server.Address)
}
