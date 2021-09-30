package users

import "github.com/gin-gonic/gin"

func (c Counters) Route(router *gin.Engine) {

	privateCounter := router.Group("/counter")
	privateCounter.Use(AuthRequired)
	{
		privateCounter.GET("/", c.getCounters)
		privateCounter.PUT("/", c.PutCounter)
		privateCounter.PUT("/:id/:userId", c.PutUserInCounter)
		privateCounter.DELETE("/:id", c.DeleteCounter)
		privateCounter.DELETE("/:id/:userId", c.DeleteUserFromCounter)
		privateCounter.GET("/:id", c.GetCounterByID)
	}

}
