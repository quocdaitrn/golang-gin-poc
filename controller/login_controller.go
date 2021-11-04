package controller

import (
	"github.com/gin-gonic/gin"

	"github/quocdaitrn/golang-gin-poc/dto"
	"github/quocdaitrn/golang-gin-poc/helpers"
	"github/quocdaitrn/golang-gin-poc/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   helpers.JWTService
}

func NewLoginController(loginService service.LoginService,
	jWtService helpers.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
