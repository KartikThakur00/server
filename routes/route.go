package routes

import (
	"react-gin/server/controller"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine){
	router.GET("/",controller.Start)
	router.GET("/api",controller.GetApi)
	router.GET("/api/:file",controller.GetFile)
	router.POST("/api",controller.PostApi)
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(404, gin.H{}) })
}
