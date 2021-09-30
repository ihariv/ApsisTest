package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PutCounter adds an counter from JSON received in the request body.
func (c *Counters) PutCounter(ctx *gin.Context) {
	var newCounter Counter

	// Call BindJSON to bind the received JSON to
	// newCounter.
	if err := ctx.BindJSON(&newCounter); err != nil {
		return
	}
	newCounter.Users = new(Users)

	// Add the new counter to the slice.

	c.Lock()
	c.Counters = append(c.Counters, newCounter)
	c.Unlock()
	ctx.IndentedJSON(http.StatusCreated, newCounter)

}

// postUsers adds an counter from JSON received in the request body.
func (c *Counters) PutUserInCounter(ctx *gin.Context) {
	userId := ctx.Param("userId")
	counterId := ctx.Param("id")
	user := c.Users.getUserByID(userId)
	counter := counters.getCounterByID(counterId)

	counter.Users.Append(user)

	ctx.IndentedJSON(http.StatusOK, counter)

}

// postUsers adds an counter from JSON received in the request body.
func (c *Counters) DeleteCounter(ctx *gin.Context) {
	id := ctx.Param("id")

	// Delete user.
	deletedItem := -1
	for i, counter := range c.Counters {
		if counter.ID == id {
			deletedItem = i
			break
		}
	}
	if deletedItem < 0 {
		ctx.IndentedJSON(http.StatusNotFound, id)
		return
	}
	c.Counters[deletedItem] = c.Counters[len(c.Counters)-1]
	c.Counters = c.Counters[:len(c.Counters)-1]
	ctx.IndentedJSON(http.StatusAccepted, c.Counters)

}

// postUsers adds an counter from JSON received in the request body.
func (c *Counters) DeleteUserFromCounter(ctx *gin.Context) {

	userId := ctx.Param("userId")
	counterId := ctx.Param("id")
	// user := c.Users.getUserByID(userId)
	counter := counters.getCounterByID(counterId)

	counter.Users.Remove(userId)
	// Delete user.
	DeletedItem := -1
	for i, user := range *counter.Users {
		if user.ID == userId {
			DeletedItem = i
			break
		}
	}

	if DeletedItem < 0 {
		ctx.IndentedJSON(http.StatusNotFound, userId)
		return
	}
	users[DeletedItem] = users[len(users)-1]
	users = users[:len(users)-1]
	ctx.IndentedJSON(http.StatusOK, counter)

}

// getUserByID locates the counter whose ID value matches the id
// parameter sent by the client, then returns that counter as a response.
func (c *Counters) GetCounterByID(ctx *gin.Context) {
	id := ctx.Param("id")
	counter := c.getCounterByID(id)

	counter.SumForUsers = counter.Users.SumCounters()
	if counter.ID != "" {
		ctx.IndentedJSON(http.StatusOK, counter)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (c *Counters) getCounterByID(id string) (cnts *Counter) {
	c.RLock()
	defer c.RUnlock()

	for _, counter := range c.Counters {
		if counter.ID == id {
			cnts = &counter
			cnts.SumForUsers = cnts.Users.SumCounters()
			break
		}
	}
	c.SumForUsers = cnts.SumCounters()
	return
}

// getUsers responds with the list of all counter as JSON.
func (c *Counters) getCounters(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, c)
}
