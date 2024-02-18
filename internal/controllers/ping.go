package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}