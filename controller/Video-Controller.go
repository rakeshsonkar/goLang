package controller

import (
	"example.com/greetings/entity"
	"example.com/greetings/services"
	"github.com/gin-gonic/gin"
)

type VidioController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VidioController {

	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var Video entity.Video
	ctx.BindJSON(&Video)
	c.service.Save(Video)
	return Video
}
