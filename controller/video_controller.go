package controller

import (
	"github/quocdaitrn/golang-gin-poc/entity"
	"github/quocdaitrn/golang-gin-poc/service"
	"github/quocdaitrn/golang-gin-poc/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type VideoController interface {
	FindAll() []*entity.Video
	Save(*gin.Context) (*entity.Video, error)
	ShowAll(ctx *gin.Context) map[string]interface{}
}

type videoController struct {
	svc service.VideoService
}

func New(s service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("isCool", validation.ValidateCoolVideoTitle)
	return &videoController{
		svc: s,
	}
}

func (c *videoController) FindAll() []*entity.Video {
	return c.svc.FindAll()
}

func (c *videoController) Save(ctx *gin.Context) (*entity.Video, error) {
	var v entity.Video
	err := ctx.ShouldBindJSON(&v)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(v)
	if err != nil {
		return nil, err
	}
	c.svc.Save(&v)
	return &v, nil
}

func (c *videoController) ShowAll(ctx *gin.Context) map[string]interface{} {
	videos := c.svc.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	return data
}