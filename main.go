package main

import (
	"github.com/gin-gonic/gin"
	"userStories/modules/login"
	"userStories/modules/users"
)

// album represents data about a record album.

func main(){

	userList := users.GetUsers()
	countersList := users.GetCounters(userList)

	router := gin.Default()
	router.Use(gin.Logger())

	cg := initCollect(router)
	cg.AddNodes(
		login.Login{},
		countersList,
		userList,
		)

	router.Run("localhost:8080")
}
