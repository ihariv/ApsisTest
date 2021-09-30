package controller

import (
	"github.com/gin-gonic/gin"
	"userStories/dto"
	"userStories/modules/users"
	"userStories/modules/users/service"
)

//login contorller interface
type (
	LoginController interface {
		Login(ctx *gin.Context) string
	}
	AccessDetails struct {
		AccessUuid string
		UserId     int
	}
)

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler() LoginController {
	var jwtService service.JWTService = service.JWTAuthService()
	var loginService service.LoginService = users.GetUsers()
	return &loginController{
		loginService: loginService,
		jWtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}