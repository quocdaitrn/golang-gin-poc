package controller

import (
	"github/quocdaitrn/golang-gin-poc/entity"
	"github/quocdaitrn/golang-gin-poc/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []*entity.Video
	Save(*gin.Context) (*entity.Video, error)
}

type videoController struct {
	svc service.VideoService
}

func New(s service.VideoService) VideoController {
	return &videoController{
		svc: s,
	}
}

func (c *videoController) FindAll() []*entity.Video {
	return c.svc.FindAll()
}

func (c *videoController) Save(ctx *gin.Context) (*entity.Video, error) {
	var s entity.Video
	err := ctx.ShouldBindJSON(&s)
	if err != nil {
		return nil, err
	}
	c.svc.Save(&s)
	return &s, nil
}