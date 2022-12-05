package main

import (
	"example.com/greetings/controller"
	"example.com/greetings/services"
	"github.com/gin-gonic/gin"
)

var (
	VideoService    services.VideoService      = services.New()
	VidioController controller.VidioController = controller.New(VideoService)
)

func main() {

	Server := gin.Default()
	Server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VidioController.FindAll())
	})
	Server.POST("/savepost", func(ctx *gin.Context) {
		ctx.JSON(200, VidioController.Save(ctx))
	})
	Server.Run(":9000")
}
