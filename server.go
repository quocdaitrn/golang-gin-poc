package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// gindump "github.com/tpkeeper/gin-dump"

	"github/quocdaitrn/golang-gin-poc/controller"
	"github/quocdaitrn/golang-gin-poc/helpers"
	"github/quocdaitrn/golang-gin-poc/middleware"
	"github/quocdaitrn/golang-gin-poc/service"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   helpers.JWTService   = helpers.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger())

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			v, err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, v)
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", videoController.ShowAll(ctx))
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
