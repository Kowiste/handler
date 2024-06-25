package main

import (
	"handler/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createUser adds a new user
func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBind(&newUser); err != nil {
		errors.RestError(c.Writer, err)
		return
	}

	users[newUser.ID] = newUser
	c.JSON(http.StatusCreated, newUser)
}

// getUser retrieves a user by ID
func getUser(c *gin.Context) {
	id := c.Param("id")
	user, exists := users[id]
	if !exists {
		errors.RestError(c.Writer, errors.NewString("User not found", errors.EErrorNotFound))
		return
	}
	c.JSON(http.StatusOK, user)
}

// updateUser updates a user by ID
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User
	if err := c.ShouldBind(&updatedUser); err != nil {
		errors.RestError(c.Writer, err)
		return
	}

	user, exists := users[id]
	if !exists {
		errors.RestError(c.Writer, errors.NewString("User not found", errors.EErrorNotFound))
		return
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	users[id] = user

	c.JSON(http.StatusOK, user)
}

// deleteUser removes a user by ID
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	_, exists := users[id]
	if !exists {
		errors.RestError(c.Writer, errors.NewString("User not found", errors.EErrorNotFound))
		return
	}

	delete(users, id)
	c.JSON(http.StatusNoContent, nil)
}
