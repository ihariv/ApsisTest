package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

const ID = "users"

// postUsers adds an users from JSON received in the request body.
func (us *Users) PutUsers(c *gin.Context) {
	var newUser User

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser.CounterUniq = new(CounterUniq)
	newUser.CounterUniq.RWMutex = new(sync.RWMutex)
	newUser.Password, newUser.PasswordSecure = newUser.PasswordSecure, newUser.Password

	// Add the new album to the slice.
	*us = append(*us, &newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

// deleteUsers adds an users from JSON received in the request body.
func (us *Users) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Delete user.
	for i, user := range users {
		if user.ID == id {
			users[i], users[len(users)-1] = users[len(users)-1], users[i]
			users = users[:len(users)-1]
		}
	}
	c.IndentedJSON(http.StatusAccepted, id)

}

// GetUserByID locates the users whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func (us *Users) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user := us.getUserByID(id)

	if user.ID != "" {
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}

// IncreateUserByID locates the users whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func (us *Users) IncreateCounter(c *gin.Context) {
	id := c.Param("id")
	user := us.getUserByID(id)

	if user.ID != "" {
		user.CounterUniq.Increate()
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (us *Users) getUserByID(id string) (u *User) {

	for _, user := range *us {
		if user.ID == id {
			u = user
			break
		}
	}
	return
}

// getUsers responds with the list of all users as JSON.
func (us *Users) getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, us)
}

// SumCounters calculate counters for all users in list || team.
func (us *Users) SumCounters() (cnt int) {
	for _, user := range *us {
		if user.ID != "" {
			cnt += user.Count
		}
	}
	return
}
