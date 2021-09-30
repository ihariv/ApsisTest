package users

import (
	"github.com/gin-gonic/gin"
)

var (
	userkey="user"
)

func (us Users) Route(router *gin.Engine) {

	private := router.Group("/" + ID)
	private.Use(AuthRequired)
	{
		private.GET("/", us.getUsers)
		private.PUT("/", us.PutUsers)
		private.GET("/:id/increate", us.IncreateCounter)
		private.DELETE("/:id", us.DeleteUser)
		private.GET("/:id", us.GetUserByID)
	}


}
