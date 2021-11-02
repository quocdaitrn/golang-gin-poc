package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// gindump "github.com/tpkeeper/gin-dump"

	"github/quocdaitrn/golang-gin-poc/controller"
	"github/quocdaitrn/golang-gin-poc/middleware"
	"github/quocdaitrn/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		v, err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, v)
		}
	})

	server.Run(":8080")
}
