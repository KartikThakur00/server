package main

import (
	"react-gin/server/config"
	"react-gin/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router:= gin.Default()
	config.Connect()
	routes.Route(router)
	router.Run(":8080")
}
