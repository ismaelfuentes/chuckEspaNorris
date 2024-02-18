package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {

	server.GET("/ping", handlePing)
	server.GET("/jokes", handlefetchJokes)

}