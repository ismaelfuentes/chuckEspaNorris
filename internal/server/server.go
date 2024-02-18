package server

import (
	"chuckEspaNorris/internal/controllers"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	server := gin.Default()
	controllers.InitRoutes(server)
  server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}