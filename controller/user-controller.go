package controller

import (
	"log"
	"net/http"

	"github.com/NajeebTyson/hello-gin/model"
	"github.com/gin-gonic/gin"
)

// GetUser return the user by id
func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, ok := model.Users[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// AddUser add a new user to store
func AddUser(c *gin.Context) {
	var newUser model.User

	err := c.BindJSON(&newUser)
	if err != nil {
		log.Fatal("Invalid add user request")

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
	}

	log.Printf("New user: %q", newUser)
	model.Users[newUser.ID] = newUser

	c.JSON(http.StatusOK, gin.H{
		"message": "User added",
		"data":    newUser,
	})
}

// DeleteUser remove user from store
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, ok := model.Users[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	delete(model.Users, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}

// UpdateUser update the user in store
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updateUser model.User
	err := c.BindJSON(&updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	_, ok := model.Users[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	model.Users[id] = updateUser
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
		"data":    updateUser,
	})
}
