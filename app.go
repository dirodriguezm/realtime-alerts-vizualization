package main

import (
	"github.com/gin-gonic/gin"
)

func start(kafkaBrokers, topic string, repo ObjectRepository, username, password *string) {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.GET("/", homeHandler)
	r.Static("/static", "./static")
	r.GET("/ws", wsHandler(kafkaBrokers, topic, repo, username, password))
	r.Run() // listen and serve on 0.0.0.0:8080
}
