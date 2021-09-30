package login

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"userStories/controller"
)

func (l Login) Route(router *gin.Engine) {

	var loginController controller.LoginController = controller.LoginHandler()

	router.Use(sessions.Sessions("token", sessions.NewCookieStore([]byte("secret"))))

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			session := sessions.Default(ctx)
			session.Set(USERKEY, token)
			if err := session.Save(); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})


	router.GET("/logout", logout)

}
